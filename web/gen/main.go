// Command gen renders the golings catalog data (catalog.ts, per-exercise
// detail markdown, redirects) from info.toml + exercise sources + per-topic
// READMEs. Run from the repo root: `go run ./web/gen`.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/madhank93/golings/golings/exercises"
	"github.com/madhank93/golings/web/internal/play"
)

const (
	infoFile    = "info.toml"
	legacyDir   = "web/src/content/docs/curriculum" // retired pages, absorbed by /catalog — removed if present
	dataDir     = "web/src/data"                    // catalog.ts + redirects for /catalog
	detailsDir  = "web/src/data/lesson-details"     // per-exercise detail markdown
	seriesDir   = "web/src/content/docs/series"     // retired chapter pages — removed if present
	chaptersDir = "web/src/data/chapters"           // per-topic chapter markdown for the catalog
	exerciseR   = "exercises"
)

// tier groups topics into a beginner→advanced section (mirrors CURRICULUM.md).
type tier struct {
	name   string
	topics []string
}

var tiers = []tier{
	{"Beginner · Fundamentals", []string{"variables", "primitive_types", "if", "switch", "functions", "more_functions", "strings"}},
	{"Beginner · Collections & Loops", []string{"arrays", "slices", "maps", "range"}},
	{"Intermediate · Types & Methods", []string{"structs", "pointers", "methods", "interfaces", "enums", "type_aliases"}},
	{"Intermediate · Functions & Errors", []string{"anonymous_functions", "defer", "errors"}},
	{"Intermediate · Generics & Modern Go", []string{"generics", "modern", "iterators"}},
	{"Intermediate · Testable Design", []string{"dependency_injection", "mocking"}},
	{"Advanced · Concurrency", []string{"concurrent", "channels", "select", "sync", "context", "concurrency_patterns", "goroutine_safety", "synctest"}},
	{"Advanced · Standard Library & I/O", []string{"stdlib_essentials", "maps_package", "structured_logging", "reflection", "unsafe_pkg", "files", "http_client"}},
	{"Advanced · Building Applications", []string{"http_server", "http_server_advanced", "cli"}},
	{"Advanced · Testing & Applied", []string{"testing_advanced", "profiling", "applied"}},
	{"Capstone · Log-Ingest Service", []string{"capstone"}},
}

// tierColors gives every topic in a tier the same chip color on /catalog.
var tierColors = []string{
	"#4fa86d", "#3b9eff", "#7c6af5", "#d29922", "#9b5de5",
	"#1f8a9c", "#c53030", "#e36f0e", "#db6d28", "#e85d9f",
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gen:", err)
		os.Exit(1)
	}
}

func run() error {
	exs, err := exercises.List(infoFile)
	if err != nil {
		return err
	}
	byTopic := map[string][]exercises.Exercise{}
	for _, e := range exs {
		t := topicOf(e.Path)
		byTopic[t] = append(byTopic[t], e)
	}

	// Every topic must be placed in a tier, or its exercises silently vanish
	// from the site (this happened to the 9 Phase-5 topics).
	if err := checkCoverage(byTopic); err != nil {
		return err
	}

	// The old per-topic curriculum pages were absorbed by /catalog (see the
	// redirects) — clear them out of stale checkouts so they don't build.
	if err := os.RemoveAll(legacyDir); err != nil {
		return err
	}

	if err := writeCatalog(byTopic); err != nil {
		return err
	}
	if err := writeChapters(); err != nil {
		return err
	}
	return writeRedirects()
}

// writeChapters renders each topic README as chapter markdown for the catalog
// modal. The README stays the single source: the TUI and GitHub read it
// directly, the site reads this rendering of it.
func writeChapters() error {
	// /series/* pages were folded into the catalog — drop them from stale
	// checkouts so they don't build.
	if err := os.RemoveAll(seriesDir); err != nil {
		return err
	}
	if err := os.RemoveAll(chaptersDir); err != nil {
		return err
	}
	if err := os.MkdirAll(chaptersDir, 0o755); err != nil {
		return err
	}

	for _, ti := range tiers {
		for _, topic := range ti.topics {
			body := readme(topic)
			if body == "" {
				continue
			}
			var b strings.Builder
			fmt.Fprintf(&b, "---\ntitle: %s\n---\n\n", pretty(topic))
			fmt.Fprintf(&b, "%s\n", chapterHTML(exercises.StripFences(body, "ascii")))
			if err := os.WriteFile(filepath.Join(chaptersDir, topic+".md"), []byte(b.String()), 0o644); err != nil {
				return err
			}
		}
	}
	return nil
}

// chapterHTML prepares a chapter for the catalog modal: neighbour links become
// catalog deep links, and mermaid fences become the <pre class="mermaid"> the
// modal's lazy mermaid pass renders (a fence would stay a code block, since the
// modal injects its HTML after the page has loaded).
func chapterHTML(md string) string {
	md = topicLink.ReplaceAllString(md, "](/catalog/?chapter=$1)")
	return mermaidFence.ReplaceAllStringFunc(md, func(block string) string {
		src := mermaidFence.FindStringSubmatch(block)[1]
		return "<pre class=\"mermaid\">" + escapeHTML(strings.TrimRight(src, "\n")) + "</pre>"
	})
}

// escapeHTML escapes the three characters that would otherwise end the <pre>
// early or be read as markup. Mermaid parses the element's text content, so it
// sees the original characters back.
func escapeHTML(s string) string {
	return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;").Replace(s)
}

// Link rewrites applied to chapters and notes on their way to the catalog.
var (
	topicLink    = regexp.MustCompile(`\]\(\.\./([a-z_]+)/\)`)
	readmeLink   = regexp.MustCompile(`\]\(\.\./README\.md\)`)
	mermaidFence = regexp.MustCompile("(?s)```mermaid\n(.*?)```")
)

// writeRedirects maps the retired curriculum URLs onto the catalog so old
// links keep working: topic pages deep-link into the topic filter.
func writeRedirects() error {
	m := map[string]string{"/curriculum/": "/catalog/"}
	for _, ti := range tiers {
		for _, topic := range ti.topics {
			m["/curriculum/"+topic+"/"] = "/catalog/?topic=" + topic
		}
	}
	out, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dataDir, "curriculum-redirects.json"), append(out, '\n'), 0o644)
}

// checkCoverage fails when info.toml has a topic the tiers table doesn't place.
func checkCoverage(byTopic map[string][]exercises.Exercise) error {
	placed := map[string]bool{}
	for _, ti := range tiers {
		for _, t := range ti.topics {
			placed[t] = true
		}
	}
	var missing []string
	for t := range byTopic {
		if !placed[t] {
			missing = append(missing, t)
		}
	}
	if len(missing) > 0 {
		sort.Strings(missing)
		return fmt.Errorf("topics missing from the tiers table: %s", strings.Join(missing, ", "))
	}
	return nil
}

// writeCatalog emits the /catalog data: src/data/catalog.ts (typed entries the
// page renders at build time) and src/data/lesson-details/<name>.md (source,
// hint and solution, fetched into the modal on demand so spoilers never load
// with the table).
func writeCatalog(byTopic map[string][]exercises.Exercise) error {
	if err := os.RemoveAll(detailsDir); err != nil {
		return err
	}
	if err := os.MkdirAll(detailsDir, 0o755); err != nil {
		return err
	}

	var b strings.Builder
	b.WriteString("// AUTO-GENERATED by `go run ./web/gen` from info.toml — do not hand-edit.\n\n")
	b.WriteString("export type CatalogEntry = {\n")
	b.WriteString("  topic: string;\n  slug: string;\n  mode: 'compile' | 'test';\n")
	b.WriteString("  description: string;\n  path: string;\n  play?: string;\n  playWasm?: true;\n};\n\n")

	b.WriteString("export const TOPICS: Record<string, { label: string; tier: string; color: string; learn: string }> = {\n")
	for i, ti := range tiers {
		color := tierColors[i%len(tierColors)]
		for _, topic := range ti.topics {
			fmt.Fprintf(&b, "  %s: { label: %s, tier: %s, color: %s, learn: %s },\n",
				js(topic), js(pretty(topic)), js(ti.name), js(color), js(learnText(topic)))
		}
	}
	b.WriteString("};\n\n")

	// Playground snippets are refreshed out of band (`mise run playground-links`)
	// because sharing needs the network and this generator must run offline in
	// CI. An exercise edited since its last share has no current snippet, so it
	// simply gets no link.
	links, err := play.LoadLinks()
	if err != nil {
		return err
	}
	var unlinked []string

	b.WriteString("export const CATALOG: CatalogEntry[] = [\n")
	for _, ti := range tiers {
		for _, topic := range ti.topics {
			for _, e := range byTopic[topic] {
				link, linked := play.LinkFor(links, e)
				if !linked {
					unlinked = append(unlinked, e.Name)
				}
				fmt.Fprintf(&b, "  { topic: %s, slug: %s, mode: %s, description: %s, path: %s%s },\n",
					js(topic), js(e.Name), js(e.Mode), js(e.Description()), js(e.Path), playFields(link, linked))
				if err := writeDetail(e); err != nil {
					return err
				}
			}
		}
	}
	if len(unlinked) > 0 {
		fmt.Fprintf(os.Stderr, "gen: no playground snippet for %d exercise(s) — run `mise run playground-links`: %s\n",
			len(unlinked), strings.Join(unlinked, " "))
	}
	b.WriteString("];\n")
	return os.WriteFile(filepath.Join(dataDir, "catalog.ts"), []byte(b.String()), 0o644)
}

// playFields renders the optional playground properties of a catalog entry.
// playWasm marks the multi-file snippets only goplay.tools can run. A blocked
// link renders nothing: the snippet exists but the playground cannot run it.
func playFields(l play.Link, linked bool) string {
	if !linked || l.Blocked {
		return ""
	}
	out := ", play: " + js(l.ID)
	if l.Wasm {
		out += ", playWasm: true"
	}
	return out
}

// writeDetail renders one exercise's modal content as markdown: the
// broken-on-purpose source, then the hint and the worked solution (read from
// the solution branch), each behind a <details> so opening the modal doesn't
// spoil anything. All of it is generated here — nothing is maintained by hand
// in the frontend.
func writeDetail(e exercises.Exercise) error {
	files, err := filepath.Glob(filepath.Join(filepath.Dir(e.Path), "*.go"))
	if err != nil {
		return err
	}
	sort.Strings(files)

	var b strings.Builder
	fmt.Fprintf(&b, "---\ntitle: %s\n---\n\n", e.Name)

	// One accordion per part, in the order a learner meets them: the broken
	// source is open, everything that gives the answer away starts closed.
	var src strings.Builder
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		fmt.Fprintf(&src, "```go title=%q\n%s\n```\n\n", filepath.Base(f), strings.TrimRight(string(data), "\n"))
	}
	b.WriteString(section("What you start with", src.String(), true))

	if h := strings.TrimSpace(e.Hint); h != "" {
		b.WriteString(section("Hint", "```text\n"+h+"\n```\n", false))
	}

	var sol strings.Builder
	for _, f := range files {
		data, ok := solutionFile(f)
		if !ok {
			continue
		}
		fmt.Fprintf(&sol, "```go title=%q\n%s\n```\n\n", filepath.Base(f), strings.TrimRight(data, "\n"))
	}
	if sol.Len() > 0 {
		b.WriteString(section("Solution", sol.String(), false))
	}

	// Teaching notes: annotated walk-through, gotchas and references. Closed by
	// default like the solution, since they contain the answer.
	if n := strings.TrimSpace(e.Notes()); n != "" {
		n = readmeLink.ReplaceAllString(n, "](/catalog/?chapter="+topicOf(e.Path)+")")
		b.WriteString(section("Explanation", exercises.StripFences(n, "ascii"), false))
	}

	return os.WriteFile(filepath.Join(detailsDir, e.Name+".md"), []byte(b.String()), 0o644)
}

// section wraps one part of an exercise in a <details> the catalog styles as an
// accordion card. Spoilers (hint, solution, explanation) stay closed.
func section(title, body string, open bool) string {
	attr := ""
	if open {
		attr = " open"
	}
	return fmt.Sprintf("<details%s>\n<summary>%s</summary>\n\n%s\n</details>\n\n", attr, title, strings.TrimRight(body, "\n")+"\n")
}

// solutionRef is the first git ref that resolves to the solution branch, or
// "" when none does (then the modal simply has no solution section).
var solutionRef = func() string {
	for _, ref := range []string{"origin/solution", "solution"} {
		if err := exec.Command("git", "rev-parse", "--verify", "--quiet", ref).Run(); err == nil {
			return ref
		}
	}
	fmt.Fprintln(os.Stderr, "gen: no solution branch found — catalog details will have no solution sections")
	return ""
}()

// solutionFile returns the worked version of path from the solution branch.
func solutionFile(path string) (string, bool) {
	if solutionRef == "" {
		return "", false
	}
	out, err := exec.Command("git", "show", solutionRef+":"+filepath.ToSlash(path)).Output()
	if err != nil {
		return "", false
	}
	return string(out), true
}

// js renders s as a JavaScript string literal.
func js(s string) string {
	out, _ := json.Marshal(s)
	return string(out)
}

var mdLink = regexp.MustCompile(`\[([^\]]+)\]\([^)]+\)`)

// learnText returns the topic's opening prose as popover-ready plain text:
// links unwrapped, paragraphs kept as blank-line breaks. Only the intro — the
// rest of the chapter has its own page at /series/<topic>.
func learnText(topic string) string {
	var paras []string
	for _, para := range strings.Split(intro(readme(topic)), "\n\n") {
		para = strings.TrimSpace(para)
		if para == "" {
			continue
		}
		para = strings.ReplaceAll(para, "\n", " ")
		paras = append(paras, mdLink.ReplaceAllString(para, "$1"))
	}
	return strings.Join(paras, "\n\n")
}

// intro returns a chapter's opening prose: everything before its first section
// heading. It has to stand on its own — the catalog popover shows only this.
func intro(body string) string {
	if i := strings.Index(body, "\n## "); i >= 0 {
		body = body[:i]
	}
	return strings.TrimSpace(body)
}

// readme returns the topic README with its leading H1 stripped (the popover
// has a title already), or "" if there is no README.
func readme(topic string) string {
	data, err := os.ReadFile(filepath.Join(exerciseR, topic, "README.md"))
	if err != nil {
		return ""
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) > 0 && strings.HasPrefix(lines[0], "# ") {
		lines = lines[1:]
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func topicOf(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 2 {
		return parts[1]
	}
	return path
}

func pretty(topic string) string {
	words := strings.Split(topic, "_")
	for i, w := range words {
		if w != "" {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

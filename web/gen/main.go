// Command gen renders the golings curriculum into Starlight content pages from
// info.toml + per-topic READMEs. Run from the repo root: `go run ./web/gen`.
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

	"github.com/mauricioabreu/golings/golings/exercises"
)

const (
	infoFile   = "info.toml"
	outDir     = "web/src/content/docs/curriculum"        // overview (index.md)
	topicsDir  = "web/src/content/docs/curriculum/topics" // one page per topic
	dataDir    = "web/src/data"                           // catalog.ts for /catalog
	detailsDir = "web/src/data/lesson-details"            // per-exercise hint markdown
	baseURL    = ""                                       // empty = served at root (custom domain); use "/golings" for project pages
	repoBlob   = "https://github.com/madhank93/golings/blob/main"
	exerciseR  = "exercises"
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

	if err := os.RemoveAll(outDir); err != nil {
		return err
	}
	if err := os.MkdirAll(topicsDir, 0o755); err != nil {
		return err
	}

	order := 0
	for _, ti := range tiers {
		for _, topic := range ti.topics {
			order++
			if err := writeTopic(order, topic, ti.name, byTopic[topic]); err != nil {
				return err
			}
		}
	}
	if err := writeOverview(byTopic); err != nil {
		return err
	}
	return writeCatalog(byTopic)
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

func writeOverview(byTopic map[string][]exercises.Exercise) error {
	total, topics := 0, 0
	for _, ti := range tiers {
		for _, topic := range ti.topics {
			total += len(byTopic[topic])
			topics++
		}
	}
	var b strings.Builder
	b.WriteString("---\ntitle: Curriculum\ndescription: The full golings track, beginner to advanced.\n---\n\n")
	fmt.Fprintf(&b, "%d exercises across %d topics, grouped into a progressive track. Work through them in order with `mise run watch`, or browse the whole set in the filterable [Catalog](%s/catalog/).\n\n", total, topics, baseURL)
	for _, ti := range tiers {
		b.WriteString("## " + ti.name + "\n\n")
		for _, topic := range ti.topics {
			n := len(byTopic[topic])
			fmt.Fprintf(&b, "- [%s](%s/curriculum/%s/) — %d exercise%s\n", pretty(topic), baseURL, topic, n, plural(n))
		}
		b.WriteString("\n")
	}
	return os.WriteFile(filepath.Join(outDir, "index.md"), []byte(b.String()), 0o644)
}

func writeTopic(order int, topic, tierName string, exs []exercises.Exercise) error {
	var b strings.Builder
	fmt.Fprintf(&b, "---\ntitle: %s\nslug: curriculum/%s\nsidebar:\n  order: %d\n---\n\n", pretty(topic), topic, order)
	fmt.Fprintf(&b, "*%s*\n\n", tierName)

	if intro := readme(topic); intro != "" {
		b.WriteString(intro + "\n\n")
	}

	b.WriteString("## Exercises\n\n")
	for _, e := range exs {
		fmt.Fprintf(&b, "### %s `%s`\n\n", e.Name, e.Mode)
		if d := e.Description(); d != "" {
			fmt.Fprintf(&b, "%s\n\n", d)
		}
		if h := strings.TrimSpace(e.Hint); h != "" {
			b.WriteString("<details>\n<summary>Show hint</summary>\n\n```\n" + h + "\n```\n\n</details>\n\n")
		}
		fmt.Fprintf(&b, "[Source](%s/%s)\n\n", repoBlob, e.Path)
	}

	name := fmt.Sprintf("%02d-%s.md", order, topic)
	return os.WriteFile(filepath.Join(topicsDir, name), []byte(b.String()), 0o644)
}

// writeCatalog emits the /catalog data: src/data/catalog.ts (typed entries the
// page renders at build time) and src/data/lesson-details/<name>.md (the hint,
// fetched into the modal on demand so spoilers never load with the table).
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
	b.WriteString("  description: string;\n  path: string;\n};\n\n")

	b.WriteString("export const TOPICS: Record<string, { label: string; tier: string; color: string; learn: string }> = {\n")
	for i, ti := range tiers {
		color := tierColors[i%len(tierColors)]
		for _, topic := range ti.topics {
			fmt.Fprintf(&b, "  %s: { label: %s, tier: %s, color: %s, learn: %s },\n",
				js(topic), js(pretty(topic)), js(ti.name), js(color), js(learnText(topic)))
		}
	}
	b.WriteString("};\n\n")

	b.WriteString("export const CATALOG: CatalogEntry[] = [\n")
	for _, ti := range tiers {
		for _, topic := range ti.topics {
			for _, e := range byTopic[topic] {
				fmt.Fprintf(&b, "  { topic: %s, slug: %s, mode: %s, description: %s, path: %s },\n",
					js(topic), js(e.Name), js(e.Mode), js(e.Description()), js(e.Path))
				if err := writeDetail(e); err != nil {
					return err
				}
			}
		}
	}
	b.WriteString("];\n")
	return os.WriteFile(filepath.Join(dataDir, "catalog.ts"), []byte(b.String()), 0o644)
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

	for _, f := range files {
		src, err := os.ReadFile(f)
		if err != nil {
			return err
		}
		fmt.Fprintf(&b, "```go title=%q\n%s\n```\n\n", filepath.Base(f), strings.TrimRight(string(src), "\n"))
	}

	if h := strings.TrimSpace(e.Hint); h != "" {
		b.WriteString("<details>\n<summary>Show hint (spoiler)</summary>\n\n```text\n" + h + "\n```\n\n</details>\n\n")
	}

	var sol strings.Builder
	for _, f := range files {
		src, ok := solutionFile(f)
		if !ok {
			continue
		}
		fmt.Fprintf(&sol, "```go title=%q\n%s\n```\n\n", filepath.Base(f), strings.TrimRight(src, "\n"))
	}
	if sol.Len() > 0 {
		b.WriteString("<details>\n<summary>Show solution (spoiler)</summary>\n\n" + sol.String() + "</details>\n")
	}

	return os.WriteFile(filepath.Join(detailsDir, e.Name+".md"), []byte(b.String()), 0o644)
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

// learnText returns the first prose paragraph of the topic README — the
// "what you learn here" line shown in the catalog's topic popover.
func learnText(topic string) string {
	intro := readme(topic)
	if intro == "" {
		return ""
	}
	for _, para := range strings.Split(intro, "\n\n") {
		para = strings.TrimSpace(para)
		if para == "" || strings.HasPrefix(para, "#") {
			continue
		}
		para = strings.ReplaceAll(para, "\n", " ")
		return mdLink.ReplaceAllString(para, "$1")
	}
	return ""
}

// readme returns the topic README with its leading H1 stripped (the page has a
// title already), or "" if there is no README.
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

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}

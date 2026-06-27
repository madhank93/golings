// Command gen renders the golings curriculum into Starlight content pages from
// info.toml + per-topic READMEs. Run from the repo root: `go run ./web/gen`.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mauricioabreu/golings/golings/exercises"
)

const (
	infoFile  = "info.toml"
	outDir    = "web/src/content/docs/curriculum"          // overview (index.md)
	topicsDir = "web/src/content/docs/curriculum/topics"   // one page per topic
	baseURL   = ""                                         // empty = served at root (custom domain); use "/golings" for project pages
	repoBlob  = "https://github.com/madhank93/golings/blob/main"
	exerciseR = "exercises"
)

// tier groups topics into a beginner→advanced section (mirrors CURRICULUM.md).
type tier struct {
	name   string
	topics []string
}

var tiers = []tier{
	{"Beginner · Fundamentals", []string{"variables", "primitive_types", "if", "switch", "functions", "more_functions", "strings"}},
	{"Beginner · Collections & Loops", []string{"arrays", "slices", "maps", "range"}},
	{"Intermediate · Types & Methods", []string{"structs", "pointers", "methods", "interfaces", "enums"}},
	{"Intermediate · Functions & Errors", []string{"anonymous_functions", "defer", "errors"}},
	{"Intermediate · Generics & Modern Go", []string{"generics", "modern"}},
	{"Intermediate · Testable Design", []string{"dependency_injection", "mocking"}},
	{"Advanced · Concurrency", []string{"concurrent", "channels", "select", "sync", "context", "concurrency_patterns"}},
	{"Advanced · Standard Library & I/O", []string{"stdlib_essentials", "reflection", "files", "http_client"}},
	{"Advanced · Building Applications", []string{"http_server", "cli"}},
	{"Advanced · Testing & Applied", []string{"testing_advanced", "applied"}},
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
	return writeOverview(byTopic)
}

func writeOverview(byTopic map[string][]exercises.Exercise) error {
	var b strings.Builder
	b.WriteString("---\ntitle: Curriculum\ndescription: The full golings track, beginner to advanced.\n---\n\n")
	b.WriteString("97 exercises across 32 topics, grouped into a progressive track. Work through them in order with `mise run watch`.\n\n")
	total := 0
	for _, ti := range tiers {
		b.WriteString("## " + ti.name + "\n\n")
		for _, topic := range ti.topics {
			n := len(byTopic[topic])
			total += n
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

package play

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/madhank93/golings/golings/exercises"
)

// write creates files under dir from a path → content map.
func write(t *testing.T, dir string, files map[string]string) {
	t.Helper()
	for name, body := range files {
		p := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestBuildModuleKeepsItsTree(t *testing.T) {
	dir := t.TempDir()
	write(t, dir, map[string]string{
		"go.mod":      "module example.com/bank\n",
		"report.go":   "package bank\n",
		"ledger/l.go": "package ledger\n",
		"notes.md":    "spoiler\n",
	})
	s, ok := Build(exercises.Exercise{Path: filepath.Join(dir, "report.go")})
	if !ok || !s.Wasm {
		t.Fatalf("Build = ok %v wasm %v, want a wasm snippet", ok, s.Wasm)
	}
	for _, want := range []string{"-- go.mod --\nmodule example.com/bank\n", "-- ledger/l.go --", "-- report.go --"} {
		if !strings.Contains(s.Source, want) {
			t.Errorf("snippet lacks %q:\n%s", want, s.Source)
		}
	}
	if strings.Contains(s.Source, "spoiler") {
		t.Error("snippet includes notes.md")
	}
}

func TestBuildSkipsWhatGoplayCannotRun(t *testing.T) {
	cases := map[string]struct {
		files  map[string]string
		noPlay bool
	}{
		"workspace": {files: map[string]string{"go.work": "go 1.26.0\n", "a.go": "package a\n"}},
		"vendored":  {files: map[string]string{"go.mod": "module m\n", "a.go": "package a\n", "vendor/modules.txt": ""}},
		"noplay":    {files: map[string]string{"go.mod": "module m\n", "a.go": "package a\n"}, noPlay: true},
	}
	for name, c := range cases {
		dir := t.TempDir()
		write(t, dir, c.files)
		if _, ok := Build(exercises.Exercise{Path: filepath.Join(dir, "a.go"), NoPlay: c.noPlay}); ok {
			t.Errorf("%s: Build returned a snippet", name)
		}
	}
}

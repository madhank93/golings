// Package play turns golings exercises into Go Playground snippets and tracks
// the share ids the catalog links to.
//
// Two shapes come out of it, because the two playgrounds run different things:
//
//   - A single-file exercise becomes one `package main` file. go.dev runs it,
//     and runs `go test` on it when it declares TestXxx and no main.
//   - A multi-file exercise (the capstone) becomes a txtar fileset with a
//     go.mod. go.dev refuses to test a multi-file snippet (golang/go#77234),
//     so those are linked to goplay.tools, whose in-browser WebAssembly
//     runtime does run them.
//
// Both are stored by the same playground snippet service, so one share id
// opens at either site.
package play

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/madhank93/golings/golings/exercises"
)

// goVersion is the language version the generated go.mod asks for. It is
// deliberately below the runtimes both playgrounds offer, and above everything
// the exercises use.
const goVersion = "1.24"

// LinksFile holds the share ids, keyed by exercise name. It is committed so
// the catalog can be regenerated offline: refreshing it is a separate,
// network-bound step (`mise run playground-links`).
const LinksFile = "web/src/data/playground.json"

// Link is one exercise's snippet: the id, whether it needs the WebAssembly
// runtime, and the hash of the source it was created from. A hash that no
// longer matches means the exercise was edited and the snippet is stale — the
// catalog drops the link rather than sending a learner to the old code.
type Link struct {
	Hash string `json:"hash"`
	ID   string `json:"id"`
	Wasm bool   `json:"wasm,omitempty"`
	// Blocked marks a snippet the playground refuses to run: a build or run
	// timeout, which is how the sandbox reports an attempted network call. The
	// catalog drops the link rather than pointing a learner at a button that
	// only ever times out. A compile error is not blocked — the exercises ship
	// broken on purpose, and that failure is the lesson.
	Blocked bool `json:"blocked,omitempty"`
}

// Snippet is an exercise rendered for a playground.
type Snippet struct {
	Source string
	// Wasm marks a multi-file snippet: only goplay.tools' WebAssembly runtime
	// runs its tests.
	Wasm bool
}

// URL is where a single-file snippet is opened — the official playground.
func URL(id string) string { return "https://go.dev/play/p/" + id }

// WasmURL is where a snippet is opened when it needs file tabs and the
// WebAssembly runtime.
func WasmURL(id string) string { return "https://goplay.tools/snippet/" + id }

// Build renders an exercise for the playground. ok is false when the source
// cannot be read or its package clause is not one the playground accepts.
func Build(e exercises.Exercise) (Snippet, bool) {
	files, err := filepath.Glob(filepath.Join(filepath.Dir(e.Path), "*.go"))
	if err != nil || len(files) == 0 {
		return Snippet{}, false
	}
	sort.Strings(files)
	if len(files) == 1 {
		src, ok := single(e.Path)
		return Snippet{Source: src}, ok
	}
	src, ok := fileset(files)
	return Snippet{Source: src, Wasm: true}, ok
}

// single rewrites a one-file exercise for go.dev, which rejects any package
// but main. The test files are `package main_test`, which is the same package
// as far as the playground's single-file test run is concerned.
func single(path string) (string, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		switch strings.TrimSpace(line) {
		case "package main":
			return string(data), true
		case "package main_test":
			lines[i] = "package main"
			return strings.Join(lines, "\n"), true
		}
	}
	return "", false
}

// fileset renders a directory as a txtar snippet. The go.mod is what makes it
// a package the runtime can test under its own name, so the capstone keeps the
// package it teaches instead of being flattened into main.
func fileset(files []string) (string, bool) {
	pkg, ok := packageName(files[0])
	if !ok {
		return "", false
	}
	var b strings.Builder
	fmt.Fprintf(&b, "-- go.mod --\nmodule %s\n\ngo %s\n", pkg, goVersion)
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return "", false
		}
		fmt.Fprintf(&b, "-- %s --\n%s\n", filepath.Base(f), strings.TrimRight(string(data), "\n"))
	}
	return b.String(), true
}

func packageName(path string) (string, bool) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", false
	}
	for _, line := range strings.Split(string(data), "\n") {
		if name, found := strings.CutPrefix(strings.TrimSpace(line), "package "); found {
			return strings.TrimSuffix(strings.TrimSpace(name), "_test"), true
		}
	}
	return "", false
}

// Hash identifies the exact source a snippet was shared from.
func Hash(src string) string {
	sum := sha256.Sum256([]byte(src))
	return hex.EncodeToString(sum[:])
}

// LoadLinks reads LinksFile, treating a missing file as an empty set so a
// fresh checkout still generates (without playground links).
func LoadLinks() (map[string]Link, error) {
	data, err := os.ReadFile(LinksFile)
	if os.IsNotExist(err) {
		return map[string]Link{}, nil
	}
	if err != nil {
		return nil, err
	}
	links := map[string]Link{}
	if err := json.Unmarshal(data, &links); err != nil {
		return nil, err
	}
	return links, nil
}

// SaveLinks writes LinksFile. json.Marshal sorts map keys, so the committed
// file stays stable across runs.
func SaveLinks(links map[string]Link) error {
	out, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(LinksFile, append(out, '\n'), 0o644)
}

// LinkFor returns the stored link for an exercise, or ok=false when there is
// none or the stored one was made from different source.
func LinkFor(links map[string]Link, e exercises.Exercise) (Link, bool) {
	s, built := Build(e)
	if !built {
		return Link{}, false
	}
	l, found := links[e.Name]
	if !found || l.ID == "" || l.Hash != Hash(s.Source) {
		return Link{}, false
	}
	return l, true
}

// Command playground refreshes the Go Playground share ids the catalog links
// to. It talks to play.golang.org, so it is a manual step — never part of the
// site build, which has to work offline: run it after editing exercises and
// commit web/src/data/playground.json.
//
//	go run ./web/playground            # share what changed
//	go run ./web/playground -force     # re-share everything
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/madhank93/golings/golings/exercises"
	"github.com/madhank93/golings/web/internal/play"
)

// shareURL is deliberately not go.dev/_/share: that front end re-encodes the
// body and stores a mangled snippet. Both hosts write to the same snippet
// store, so ids created here open at go.dev/play/p/<id>.
const shareURL = "https://play.golang.org/share"

// compileURL runs a snippet the way the play button will. Probing is the only
// thing that separates an exercise which fails on purpose from one the sandbox
// cannot run at all.
const compileURL = "https://go.dev/_/compile"

const infoFile = "info.toml"

func main() {
	force := flag.Bool("force", false, "re-share every exercise, ignoring the stored hashes")
	flag.Parse()

	if err := run(*force); err != nil {
		fmt.Fprintln(os.Stderr, "playground:", err)
		os.Exit(1)
	}
}

func run(force bool) error {
	exs, err := exercises.List(infoFile)
	if err != nil {
		return err
	}
	links, err := play.LoadLinks()
	if err != nil {
		return err
	}

	var kept, shared, skipped []string
	for _, e := range exs {
		snip, ok := play.Build(e)
		if !ok {
			delete(links, e.Name)
			skipped = append(skipped, e.Name)
			continue
		}
		hash := play.Hash(snip.Source)
		if l, found := links[e.Name]; found && l.ID != "" && l.Hash == hash && !force {
			kept = append(kept, e.Name)
			continue
		}
		id, err := share(snip.Source)
		if err != nil {
			return fmt.Errorf("%s: %w", e.Name, err)
		}
		link := play.Link{Hash: hash, ID: id, Wasm: snip.Wasm}
		// go.dev cannot run a multi-file snippet at all, so probing one there
		// proves nothing: the wasm snippets are opened at goplay.tools.
		if !snip.Wasm {
			stuck, err := blocked(snip.Source)
			if err != nil {
				return fmt.Errorf("%s: %w", e.Name, err)
			}
			link.Blocked = stuck
		}
		links[e.Name] = link
		shared = append(shared, e.Name)
		target := play.URL(id)
		if snip.Wasm {
			target = play.WasmURL(id)
		}
		note := ""
		if link.Blocked {
			note = "  (the playground cannot run it — no link)"
		}
		fmt.Printf("shared %-14s %s%s\n", e.Name, target, note)
	}

	if err := play.SaveLinks(links); err != nil {
		return err
	}
	fmt.Printf("\n%d unchanged, %d shared, %d without a snippet%s\n",
		len(kept), len(shared), len(skipped), listing(skipped))
	fmt.Println("regenerate the catalog to pick the ids up: mise run gen-site")
	return nil
}

func listing(names []string) string {
	if len(names) == 0 {
		return ""
	}
	return " (" + strings.Join(names, " ") + ")"
}

// share posts one snippet and reads it back. The read-back is not paranoia: a
// share that silently stores re-encoded source still returns a valid-looking
// id, and the only way to see it is to fetch the snippet.
func share(src string) (string, error) {
	id, err := post(src)
	if err != nil {
		return "", err
	}
	got, err := fetch(id)
	if err != nil {
		return "", err
	}
	if got != src {
		return "", fmt.Errorf("snippet %s does not match the source it was shared from", id)
	}
	return id, nil
}

func post(src string) (string, error) {
	var last error
	for attempt := range 3 {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
		}
		resp, err := http.Post(shareURL, "text/plain", strings.NewReader(src))
		if err != nil {
			last = err
			continue
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			last = err
			continue
		}
		if resp.StatusCode != http.StatusOK {
			last = fmt.Errorf("share: %s: %s", resp.Status, bytes.TrimSpace(body))
			continue
		}
		id := string(bytes.TrimSpace(body))
		if id == "" {
			last = fmt.Errorf("share returned an empty id")
			continue
		}
		return id, nil
	}
	return "", last
}

func fetch(id string) (string, error) {
	resp, err := http.Get(play.URL(id) + ".go")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("fetch %s: %s", id, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// blocked reports whether the playground refused to run the snippet. Only a
// build or run timeout counts. Those timeouts are intermittent, so a snippet is
// called blocked only when every attempt timed out; a single clean run proves
// it is runnable.
func blocked(src string) (bool, error) {
	var last error
	for attempt := range 3 {
		if attempt > 0 {
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
		}
		out, err := compile(src)
		if err != nil {
			last = err
			continue
		}
		if !timedOut(out) {
			return false, nil
		}
		last = nil
	}
	if last != nil {
		return false, last
	}
	return true, nil
}

// result is the part of the playground's compile response this command reads.
type result struct {
	Errors string
	Events []struct {
		Message string
	}
}

func compile(src string) (result, error) {
	form := url.Values{"version": {"2"}, "body": {src}, "withVet": {"false"}}
	resp, err := http.PostForm(compileURL, form)
	if err != nil {
		return result{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return result{}, fmt.Errorf("compile: %s", resp.Status)
	}
	var out result
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return result{}, err
	}
	return out, nil
}

// timedOut distinguishes the sandbox giving up from the exercise failing. A
// build timeout is what the playground returns for a network call it will not
// make; a run that takes too long is the same refusal seen from the other side.
func timedOut(r result) bool {
	if strings.Contains(r.Errors, "timeout running go build") {
		return true
	}
	for _, ev := range r.Events {
		if strings.Contains(ev.Message, "process took too long") {
			return true
		}
	}
	return false
}

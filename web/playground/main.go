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
	"flag"
	"fmt"
	"io"
	"net/http"
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
		links[e.Name] = play.Link{Hash: hash, ID: id, Wasm: snip.Wasm}
		shared = append(shared, e.Name)
		url := play.URL(id)
		if snip.Wasm {
			url = play.WasmURL(id)
		}
		fmt.Printf("shared %-14s %s\n", e.Name, url)
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

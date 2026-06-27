# golings

[![CI](https://github.com/madhank93/golings/actions/workflows/ci.yml/badge.svg)](https://github.com/madhank93/golings/actions/workflows/ci.yml)
[![Docs](https://github.com/madhank93/golings/actions/workflows/pages.yml/badge.svg)](https://golings.madhan.app)

![gopher](misc/gopher-dance.gif)

> rustlings, but for Go.

Learn Go the [rustlings](https://github.com/rust-lang/rustlings) way: **112 small,
broken-on-purpose exercises** you fix one at a time, from variables to
concurrency, current through Go 1.26. An interactive terminal UI re-runs each
exercise the moment you save and only lets you advance when the tests pass **and**
`golangci-lint` is clean.

📖 **Docs & full curriculum: <https://golings.madhan.app>**

> This is a maintained fork of the original
> [golings by mauricioabreu](https://github.com/mauricioabreu/golings), rebuilt
> around a [mise](https://mise.jdx.dev/) toolchain, a Bubble Tea TUI, an expanded
> 32-topic curriculum, and a docs site.

## Quick start

This project uses [mise](https://mise.jdx.dev/) to pin the toolchain (Go 1.26,
gopls, golangci-lint) — you don't need Go installed globally.

```sh
brew install mise        # macOS; see mise docs for other platforms
git clone https://github.com/madhank93/golings
cd golings
mise install             # provisions Go, gopls, golangci-lint
mise run watch           # build + launch the interactive TUI
```

That's it. The TUI highlights the next unfinished exercise — open the file, make
it compile / pass its tests, remove the `// I AM NOT DONE` marker, and save.

### GitHub Codespaces

Open the repo in a Codespace (or VS Code **Dev Containers**). The devcontainer
installs mise and provisions the same pinned toolchain, then `mise run watch`.

## Working through the exercises

Exercises live in `exercises/<topic>/`. Each topic has a README with resources —
read it before you start.

In the watch TUI:

| Key | Action |
| --- | --- |
| `↑` / `↓` (or `k` / `j`) | move between exercises |
| `⏎` | run the selected exercise |
| `e` | open the exercise in `$VISUAL` / `$EDITOR` (falls back to `vi`) |
| `h` | toggle the hint |
| `r` | reset the exercise to its original state |
| `n` | jump to the next unfinished exercise |
| `/` | search exercises by name or topic (`⏎` runs the match, `Esc` cancels) |
| `q` | quit |

## Common commands

```sh
mise run watch           # interactive watch loop (the main one)
mise run list            # list every exercise + your progress
mise run update          # pull latest changes without losing your progress
mise run test            # test the tool itself (-race)
mise run lint            # lint the tool source
mise run site            # run the docs site locally
```

`mise run update` shelves your in-progress edits, pulls, then restores them. Your
completion record (`.golings-state.json`) is untracked, so your done count and
streak survive every update.

## Editor setup (VS Code)

The repo ships a `.vscode/` config. On open, VS Code prompts you to install the
recommended **Go** and **mise** extensions — accept both and reload. The mise
extension feeds VS Code the pinned `go`/`gopls`, so completion, format-on-save,
and `golangci-lint` work even when VS Code is launched outside a mise shell.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md).

## Learning resources

* [A Tour of Go](https://go.dev/tour/)
* [Go by Example](https://gobyexample.com)
* [Effective Go](https://go.dev/doc/effective_go)

## Other 'lings

* [rustlings](https://github.com/rust-lang/rustlings)
* [ziglings](https://github.com/ratfactor/ziglings)

## Credits

* Original [golings](https://github.com/mauricioabreu/golings) by Maurício Antunes.
* Gopher artwork by [@egonelbre](https://github.com/egonelbre/gophers) (CC0) and
  the Go gopher by Renée French.

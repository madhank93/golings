---
title: Getting started
description: Install golings and start the interactive exercise runner.
---

## Prerequisites

golings uses [mise](https://mise.jdx.dev/) to pin the toolchain (Go 1.26) and the
linter, so you don't need Go installed globally.

```sh
brew install mise        # macOS; see mise docs for other platforms
```

## Set up

```sh
git clone https://github.com/madhank93/golings
cd golings
mise install             # provisions Go 1.26, gopls, golangci-lint
mise run watch           # launches the interactive TUI
```

## Use in GitHub Codespaces

Prefer a zero-setup cloud environment? Open the repo in a
[devcontainer](https://containers.dev/) (GitHub Codespaces, or VS Code's
**Dev Containers** extension). It installs **mise** and provisions the exact same
pinned toolchain as local — Go 1.26, gopls, golangci-lint — then builds the
binary. Once the container is ready, just run:

```sh
mise run watch
```

`go`, `gopls`, and `golangci-lint` are on your `PATH` in the integrated terminal,
so editor completion and linting work too.

## Editor setup (VS Code)

The repo ships a `.vscode/` config. When you open it, VS Code prompts you to
install the recommended extensions — the **Go** extension and the **mise**
extension. Accept both and reload. The mise extension feeds VS Code the
mise-pinned `go` and `gopls`, so completion, hover, go-to-definition, format on
save, and `golangci-lint` work — even when VS Code is launched outside a mise
shell. No global Go install required.

## How it works

1. The TUI highlights the next unfinished exercise and shows its file path.
2. Open that file and make it compile / pass its tests.
3. Remove the `// I AM NOT DONE` marker when you think it's done.
4. **Save** — golings re-runs the exercise. It only advances when the tests pass
   **and** `golangci-lint` is clean.

### Keys

| Key | Action |
| --- | --- |
| `↑` / `↓` (or `k` / `j`) | move between exercises |
| `⏎` | run the selected exercise |
| `e` | open the exercise in `$EDITOR` (vi/vim/emacs/code…) |
| `h` | toggle the hint |
| `r` | reset the exercise to its original state |
| `n` | jump to the next unfinished exercise |
| `q` | quit |

### Streaks

Your streak is the number of **consecutive days, ending on your most recent
completion, on which you finished at least one exercise.** It is never stored as
a number — golings only records the timestamp of each first completion in
`.golings-state.json`, and the streak is *derived* from those timestamps each
time it's shown. Reset an exercise (`r`) and its completion is dropped, so the
streak adjusts automatically; miss a day and your next completion starts a fresh
streak of 1.

## Staying up to date

To pull the latest exercises and tool improvements **without losing your
progress**, run:

```sh
mise run update
```

It shelves your in-progress exercise edits, pulls, then restores them. Your
completion record (`.golings-state.json`) is untracked, so it is never touched —
your done count and streak survive every update.

## Other commands

```sh
mise run list            # list every exercise + progress
mise run update          # pull latest, keeping your progress
./bin/golings hint <name>   # print a hint
./bin/golings reset <name>  # restore an exercise to its original state
mise run test            # test the tool itself (-race)
mise run lint            # lint the tool source
```

Next: browse the [Curriculum](/curriculum/) to see the full track.

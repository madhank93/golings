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

## Other commands

```sh
mise run list            # list every exercise + progress
./bin/golings hint <name>   # print a hint
./bin/golings reset <name>  # restore an exercise to its original state
mise run test            # test the tool itself (-race)
mise run lint            # lint the tool source
```

Next: browse the [Curriculum](/golings/curriculum/) to see the full track.

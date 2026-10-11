## modules10 — go mod tidy

```
$ go mod tidy
```

```
require example.com/table v0.0.0

require example.com/pad v0.0.0 // indirect
```

**Why it works**

- `go mod tidy` reads every import in the module (tests included, for all
  platforms) and rewrites `go.mod` to match: it drops `oldfmt`, which nothing
  imports, and adds `pad`, which `table` needs.

**Under the hood**

- `pad` is marked `// indirect` because no package in this module imports it.
  Since go 1.17, `go.mod` lists every module that provides a package in the
  build, direct or not — that is what lets the `go` command prune the module
  graph and skip loading the `go.mod` files of irrelevant dependencies.
- tidy also adds any missing `go.sum` entries and removes stale ones.

**Common mistake**

- Editing `require` lines by hand. It is easy to miss an indirect dependency or
  a checksum; tidy derives both from the code and never forgets one.

**Key detail:** `go mod tidy -diff` (Go 1.23) changes nothing, prints the
needed edits as a diff and exits non-zero if there are any — a ready-made CI
check, and exactly what this exercise's test runs. Note tidy leaves `replace`
directives alone: the unused `oldfmt` replacement stays.

**See also:** modules5 (`require` + `replace`) · modules9 (workspaces) · the
[chapter](../README.md)

**References**

- Go Modules Reference — `go mod tidy`: https://go.dev/ref/mod#go-mod-tidy
- Go Modules Reference — Module graph pruning: https://go.dev/ref/mod#graph-pruning
- Go 1.23 release notes — Go command: https://go.dev/doc/go1.23#go-command

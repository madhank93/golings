## modules11 — vendoring

```
$ go mod vendor
```

```
# vendor/modules.txt
# example.com/slug v0.0.0 => ./slug
## explicit; go 1.26.0
example.com/slug
# example.com/slug => ./slug
```

**Why it works**

- `go mod vendor` copies the packages the build needs into `vendor/` and
  writes `vendor/modules.txt`, the manifest of which module version each came
  from. Once it agrees with `go.mod` again, the build reads `slug` from
  `vendor/`.

**Under the hood**

- When `vendor/` exists and `go.mod` says go 1.14 or later, the `go` command
  acts as if `-mod=vendor` were set: it builds from `vendor/` and never touches
  the module cache or network. Before using it, it checks `modules.txt` against
  `go.mod` — every explicit `require` and `replace` must be recorded there — and
  stops with `inconsistent vendoring` if they disagree.

**Common mistake**

- Adding or bumping a `require` and committing without re-vendoring. Every
  `go.mod` change needs another `go mod vendor`; `-mod=mod` only sidesteps the
  check for one command and is not a fix.

**Key detail:** vendoring trades repository size for a self-contained build —
useful for air-gapped or audited builds. Most projects rely on the module
proxy plus `go.sum` instead; vendor only when you need the copies in the repo.

**See also:** modules5 (`replace` to a local path) · modules10 (`go mod tidy`)
· the [chapter](../README.md)

**References**

- Go Modules Reference — Vendoring: https://go.dev/ref/mod#vendoring
- Go Modules Reference — `go mod vendor`: https://go.dev/ref/mod#go-mod-vendor
- Go 1.14 release notes — Go command: https://go.dev/doc/go1.14#go-command

## modules7 — the `go` directive sets the language version

```
module example.com/countdown

go 1.26.0
```

**Why it works**

- The `go` line is not documentation: it is the language version the module is
  compiled as. `for i := range n` over an int arrived in Go 1.22, so under
  `go 1.21` the compiler rejects it even though the installed toolchain is newer.

**Under the hood**

- `go build` passes the `go` line to the compiler as `-lang`. That is why the
  error reads `requires go1.22 or later (-lang was set to go1.21; check go.mod)`
  — the toolchain can do it, the module has not opted in.
- Since Go 1.21 the line is also a minimum: a toolchain older than the declared
  version refuses to build the module (or, with `GOTOOLCHAIN=auto`, downloads a
  newer one). The separate `toolchain` directive only *suggests* a toolchain to
  use; it never changes the language version.

**Common mistake**

- Bumping the Go install, or adding `toolchain go1.26.0`, and expecting new
  syntax to compile. Only the `go` line moves the language version.

**Key detail:** language semantics follow the line too, not just new syntax —
the Go 1.22 per-iteration loop variable only applies to modules declaring
`go 1.22` or later. Raising it is a deliberate change; `go mod edit -go=1.26.0`
does it without hand-editing.

**See also:** modules1 (writing go.mod) · modern2 (range over an int) · the
[chapter](../README.md)

**References**

- Go Modules Reference — `go` directive: https://go.dev/ref/mod#go-mod-file-go
- Go Toolchains: https://go.dev/doc/toolchain
- Go 1.22 release notes: https://go.dev/doc/go1.22#language

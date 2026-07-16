# The unsafe Package

`unsafe` steps outside Go's type safety for the rare cases that need it: struct
layout introspection and zero-copy conversions in hot paths. These exercises use
only the safe-ish, well-established idioms — but remember the tradeoff: the
compiler no longer protects you.

- **unsafe1** — `unsafe.Offsetof` for a field's byte offset (no reflection)
- **unsafe2** — zero-copy `[]byte` ⇄ `string` via `unsafe.String`/`SliceData`

## Resources

- [unsafe package](https://pkg.go.dev/unsafe)
- [Go spec: Package unsafe](https://go.dev/ref/spec#Package_unsafe)

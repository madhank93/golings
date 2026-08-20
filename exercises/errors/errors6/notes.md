## errors6 — %v severs the chain

```go
level1 := fmt.Errorf("read file: %w", ErrNotFound)
level2 := fmt.Errorf("load config: %w", level1) // was %v — the break
level3 := fmt.Errorf("startup: %w", level2)
```

**Why it works**

- `errors.Is` walks the wrap chain link by link. Every layer must use `%w` to
  keep a link; the one `%v` in the middle flattened `level1` into plain text, so
  from `level2` upward there was nothing left to unwrap.

**Under the hood**

- `%w` produces an error with an `Unwrap() error` method pointing at the
  original; `%v` calls `Error()` and embeds the **string**. Both messages read
  identically — `"startup: load config: read file: not found"` — which is what
  makes this bug so quiet. Only a test that asserts `errors.Is` finds it.

**Common mistake**

- Assuming one `%w` anywhere is enough. The chain is only as strong as its
  weakest layer: a single `%v` between the sentinel and the caller hides it
  completely.

**Key detail:** `fmt.Errorf` accepts multiple `%w` verbs (Go 1.20), so a layer can
wrap two causes at once. And when you *want* to hide an internal error from
callers — not leaking a database driver's type into your API — `%v` is the
deliberate choice.

**See also:** errors2 (wrapping basics) · errors5 (`Join`) · errors3 (`As`
walks the same chain) · the [chapter](../README.md)

**References**

- pkg.go.dev — fmt.Errorf: https://pkg.go.dev/fmt#Errorf ·
  errors.Unwrap: https://pkg.go.dev/errors#Unwrap
- Go blog — Working with Errors in Go 1.13: https://go.dev/blog/go1.13-errors

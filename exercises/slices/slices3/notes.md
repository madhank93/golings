## slices3 — grow a slice with append

```go
func addPeter(names []string) []string {
    return append(names, "Peter")
}
```

**Why it works**

- `append(names, "Peter")` returns a slice with the element added at the end. The
  broken `append()` gave it nothing to append.

**Nuance — you must use the return value.** `append` may allocate a **new**
backing array (when capacity runs out), so `names = append(names, x)` is the
rule. Ignoring the result — `append(names, x)` on its own — is a classic bug that
silently drops the new element.

**References**

- Go by Example — Slices: https://gobyexample.com/slices
- The Go Blog — Slices: usage and internals: https://go.dev/blog/slices-intro

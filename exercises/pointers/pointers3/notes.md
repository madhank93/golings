## pointers3 — value copy vs. pointer

The two functions differ only in the parameter type:

```go
func incValue(c Counter)   { c.N++ } // c is a COPY
func incPointer(c *Counter) { c.N++ } // c is the address of the original
```

**Why the test proves the point**

- `incValue(c)` passes the `Counter` **by value** — Go copies the whole struct
  into the parameter, so `c.N++` bumps the copy and the caller's `c.N` stays
  `0`.
- `incPointer(&c)` passes the **address**, so `c.N++` writes through to the
  caller's struct and `c.N` becomes `1`.

**Nuance — Go is *always* pass-by-value.** Even the pointer case is
pass-by-value: Go copies the **pointer** (the address), and both copies point at
the same struct. There is no pass-by-reference in Go. As Dave Cheney puts it:

> Go does not have reference variables, so Go does not have pass-by-reference
> function call semantics.

To mutate the caller's data, pass a pointer explicitly; to guarantee you
**can't**, pass by value. Copying a large struct every call is also wasteful —
another reason methods on big structs often take a pointer receiver.

**References**

- Dave Cheney — There is no pass-by-reference in Go:
  https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
- Effective Go — Pointers vs. Values: https://go.dev/doc/effective_go#pointers_vs_values

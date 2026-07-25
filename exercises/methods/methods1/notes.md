## methods1 — value vs pointer receivers

```go
func (r Rectangle) Area() float64 { return r.Width * r.Height } // reads a copy
func (r *Rectangle) Scale(factor float64) {                     // mutates original
    r.Width *= factor
    r.Height *= factor
}
```

**Why it works**

- `Area` only reads, so a **value receiver** `(r Rectangle)` — a copy — is enough.
- `Scale` must change the caller's rectangle, so it needs a **pointer receiver**
  `(r *Rectangle)`. A value receiver would scale a throwaway copy.

**Nuance:** Go auto-takes the address, so `r.Scale(2)` works on an addressable
`r` without writing `(&r).Scale(2)`. Rule of thumb: **be consistent** — if any
method needs a pointer receiver, give them all pointer receivers.

**References**

- A Tour of Go — Methods: https://go.dev/tour/methods/1
- Go by Example — Methods: https://gobyexample.com/methods
- Effective Go — Pointers vs. Values: https://go.dev/doc/effective_go#pointers_vs_values

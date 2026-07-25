## pointers2 — mutate a struct through a `*Account`

The solution:

```go
func deposit(a *Account, amount int) {
    a.Balance += amount
}
```

**Why it works**

- `a` is an `*Account` — the **address** of the caller's struct. Writing through
  it mutates the original `acc`, which is why the balance the test sees changes.
- Note what you did **not** have to write: `(*a).Balance`. For a pointer to a
  struct, a field selector **auto-dereferences** — `a.Balance` is shorthand for
  `(*a).Balance`. The Go spec states it plainly:

  > Selectors automatically dereference pointers to structs. If `x` is a pointer
  > to a struct, `x.y` is shorthand for `(*x).y`.

**Nuance:** in `pointers1` you had to spell out `*p` because Go does **not**
auto-dereference a `*int`. Here, because `a` points to a **struct**,
`a.Balance` just works. The auto-deref shorthand is a struct-only convenience.

**References**

- Go spec — Selectors: https://go.dev/ref/spec#Selectors
- Effective Go — Pointers vs. Values: https://go.dev/doc/effective_go#pointers_vs_values

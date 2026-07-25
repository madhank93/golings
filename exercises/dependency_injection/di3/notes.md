## di3 — constructor injection

```go
type Store interface {
    Save(name string)
    Count() int
}
type Greeter struct{ store Store } // dependency held as a field

func (g Greeter) Greet(name string) string {
    g.store.Save(name)
    return "Hi, " + name
}
```

**Why it works**

- `Greeter` holds its dependency (`Store`) as an **interface field**, set when the
  struct is built (`Greeter{store: store}`). `Greet` delegates to it. The test
  injects a `*memStore` and then asserts the store was actually used.

**Nuance:** this is the struct-level version of di1/di2 — inject collaborators
through the constructor/fields so nothing reaches for a global. Depend on the
**interface** (`Store`), not a concrete type, so production and test
implementations are interchangeable. This is Go's answer to "DI frameworks":
plain interfaces and struct fields.

**References**

- Effective Go — Interfaces: https://go.dev/doc/effective_go#interfaces

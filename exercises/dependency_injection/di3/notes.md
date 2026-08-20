## di3 — constructor injection

```go
type Greeter struct {
    store Store // the field is the INTERFACE, not a concrete type
}

func (g Greeter) Greet(name string) string {
    g.store.Save(name)
    return "Hi, " + name
}
```

**Why it works**

- The dependency is supplied when the `Greeter` is built and held as an
  interface field, so every method can use it. The test constructs
  `Greeter{store: &memStore{}}`; production passes the real store. One code
  path, two collaborators.

**Under the hood**

- The substitution hinges on the field's **type**. Declaring `store *postgres.DB`
  would compile and buy nothing — no other implementation could be assigned. The
  interface is what makes the seam.

**Common mistake**

- A constructor that reaches out instead of receiving: `NewGreeter()` calling
  `postgres.New(...)` internally. That is the original problem wearing a
  constructor's clothes — the dependency is still hard-coded, just one level
  down.

**Key detail:** define `Store` in the package that **consumes** it, listing only
the methods that package uses. "Accept interfaces, return structs" — the
implementation returns its concrete type and is free to have fifty other methods
nobody depends on.

**See also:** di1 · di2 · mock1 / mock2 (the doubles that get injected here) ·
interfaces1 (implicit satisfaction) · the [chapter](../README.md)

**References**

- Go Code Review Comments — interfaces: https://go.dev/wiki/CodeReviewComments#interfaces
- Effective Go — Interfaces: https://go.dev/doc/effective_go#interfaces

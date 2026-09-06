## structs2 — embedding promotes fields

```go
type ContactDetails struct {
    phone string
}

type Person struct {
    name string
    age  int
    ContactDetails // embedded — no field name
}

person.phone // promoted: shorthand for person.ContactDetails.phone
```

**Why it works**

- An embedded field is declared by **type only**. Its fields and methods are
  *promoted* onto the outer struct, so `person.phone` resolves without naming
  the inner struct.

**Under the hood**

- Promotion is a name-lookup rule, not inheritance. The embedded value is a real
  field whose name **is its type name** — which is why the literal still says
  `ContactDetails: ContactDetails{…}`. Selector lookup goes shallowest-first, so
  a `phone` field declared directly on `Person` would win, with the inner one
  still reachable at `person.ContactDetails.phone`.

**Common mistake**

- Expecting subtyping. A `Person` is **not** a `ContactDetails` and cannot be
  passed where one is expected; nothing is overridden and nothing is virtual.
  Embedding an *interface*, by contrast, does make the outer type satisfy it.

**Key detail:** promoted **methods** are the reason embedding matters — a struct
embedding `sync.Mutex` gets `Lock`/`Unlock` as its own methods, which is how
`sync1`-style types are usually written.

**See also:** structs1 · structs3 · interfaces4 (embedding interfaces) ·
the [chapter](../README.md)

**References**

- Go spec — Selectors (promotion rules): https://go.dev/ref/spec#Selectors
- Effective Go — Embedding: https://go.dev/doc/effective_go#embedding

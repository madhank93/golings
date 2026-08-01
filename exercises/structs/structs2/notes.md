## structs2 — embed a struct

```go
type ContactDetails struct{ phone string }

type Person struct {
    name string
    age  int
    ContactDetails // embedded (no field name)
}
person.phone // promoted from ContactDetails
```

**Why it works**

- Writing the type **without a field name** embeds it. The inner struct's fields
  are **promoted**, so `person.phone` works directly instead of
  `person.contact.phone`.

**Key detail:** embedding is Go's **composition**, not inheritance. `Person` gains
`ContactDetails`' fields *and* methods, but it's not a subtype of it. Set an
embedded field in a literal by its type name: `ContactDetails: ContactDetails{...}`.

**References**

- Go by Example — Struct Embedding: https://gobyexample.com/struct-embedding
- Effective Go — Embedding: https://go.dev/doc/effective_go#embedding

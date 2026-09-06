## structs3 — attach a method

```go
func (p Person) FullName() string {
    return p.firstName + " " + p.lastName
}
```

**Why it works**

- The receiver `(p Person)` binds the function to the type, so it can be called
  as `person.FullName()`. The body reads the two fields and joins them with a
  single space — the test pins the exact string, which a missing space would
  fail.

**Common mistake**

- Building strings this way in a loop. Each `+` allocates a new string, so
  concatenating n pieces copies O(n²) bytes. Two operands is fine; a loop wants
  `strings.Builder` or `strings.Join`.

**Key detail:** a **value receiver** is right here because the method only reads.
Use a pointer receiver when the method mutates the struct or the struct is large
enough that copying costs — and then use pointer receivers consistently across
the type, because the choice also decides what satisfies an interface.

**See also:** methods1 (value vs pointer receivers) · interfaces2 (`String()`) ·
strings1 (building strings) · the [chapter](../README.md)

**References**

- Go spec — Method declarations: https://go.dev/ref/spec#Method_declarations
- Effective Go — Pointers vs. Values: https://go.dev/doc/effective_go#pointers_vs_values

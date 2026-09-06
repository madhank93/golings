## structs1 — declare the fields

```go
type Person struct {
    name string
    age  int
}

person := Person{name: "John", age: 42}
```

**Why it works**

- A struct type is its field list. The test constructs `Person{name: …, age: …}`,
  so those two fields, with those types, have to exist on the type.

**Common mistake**

- Constructing positionally — `Person{"John", 42}`. It compiles today and breaks
  silently the moment a field is added or reordered. Keyed literals are the
  house style for exactly that reason.

**Key detail:** the zero value of a struct is every field's zero value, so
`var p Person` is immediately usable (`""`, `0`) with no constructor. Field names
starting lowercase are package-private — the same capital-letter rule that
governs everything else, and what decides whether `encoding/json` can see them.

**See also:** structs2 (embedding) · structs3 (methods) · variables3 (zero
values) · the [chapter](../README.md)

**References**

- Go spec — Struct types: https://go.dev/ref/spec#Struct_types
- Go by Example — Structs: https://gobyexample.com/structs

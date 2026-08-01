## structs1 — define a struct's fields

```go
type Person struct {
    name string
    age  int
}
person := Person{name: "John", age: 42}
```

**Why it works**

- A `struct` groups named fields, each with its own type. The empty `struct {}`
  had no `name`/`age`, so `Person{name: ..., age: ...}` wouldn't compile.

**Key detail:** prefer **field-name** literals (`Person{name: "John", age: 42}`) over
positional ones (`Person{"John", 42}`) — named literals survive field reordering
and new fields without silently breaking. A struct's zero value has every field
at *its* zero value.

**References**

- Go by Example — Structs: https://gobyexample.com/structs

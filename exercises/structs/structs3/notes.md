## structs3 — attach a method to a struct

```go
func (p Person) FullName() string {
    return p.firstName + " " + p.lastName
}
```

**Why it works**

- The `(p Person)` **receiver** binds the function to `Person`, so you can call
  `person.FullName()`. It returns the two names joined by a single space.

**Key detail:** this is exactly the kind of bug a test catches but the compiler
can't — `p.firstName + p.lastName` (no space) still compiles and returns
`"MaurícioAntunes"`. The receiver here is a **value** (`p Person`), a read-only
copy, which is right since `FullName` only reads.

**References**

- Go by Example — Methods: https://gobyexample.com/methods

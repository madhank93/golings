## generics1 — a type parameter

```go
func print[T any](value T) {
    fmt.Println(value)
}
print("Hello, World!") // T inferred as string
print(42)              // T inferred as int
```

**Why it works**

- `[T any]` declares a **type parameter**. One function body now works for every
  type; `any` is the constraint that permits all of them. The broken `print(value)`
  had no type for its parameter.

**Nuance:** the compiler **infers** `T` from the argument, so you rarely write
`print[int](42)` explicitly. Generics let you write one function instead of one
per type — without giving up compile-time type safety (unlike `any` + type
assertions).

**References**

- A Tour of Go — Generics: https://go.dev/tour/generics/1
- Go by Example — Generics: https://gobyexample.com/generics

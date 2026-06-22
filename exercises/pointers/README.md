# Pointers

A pointer holds the memory address of a value. `&x` takes an address; `*p`
dereferences. Pointers let functions mutate the caller's data and avoid copying
large structs.

- **pointers1** — read/write an int through `*p`
- **pointers2** — mutate a struct via `*Struct`
- **pointers3** — value copy vs pointer: only a pointer reaches the caller

## Resources

- [A Tour of Go: Pointers](https://go.dev/tour/moretypes/1)
- [Go by Example: Pointers](https://gobyexample.com/pointers)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)

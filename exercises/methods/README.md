# Methods

A method is a function with a receiver. Choose the receiver type deliberately:

- **value receiver** `(r T)` — operates on a copy; use for read-only behavior
- **pointer receiver** `(r *T)` — can mutate the original; use to change state or
  avoid copying large values

Methods can be defined on any named type you declare in your package, not only
structs.

- **methods1** — value (`Area`) vs pointer (`Scale`) receivers
- **methods2** — a method on a named non-struct type (`Celsius`)

## Resources

- [A Tour of Go: Methods](https://go.dev/tour/methods/1)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go by Example: Methods](https://gobyexample.com/methods)

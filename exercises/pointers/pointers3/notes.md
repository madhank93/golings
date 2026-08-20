## pointers3 — value copies, pointer reaches

```go
func incValue(c Counter)   { c.N++ } // mutates the copy — caller unaffected
func incPointer(c *Counter) { c.N++ } // mutates the caller's Counter
```

**Why it works**

- Go passes everything by value. `incValue` receives a copy of the struct, so its
  increment dies with the call. `incPointer` receives an address, so its
  increment lands on the caller's memory. The test asserts both outcomes.

**Under the hood**

- Slices, maps and channels appear to break this rule — a function can change
  their contents without a pointer. They do not: what is copied is a small
  header pointing at shared data. The header is a copy, the data is shared,
  which is exactly why `append` inside a function does not extend the caller's
  slice (it changes the copy's length).

**Common mistake**

- Assuming a method with a value receiver can "just" mutate. It cannot, for the
  same reason — and a `*T` method called on a map entry
  (`m["k"].Scale(2)`) does not compile at all, because a map entry has no
  address.

**Key detail:** the rule of thumb — pointer to **mutate**, pointer for **large**
structs, pointer when the value must not be copied (anything holding a
`sync.Mutex`). Value otherwise.

**See also:** pointers2 · methods1 (receivers) · slices3 (`append` and headers) ·
sync1 (never copy a mutex) · the [chapter](../README.md)

**References**

- Go FAQ — When are function parameters passed by value?: https://go.dev/doc/faq#pass_by_value
- Effective Go — Pointers vs. Values: https://go.dev/doc/effective_go#pointers_vs_values

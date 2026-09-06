## pointers2 — mutate a struct through a pointer

```go
func deposit(a *Account, amount int) {
    a.Balance += amount
}

acc := Account{Balance: 100}
deposit(&acc, 50) // acc.Balance == 150
```

**Why it works**

- `*Account` gives the function the caller's account rather than a copy, so
  `a.Balance += amount` updates the original.

**Under the hood**

- `a.Balance` on a pointer is **automatic dereferencing**: the compiler rewrites
  it as `(*a).Balance`. The explicit form is legal and never written. The same
  convenience runs the other way for method calls — `acc.Withdraw()` on an
  addressable value becomes `(&acc).Withdraw()`.

**Common mistake**

- Taking a pointer reflexively "to avoid the copy". For a small struct the copy
  is cheaper than the heap allocation and indirection a pointer can cause. Take
  a pointer when you need to **mutate**, or when the struct is genuinely large.

**Key detail:** returning `&local` from a function is safe in Go. Escape analysis
moves any value whose address outlives the call to the heap — check with
`go build -gcflags='-m'`, which prints `moved to heap: …`.

**See also:** pointers3 (the copy that cannot mutate) · methods1 (pointer
receivers) · structs1 · the [chapter](../README.md)

**References**

- Go spec — Selectors: https://go.dev/ref/spec#Selectors
- Go FAQ — stack or heap: https://go.dev/doc/faq#stack_or_heap

## pointers1 — read & write through a `*int`

The solution is one line:

```go
func double(p *int) {
    *p = *p * 2
}
```

**Why it works**

- `p` has type `*int` — it holds the **address** of the caller's `n`, not a
  copy of the value. The test calls `double(&n)`, and `&` is the operator that
  takes that address.
- `*p` is the **dereference**. On the **right** of `=` it *reads* the value at
  the address; on the **left** it *writes* to it. So `*p = *p * 2` reads `21`
  and stores `42` back into the very same `n` the caller owns — **no `return`
  needed**.

**Key detail:** for a primitive like `*int` you must write the dereference (`*p`)
explicitly every time you touch the value. Forgetting the `*` and writing
`p = p * 2` is a type error — you can't multiply an address.

**References**

- A Tour of Go — Pointers: https://go.dev/tour/moretypes/1
- Go by Example — Pointers: https://gobyexample.com/pointers

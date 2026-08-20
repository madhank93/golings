## variables4 — a new block can shadow a name

```go
x := "TEN" // outer x is a string
if true {
    x := 1 // NEW x, an int, local to this block
    fmt.Println(x + 1)
}
fmt.Println(x) // still "TEN"
```

**Why it works**

- The broken code wrote `x = 1`, assigning an `int` to the string `x` → type
  error. Using `x :=` inside the `if` **declares a second `x`** scoped to that
  block, so the types never clash.

**Under the hood**

- Every `{ … }` opens a scope, and a name resolves to the innermost declaration
  visible at that point. The inner `x` is a different variable at a different
  address; the outer one is untouched and comes back into view when the block
  ends. A variable's type is fixed at declaration and can never change, which is
  why assignment could not have worked here.

**Common mistake**

- Shadowing by accident, usually with `err` or a result variable inside an `if`
  or a loop: you assign the inner copy, the outer one keeps its old value, and
  the fix "does nothing". If you meant to update the outer variable, use `=`.

**Key detail:** deliberate shadowing is idiomatic — `if err := f(); err != nil`
keeps `err` from leaking into the rest of the function. `go vet`'s shadow
analyzer exists for the accidental kind.

**See also:** variables2 (`:=` rules) · if2 (statement-scoped variables) ·
the [chapter](../README.md)

**References**

- Go spec — Declarations and scope: https://go.dev/ref/spec#Declarations_and_scope
- Go by Example — Variables: https://gobyexample.com/variables

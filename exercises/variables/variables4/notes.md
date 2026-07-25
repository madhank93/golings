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

**Nuance — shadowing:** `:=` in an inner block creates a *new* variable that
hides the outer one; changes to it don't touch the outer `x`. This is a common
source of bugs — `go vet`'s shadow check exists precisely for it.

**References**

- Go by Example — Variables: https://gobyexample.com/variables
- Go spec — Declarations and scope: https://go.dev/ref/spec#Declarations_and_scope

## variables2 — declare before you assign

```go
x := 5
fmt.Printf("x has the value %d", x)
```

**Why it works**

- `x = 5` assigns to an `x` that was never declared → compile error. `:=`
  **declares and assigns** in one step, so `x` now exists.

**Common mistake**

- Reaching for `:=` at package level. A file's top level takes declarations
  only, so package variables are always `var x = 5`. `:=` is a function-body
  form.

**Key detail:** `:=` must introduce **at least one new variable** on its left. That
is why `v, err := f()` followed by `w, err := g()` is legal — `w` is new, `err`
is merely assigned. When no name is new, drop the colon.

**See also:** variables1 (the `var` forms) · variables4 (when `:=` shadows
instead of assigning) · the [chapter](../README.md)

**References**

- Go spec — Short variable declarations: https://go.dev/ref/spec#Short_variable_declarations
- Go by Example — Variables: https://gobyexample.com/variables

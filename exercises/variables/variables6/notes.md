## variables6 — constants can't be reassigned

```go
var x = 10 // was: const x = 10
func main() {
    fmt.Println(x)
    x = x + 1 // now legal — x is a variable
    fmt.Println(x)
}
```

**Why it works**

- The code reassigns `x` with `x = x + 1`. A `const` is immutable, so that line
  won't compile. Changing the declaration to `var` makes `x` mutable.

**Under the hood**

- There is nothing to assign *to*: a constant has no storage, only a value the
  compiler pastes in wherever the name appears. So this is not a permission
  check that could be overridden — the target simply does not exist.

**Common mistake**

- Reaching for `const` as "the safe default" and then discovering the value has
  to change. The choice is about the binding, not about caution: a value you
  ever reassign is a `var`, full stop.

**Key detail:** package-level `var` is initialised before `main` runs, in
dependency order, followed by any `init()` functions — so a `var` can be
computed from other package state where a `const` cannot.

**See also:** variables5 (constants need a value) · variables1 (declaration
forms) · the [chapter](../README.md)

**References**

- Go spec — Constant declarations: https://go.dev/ref/spec#Constant_declarations
- Go spec — Package initialization: https://go.dev/ref/spec#Package_initialization

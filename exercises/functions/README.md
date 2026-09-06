# Functions

A function in Go declares its name, its parameters *with their types*, and what
it returns. All three are checked at compile time and none of them can be
inferred from the call site, which is why all four exercises here fail to
compile rather than misbehave at run time — a missing type, a missing argument,
a missing return type.

The bigger idea underneath the syntax is that **functions are values**. A
function has a type (`func(int) string`), can be stored in a variable, passed to
another function, returned from one, and — with a receiver — attached to a type
as a method. Everything from `sort.Slice` to `http.HandlerFunc` to the
range-over-func iterators of Go 1.23 rests on that.

## 1. The declaration

```go
func addNumbers(a int, b int) int {
    return a + b
}

func addNumbers(a, b int) int { … }   // same thing: shared type, written once
```

Parameter types are mandatory (`functions2` omits one) and go *after* the name —
the same order as in variable declarations, and the reason Go's declaration
syntax reads left to right instead of spiralling like C's.

Order in the file does not matter: `main` may call a function declared below it
(`functions1`), because the compiler reads the whole package before checking
bodies. There are no forward declarations and no header files.

Naming is one word, mixedCaps, not snake_case — `callMe`, not `call_me`. And
the case of the first letter is not style but **access control**: an
uppercase name is exported from the package, a lowercase one is not.

## 2. Arity is exact

```go
func call_me(num int) { … }

call_me()      // compile error: not enough arguments
call_me(1, 2)  // compile error: too many
```

No default parameters, no optional arguments, no overloading. `functions3` is
this error. When a function genuinely needs optional configuration, Go reaches
for a struct of options or the functional-options pattern (`func(*Server)`)
rather than adding parameters with defaults.

The one flexible form is the variadic parameter, `nums ...int`, which the next
chapter covers.

## 3. Returns: declared, and often plural

```go
func half(n int) int          { return n / 2 }
func divide(a, b int) (int, error) { … }   // the Go signature you will write most
func swap(a, b string) (string, string) { return b, a }
```

If a function returns something, the type must be in the signature
(`functions4` omits it, so `return a + b` has nowhere to go). Multiple returns
are ordinary here, and they are why Go needs no exceptions: the second value
carries the error, and the caller must acknowledge it.

**Named results** are allowed:

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return          // "naked" return: returns the named values
}
```

They are worth using when the names document what the two `int`s mean, or when
a `defer` needs to modify the result before it leaves (`defer1`). Naked returns
in a long function are a readability trap — the values being returned are
nowhere near the `return`.

## 4. Functions as values

```go
var op func(int, int) int = addNumbers   // a variable of function type
op(2, 3)                                  // 5

func apply(nums []int, f func(int) int) []int {   // taken as a parameter
    out := make([]int, len(nums))
    for i, n := range nums {
        out[i] = f(n)
    }
    return out
}
```

The type of a function is its signature — parameter types and result types, not
its name. That single fact is what makes `sort.Slice(s, func(i, j int) bool)`,
`http.HandleFunc(path, handler)`, and dependency injection through a function
parameter all work with no interfaces involved. Anonymous functions and closures
build directly on it, in the `anonymous_functions` chapter.

## Gotchas

- **Everything is passed by value.** A struct argument is copied; to mutate the
  caller's value, take a pointer. Slices and maps *look* like exceptions because
  the copied header points at the same backing data (`pointers3` makes this
  concrete).
- **An unused parameter is fine; an unused local variable is not.** Name a
  parameter `_` when you must satisfy a signature you do not use.
- **A function with results must return on every path** — the compiler rejects
  "missing return", including at the end of a `for` with no exit.
- **`return` in a function with named results returns the current values**, not
  the zero values.
- **Exported means uppercase.** Renaming `Handler` to `handler` silently removes
  it from every other package's view.
- **`call_me` compiles but is not Go.** Linters (and reviewers) will say so.

## The exercises

- **functions1** — define the function `main` calls; order in the file does not
  matter.
- **functions2** — give the parameter its type.
- **functions3** — pass the argument the function declares.
- **functions4** — declare the return type for a function that returns a value.

## Source references

- [Go spec: Function declarations](https://go.dev/ref/spec#Function_declarations) ·
  [Function types](https://go.dev/ref/spec#Function_types) ·
  [Return statements](https://go.dev/ref/spec#Return_statements)
- [Effective Go: Functions](https://go.dev/doc/effective_go#functions) — multiple
  returns and named results, with the rationale
- [Go Code Review Comments: named result parameters](https://go.dev/wiki/CodeReviewComments#named-result-parameters)
- [A Tour of Go: Functions](https://go.dev/tour/basics/4)

**Next: [more_functions](../more_functions/) →** — a function that calls itself,
and one that takes as many arguments as you like.

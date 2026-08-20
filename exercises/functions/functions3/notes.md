## functions3 — pass the argument the signature declares

```go
func main() {
    call_me(5)
}

func call_me(num int) { … }
```

**Why it works**

- `call_me` takes one `int`, so calling it with none is
  `not enough arguments in call to call_me`. Supplying a value satisfies the
  signature.

**Common mistake**

- Looking for a default value — `func call_me(num int = 10)`. Go has no default
  parameters, no optional arguments, and no overloading: one name, one
  signature, exact arity.

**Key detail:** when a function genuinely needs optional configuration, Go uses a
config struct or the functional-options pattern (`func(*Server)`) rather than
growing the parameter list. The one built-in flexibility is a variadic
parameter (`nums ...int`), which accepts zero or more values — morefn2.

**See also:** functions2 (parameter types) · morefn2 (variadic) ·
di1 (passing collaborators in) · the [chapter](../README.md)

**References**

- Go spec — Calls: https://go.dev/ref/spec#Calls
- Go by Example — Functions: https://gobyexample.com/functions

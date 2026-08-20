## primitive_types2 — declare a string before you format it

```go
who := "Gopher"
fmt.Printf("Hello, %s\n", who)
```

**Why it works**

- `who` was used in `Printf` without ever being declared. Declaring it with
  `who := "Gopher"` gives `%s` something to substitute.

**Common mistake**

- Assuming a verb mismatch is a compile error. It is not — `fmt` takes `...any`,
  so `%d` with a string compiles and prints `%!d(string=Gopher)` at run time.
  `go vet` catches the common cases, which is why this repo's lint step runs it.

**Key detail:** match the verb to the type — `%s` strings, `%d` integers, `%t`
bools, `%f` floats, `%q` a quoted string (the quickest way to see stray
whitespace), `%v` anything, `%T` the value's type.

**See also:** primitive_types3 (two verbs, two arguments) · strings1 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — fmt: https://pkg.go.dev/fmt
- Go by Example — String Formatting: https://gobyexample.com/string-formatting

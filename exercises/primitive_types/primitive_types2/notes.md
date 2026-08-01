## primitive_types2 — declare a string before you format it

```go
who := "Gopher"
fmt.Printf("Hello, %s\n", who)
```

**Why it works**

- `who` was used in `Printf` without ever being declared. Declaring it with
  `who := "Gopher"` gives `%s` something to substitute.

**Key detail:** `%s` is the verb for strings. Match the verb to the type — `%d` for
integers, `%v` for "any value, default format", `%q` for a quoted string. A
mismatch compiles but prints a `%!s(...)` error marker at run time.

**References**

- Go by Example — String Formatting: https://gobyexample.com/string-formatting

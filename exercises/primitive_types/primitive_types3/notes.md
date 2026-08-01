## primitive_types3 — every format argument must exist

```go
who := "Gopher"
country := "India"
fmt.Printf("Hello, I am %s and live in %s\n", who, country)
```

**Why it works**

- The format string has two `%s` verbs, so `Printf` needs two declared
  arguments. Declaring both `who` and `country` satisfies them.

**Key detail:** `Printf` takes its arguments **left to right**, one per verb. Too few
arguments prints a `%!s(MISSING)` marker; extra ones print `%!(EXTRA ...)`. The
count and order matter.

**References**

- Go by Example — String Formatting: https://gobyexample.com/string-formatting

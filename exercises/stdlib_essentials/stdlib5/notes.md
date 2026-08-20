## stdlib5 — strconv

```go
func parseAndDouble(s string) (int, error) {
    n, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    return n * 2, nil
}
```

**Why it works**

- `strconv.Atoi` parses a string as an `int` and returns `(int, error)`. Checking
  the error is what separates "the user typed nonsense" from "the user typed 0" —
  both would otherwise look like the zero value.

**Under the hood**

- The failure is a `*strconv.NumError` carrying the function name, the input, and
  the underlying cause (`ErrSyntax` or `ErrRange`), so
  `errors.Is(err, strconv.ErrRange)` distinguishes "not a number" from "too big
  for an int".

**Common mistake**

- Using `string(65)` to render a number. That is a **rune conversion** producing
  `"A"`; `strconv.Itoa(65)` gives `"65"`. `go vet` flags the former.

**Key detail:** `strconv.Itoa`/`Atoi` are several times faster than
`fmt.Sprintf("%d", n)` / `fmt.Sscanf` and say exactly what they do. The package
also has `ParseFloat`, `ParseBool`, `ParseInt` (with base and bit size), and
`Quote`.

**See also:** errors1 (returning errors) · primitive_types5 (numeric types) ·
cli1 (parsing user input) · the [chapter](../README.md)

**References**

- pkg.go.dev — strconv: https://pkg.go.dev/strconv
- pkg.go.dev — strconv.NumError: https://pkg.go.dev/strconv#NumError

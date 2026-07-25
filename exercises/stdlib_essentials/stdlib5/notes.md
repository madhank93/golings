## stdlib5 — strconv

```go
n, err := strconv.Atoi(s) // string → int, with an error
if err != nil {
    return 0, err
}
```

**Why it works**

- `strconv.Atoi` parses a string into an `int`, returning an `error` when the text
  isn't a number (`"nope"`). Propagating that error lets the caller handle bad
  input.

**Nuance:** conversion between strings and numbers is **explicit and fallible** in
Go — there's no implicit coercion. `Atoi`/`Itoa` cover base-10 ints; use
`ParseInt`/`ParseFloat`/`FormatInt` for other bases, bit sizes, and floats.
Contrast with a numeric *type* conversion like `int(f)`, which never errors.

**References**

- pkg.go.dev — strconv: https://pkg.go.dev/strconv

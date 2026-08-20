## stdlib4 — the reference time layout

```go
func parseDate(s string) (time.Time, error) {
    return time.Parse("2006-01-02", s)
}
```

**Why it works**

- Go states time formats by **example**, not with `%Y-%m-%d` codes. The layout
  spells out how the reference instant would look, so `"2006-01-02"` means
  four-digit year, two-digit month, two-digit day.

**Under the hood**

- The reference is `Mon Jan 2 15:04:05 MST 2006` — which is
  `01/02 03:04:05PM '06 -0700`, the numbers 1 through 7 in order. That is the
  mnemonic: month 1, day 2, hour 3, minute 4, second 5, year 6, zone 7. `15` is
  the 24-hour form of hour 3.

**Common mistake**

- Writing `"YYYY-MM-DD"`. It is a valid layout — it just does not contain any of
  the reference components, so parsing fails with a confusing error and
  formatting emits the literal text.

**Key detail:** `time.Parse` assumes UTC when the layout has no zone; `ParseInLocation`
takes one explicitly. And never compare times with `==` — a `time.Time` carries
a monotonic reading and a location, so use `Equal`, `Before`, `After`.

**See also:** stdlib7 (a zero `time.Time` in JSON) · di2 (injecting a clock) ·
synctest2 (virtual time) · the [chapter](../README.md)

**References**

- pkg.go.dev — time.Parse: https://pkg.go.dev/time#Parse ·
  Layout constants: https://pkg.go.dev/time#pkg-constants
- Go by Example — Time Formatting: https://gobyexample.com/time-formatting-parsing

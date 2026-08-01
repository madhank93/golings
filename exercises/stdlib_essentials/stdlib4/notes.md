## stdlib4 — time layouts

```go
time.Parse("2006-01-02", s) // layout uses the reference date
```

**Why it works**

- Go doesn't use `YYYY-MM-DD` format codes. Instead you write the **reference
  date** — Mon Jan 2 15:04:05 MST 2006 — in the shape you want. `"2006-01-02"`
  therefore means "4-digit year - 2-digit month - 2-digit day".

**Key detail:** the magic numbers are a mnemonic: `01/02 03:04:05 PM '06 -0700` =
1,2,3,4,5,6,7. Use the **same** layout string for both `Parse` and `Format`. A
mismatch between the layout and the input is a run-time error, not a compile
error — so the layout must exactly describe the data.

**References**

- pkg.go.dev — time (Layout): https://pkg.go.dev/time#pkg-constants

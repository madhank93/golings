## files2 — streaming with bufio.Scanner

```go
func countLines(r io.Reader) int {
    scanner := bufio.NewScanner(r)
    n := 0
    for scanner.Scan() {
        n++
    }
    return n
}
```

**Why it works**

- `Scan()` advances to the next line and returns false when there is none, so the
  loop counts lines without ever holding more than one in memory. `scanner.Text()`
  returns the line, newline stripped.

**Under the hood**

- `Scan` returns false at EOF **or on error**, which is why real code checks
  `scanner.Err()` afterwards — without it, a read failure is indistinguishable
  from a short file. Note that this counting version does not, which is fine for
  a `strings.Reader` and would not be for a network stream.

**Common mistake**

- Hitting the 64 KB default token limit. A longer line stops the scan with
  `bufio.ErrTooLong` — raise it deliberately with `scanner.Buffer(buf, max)` for
  formats like JSONL.

**Key detail:** the split function is configurable — `scanner.Split(bufio.ScanWords)`
or `ScanRunes`, or your own `SplitFunc`. And `scanner.Bytes()` avoids the
per-line string allocation but is only valid until the next `Scan`.

**See also:** files1 (whole-file) · stdlib2 (`io.Reader`) · strings2 ·
logingest7 (parsing untrusted lines) · the [chapter](../README.md)

**References**

- pkg.go.dev — bufio.Scanner: https://pkg.go.dev/bufio#Scanner
- Go by Example — Line Filters: https://gobyexample.com/line-filters

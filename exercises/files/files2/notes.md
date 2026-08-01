## files2 — read line by line with bufio.Scanner

```go
scanner := bufio.NewScanner(r)
for scanner.Scan() { // advances one line at a time
    n++
}
```

**Why it works**

- `bufio.Scanner` wraps any `io.Reader` and, by default, splits on lines.
  `Scan()` returns `true` and buffers the next line each call, `false` at EOF —
  so the loop counts lines without loading the whole input.

**Key detail:** this streams, so it handles arbitrarily large inputs with constant
memory — unlike `ReadFile` + split. Read the line text with `scanner.Text()`;
change the split with `scanner.Split(bufio.ScanWords)`. After the loop, check
`scanner.Err()` — `Scan` returning `false` can mean EOF *or* an error.

**References**

- Go by Example — Line Filters: https://gobyexample.com/line-filters
- pkg.go.dev — bufio: https://pkg.go.dev/bufio

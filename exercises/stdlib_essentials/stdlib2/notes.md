## stdlib2 — io.Reader / io.Writer

```go
func readAll(r io.Reader) (string, error) {
    var buf bytes.Buffer
    if _, err := io.Copy(&buf, r); err != nil {
        return "", err
    }
    return buf.String(), nil
}
```

**Why it works**

- `io.Copy(dst, src)` streams bytes from **any** `io.Reader` to **any**
  `io.Writer`. A `strings.Reader` and a `bytes.Buffer` both satisfy those
  one-method interfaces, so `readAll` works without caring about the concrete
  types.

**Nuance:** `io.Reader`/`io.Writer` are the two most important interfaces in Go —
files, network sockets, HTTP bodies, buffers, and compressors all implement them,
so code written against them composes universally. Streaming with `io.Copy` also
avoids loading everything into memory at once.

**References**

- pkg.go.dev — io: https://pkg.go.dev/io

## stdlib2 — io.Copy

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

- `io.Copy(dst, src)` streams from any `io.Reader` to any `io.Writer` until the
  reader is exhausted. `bytes.Buffer` is both, so it can receive the copy and
  hand back a string.

**Under the hood**

- `Copy` uses a fixed 32 KB buffer, so memory stays flat no matter how large the
  source is — the difference between this and `io.ReadAll`, which holds the whole
  input. It also checks for `WriterTo`/`ReaderFrom` and lets those take over,
  which is how a file-to-socket copy becomes a kernel `sendfile`.

**Common mistake**

- Treating `io.EOF` as a failure. It is how a stream ends, and `Copy` already
  handles it — a non-nil error from `Copy` is a *real* problem.

**Key detail:** `io.Reader` and `io.Writer` are one method each, and everything
speaks them: files, sockets, buffers, compressors, hashers, HTTP bodies. Take
those interfaces in your own functions and they work with all of it — plus a
`strings.NewReader` in the test.

**See also:** files2 (streaming lines) · di1 (`io.Writer` as a dependency) ·
http1 (reading a response body) · the [chapter](../README.md)

**References**

- pkg.go.dev — io.Copy: https://pkg.go.dev/io#Copy
- pkg.go.dev — io.Reader: https://pkg.go.dev/io#Reader

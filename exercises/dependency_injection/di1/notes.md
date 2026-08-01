## di1 — inject an io.Writer

```go
func Greet(w io.Writer, name string) {
    fmt.Fprintf(w, "Hello, %s!", name)
}
var buf bytes.Buffer
Greet(&buf, "Go") // capture output in a test
```

**Why it works**

- Instead of hardcoding `fmt.Println` (which writes to stdout), `Greet` takes an
  `io.Writer`. In production you pass `os.Stdout`; in a test you pass a
  `bytes.Buffer` and assert on what was written.

**Key detail:** this is **dependency injection** at its simplest — pass a dependency in
rather than reaching for a global. Because `io.Writer` is a tiny interface, files,
buffers, network connections, and stdout all satisfy it, so the same function
works everywhere and is trivially testable.

**References**

- Learn Go with Tests — Dependency Injection: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

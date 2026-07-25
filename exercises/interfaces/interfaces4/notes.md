## interfaces4 — embed interfaces

```go
type Reader interface{ Read() string }
type Writer interface{ Write(s string) }

type ReadWriter interface {
    Reader
    Writer
}
```

**Why it works**

- `ReadWriter` embeds `Reader` and `Writer`, so its method set is the **union**:
  a type satisfies `ReadWriter` by having both `Read` and `Write`. `*Buffer` has
  both, so it fits.

**Nuance:** this is exactly how the standard library builds `io.ReadWriter` from
`io.Reader` + `io.Writer`. Compose small, single-method interfaces into larger
ones rather than declaring one big interface up front — smaller interfaces are
easier to satisfy and mock.

**References**

- io.ReadWriter: https://pkg.go.dev/io#ReadWriter
- Effective Go — Interfaces: https://go.dev/doc/effective_go#interfaces

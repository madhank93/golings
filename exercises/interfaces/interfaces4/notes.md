## interfaces4 — interfaces compose by embedding

```go
func (b *Buffer) Write(s string) { b.data += s }
func (b *Buffer) Read() string   { return b.data }

// ReadWriter embeds Reader + Writer; *Buffer satisfies it by having both.
```

**Why it works**

- An interface may embed other interfaces; the result is the **union** of their
  methods. Implement `Read` and `Write` and `*Buffer` satisfies `ReadWriter`
  automatically — again with nothing declared.

**Under the hood**

- Embedding is flattening at compile time, not delegation: `ReadWriter`'s method
  set is exactly `{Read, Write}`. This is how `io.ReadWriter`, `io.ReadCloser`
  and `io.ReadWriteSeeker` are built from the one-method `io.Reader` and
  `io.Writer`.

**Common mistake**

- Implementing the methods on `Buffer` but passing a `Buffer` value. With
  **pointer receivers** only `*Buffer` has those methods, so `useRW(b)` fails and
  `useRW(&b)` works — the method-set rule again, and the most common interface
  error message in Go.

**Key detail:** the same trick works on structs: embedding an interface *field*
in a struct gives the struct those methods by delegation, which is how test
fakes override one method of a large interface and leave the rest to panic.

**See also:** interfaces1 · methods1 (method sets) · structs2 (struct
embedding) · mock2 (fakes built this way) · the [chapter](../README.md)

**References**

- Go spec — Interface types (embedding): https://go.dev/ref/spec#Embedded_interfaces
- pkg.go.dev — io.ReadWriter: https://pkg.go.dev/io#ReadWriter

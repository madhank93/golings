## binary1 — a fixed header, big-endian

```go
var buf [HeaderSize]byte
if _, err := io.ReadFull(r, buf[:]); err != nil {
    return Header{}, err
}
h := Header{
    ID:    binary.BigEndian.Uint16(buf[0:2]),
    Flags: binary.BigEndian.Uint16(buf[2:4]),
    // …
}
```

**Why it works**

- A fixed header is a byte offset table. Each field has a known position and a
  known width, so decoding is slicing plus `binary.BigEndian.Uint16`. Nothing is
  self-describing, which is exactly why it is fast and exactly why an off-by-two
  offset gives you a plausible wrong number rather than an error.

**Under the hood**

- **Big-endian is network byte order** (RFC 1700), most significant byte first,
  and nearly every wire protocol uses it while nearly every CPU you own is
  little-endian. `binary.BigEndian` is a `struct{}` with methods, so the shifts
  are inlined — the abstraction costs nothing.

**Common mistake**

- Using `io.Read` instead of `io.ReadFull`. A single `Read` may legally return
  three bytes of a twelve-byte header, and code that ignores the count decodes
  whatever happened to be in the rest of the buffer. `ReadFull` also draws the
  distinction that matters: `io.EOF` for "nothing was there",
  `io.ErrUnexpectedEOF` for "it was truncated".

**Key detail:** `binary.Read(r, binary.BigEndian, &h)` fills the struct in one
call, at the price of reflection and of your struct layout becoming the wire
format. Explicit offsets survive the day the wire adds a field and your struct
does not.

**See also:** binary2 (the flag bits inside this header) · binary3 (framing) ·
stdlib_essentials2 (`io`) · the [chapter](../README.md)

**References**

- pkg.go.dev — encoding/binary: https://pkg.go.dev/encoding/binary
- RFC 1035 §4.1.1 (this header, named): https://www.rfc-editor.org/rfc/rfc1035#section-4.1.1

## binary3 — length-prefixed frames

```go
var lenBuf [4]byte
if _, err := io.ReadFull(r, lenBuf[:]); err != nil {
    return nil, err            // io.EOF here = the sender finished
}
n := binary.BigEndian.Uint32(lenBuf[:])
if n > MaxFrame {
    return nil, fmt.Errorf("%w: %d bytes", ErrFrameTooLarge, n)
}
payload := make([]byte, n)
if _, err := io.ReadFull(r, payload); err != nil {
    return nil, err            // io.ErrUnexpectedEOF = truncated
}
```

**Why it works**

- Same problem as tcp2 — where does a message end — with the other answer.
  A delimiter suits text; a length prefix suits binary, where any byte value can
  appear in the payload and no delimiter is safe. The reader learns the size
  before the content, so it reads exactly that much and stops.

**Under the hood**

- This is the frame shape of Kafka, of the BitTorrent peer protocol, of gRPC's
  HTTP/2 data frames and of TLS records. The prefix width is the design
  decision: four bytes caps a message at 4 GiB, two bytes at 64 KiB, and a
  varint trades a byte of overhead for no cap at all.

**Common mistake**

- `make([]byte, n)` before checking `n`. The length came from the network, so it
  is attacker-controlled: four bytes of `0xff` ask for a four-gigabyte
  allocation, and a handful of those is a denial of service that needs no
  bandwidth at all. Bound it first, always.

**Key detail:** the two EOFs mean different things and callers depend on it.
`io.EOF` before the prefix is a clean end of stream; `io.ErrUnexpectedEOF`
anywhere after it means the peer died mid-frame.

**See also:** tcp2 (delimiter framing) · binary1 (`io.ReadFull`) · errors2
(`%w`) · the [chapter](../README.md)

**References**

- pkg.go.dev — io.ReadFull: https://pkg.go.dev/io#ReadFull
- Kafka protocol — the size-delimited request format: https://kafka.apache.org/protocol.html

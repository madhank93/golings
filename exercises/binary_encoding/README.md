# Binary encoding

Text formats tell you what they contain; binary formats do not. A wire header
is a fixed run of bytes at known offsets, and everything you get back depends on
agreeing with the sender about **width**, **order** and **position**. Get one of
them wrong and you do not get an error — you get a plausible wrong number.

This chapter is `encoding/binary`, bit twiddling, and the two `io` interfaces
that make it safe: `io.ReadFull` for streams and `io.ReaderAt` for files.

## 1. Fixed headers and network byte order

```go
var buf [12]byte
if _, err := io.ReadFull(r, buf[:]); err != nil {
    return Header{}, err
}
id := binary.BigEndian.Uint16(buf[0:2])
```

`binary1`. **Big-endian is network byte order** — most significant byte first —
and effectively every wire protocol uses it while effectively every CPU you own
is little-endian. `binary.BigEndian` is a zero-size type whose methods compile
to shifts, so naming the endianness costs nothing at runtime and prevents the
one bug that never shows up in a hex dump you read left to right.

```
   offset  0     2     4     6     8    10    12
           ├─────┼─────┼─────┼─────┼─────┼─────┤
           │  ID │flags│  qd │  an │  ns │  ar │
           └─────┴─────┴─────┴─────┴─────┴─────┘
             u16   u16   u16   u16   u16   u16      all big-endian
```

`io.ReadFull` is not optional. A single `Read` may return three bytes of a
twelve-byte header, and it distinguishes the two ends: `io.EOF` for "nothing was
there", `io.ErrUnexpectedEOF` for "it was truncated".

`binary.Read(r, binary.BigEndian, &h)` fills a struct in one reflective call.
Convenient, slower, and it makes your Go struct layout part of the wire
contract — explicit offsets age better.

## 2. Bitfields: shift and mask

```go
Opcode: uint8(w >> 11 & 0xf),      // read:  shift down, mask to width
w |= uint16(f.Opcode&0xf) << 11    // write: mask to width, shift up
```

`binary2`. When a field is narrower than a byte, its address is a *shift* and a
*width*:

```
    bit 15                                          bit 0
      │                                                │
      ▼                                                ▼
     ┌──┬──────────────┬──┬──┬──┬──┬──────────┬────────────┐
     │QR│    opcode    │AA│TC│RD│RA│  z (3)   │ rcode (4)  │
     └──┴──────────────┴──┴──┴──┴──┴──────────┴────────────┘
      1        4         1  1  1  1     3            4       bits
```

The mask is the part people skip, because omitting it works until a value
overflows its field and corrupts the neighbour. Go has no C-style bitfield
syntax on purpose: C leaves the packing implementation-defined, and explicit
shifts are both portable and readable.

## 3. Length-prefixed framing

```go
n := binary.BigEndian.Uint32(lenBuf[:])
if n > MaxFrame {
    return nil, fmt.Errorf("%w: %d bytes", ErrFrameTooLarge, n)
}
payload := make([]byte, n)
_, err := io.ReadFull(r, payload)
```

`binary3`. The other answer to `net_tcp`'s framing question, and the one binary
protocols use — a delimiter is unsafe when any byte can appear in the payload.

The security note is the point of the exercise: **the length came off the wire**,
so `make([]byte, n)` before checking `n` turns four bytes of `0xff` into a
four-gigabyte allocation. Bound it first. This is a real denial-of-service class,
cheap to trigger and cheap to prevent.

## 4. Random access with `io.ReaderAt`

```go
offset := int64(n-1) * int64(pageSize)     // pages are 1-based
sr := io.NewSectionReader(ra, offset, int64(size))
_, err := io.ReadFull(sr, page)
```

`binary4`. `io.Reader` has a cursor; `io.ReaderAt` does not — it reads at an
absolute offset and is therefore safe to call from several goroutines at once.
That is why every random-access format (databases, archives, disk images) is
written against it.

`io.NewSectionReader` bounds a reader to one region, so a page decoder written
against `io.Reader` physically cannot read past its page.

## Gotchas

- **Always say the endianness.** `binary.BigEndian` for wire formats.
- **`io.ReadFull`, not `Read`.** A short read is an error, not fewer bytes.
- **Bound any length that came from outside** before allocating.
- **Mask every bitfield to its width**, on the way in and on the way out.
- **`io.EOF` vs `io.ErrUnexpectedEOF`** — clean end versus truncation. Callers
  depend on the difference.
- **`ReadAt` may return a short count with a nil error.** Wrap it.
- **Reserved bits are for a future version**; preserve them rather than zeroing.
- **`unsafe` struct casts are not a shortcut** — padding and host endianness
  make them wrong on the first machine that differs.

## The exercises

- **binary1** — decode (and re-encode) a twelve-byte big-endian header.
- **binary2** — unpack and pack the bitfields inside one 16-bit word.
- **binary3** — read length-prefixed frames, safely, off a stream.
- **binary4** — read a fixed-size page at an offset with `io.ReaderAt`.

## Where this goes next

The four shapes here are not invented. The header is a DNS message header (byox
dns-server), the flag word is its second field, length-prefixed frames are the
Kafka and BitTorrent peer protocols (25 and 19 stages), and 1-based pages after
a 100-byte file header is the SQLite database format (9 stages). byox's git
course adds zlib and SHA-1 on top of exactly this decoding.

## Source references

- [pkg.go.dev: encoding/binary](https://pkg.go.dev/encoding/binary) ·
  [io.ReadFull](https://pkg.go.dev/io#ReadFull) ·
  [io.ReaderAt](https://pkg.go.dev/io#ReaderAt)
- [RFC 1035 §4.1.1](https://www.rfc-editor.org/rfc/rfc1035#section-4.1.1) — the
  header and flag word this chapter borrows
- [SQLite file format](https://www.sqlite.org/fileformat.html) — pages, 1-based,
  with a 100-byte header on page 1

**Next: [process_control](../process_control/) →** — from bytes on a wire to
processes on a machine.

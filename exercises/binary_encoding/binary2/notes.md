## binary2 — bitfields: shift, then mask

```go
// read
Opcode: uint8(w >> 11 & 0xf),

// write
w |= uint16(f.Opcode&0xf) << 11
```

**Why it works**

- A bitfield is a **shift** (which bit does the field start at) and a **width**
  (how many bits it owns). Reading: shift the field down to bit 0, mask off
  everything above it. Writing: mask the value to its width, shift it into
  place, OR it in. Getting the mask wrong is silent — one field's overflow
  becomes another field's value.

**Under the hood**

- Go has no C-style bitfield syntax, deliberately: C leaves the layout
  implementation-defined, so the same struct packs differently across compilers.
  Explicit shifts are portable and the reader can see the wire format.

**Common mistake**

- `w >> 11` without `& 0xf`. It works for every value you happen to test with
  because the bits above are zero, and breaks the moment they are not.
- Precedence: in Go `>>` and `&` bind tighter than `==`, so `w>>15&1 == 1`
  parses as `((w>>15)&1) == 1`. Correct, and worth parenthesising anyway.

**Key detail:** reserved bits ("must be zero") are reserved for a *future*
version. Preserving them on a round trip is safer than zeroing them, and both
are better than assuming they are yours.

**See also:** binary1 (the header holding this word) · primitive_types3
(conversions) · the [chapter](../README.md)

**References**

- Go spec — arithmetic operators: https://go.dev/ref/spec#Arithmetic_operators
- RFC 1035 §4.1.1 (these exact bits): https://www.rfc-editor.org/rfc/rfc1035#section-4.1.1

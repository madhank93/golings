## tcp2 — framing: find the message inside the stream

```go
r := bufio.NewReader(conn)
for {
    line, err := r.ReadString('\n')
    if err != nil {
        return nil // the client closed; not a failure
    }
    cmd := strings.TrimSuffix(line, "\n")
    if cmd == "QUIT" {
        return nil
    }
    fmt.Fprintf(conn, "%s\n", strings.ToUpper(cmd))
}
```

**Why it works**

- TCP is a **byte stream**, not a message queue. One `Write` by the client can
  arrive as three `Read`s, and three `Write`s can arrive as one. The delimiter —
  here a newline — is the only thing that says where a message ends, and
  `bufio.Reader` keeps the leftovers between calls so a split message is
  reassembled rather than lost.

**Under the hood**

- Nagle's algorithm on the sender and delayed ACK on the receiver both coalesce
  small writes, so "one write, one read" holds on loopback right up until it
  does not. A framing bug reproduces on a real network and never on your laptop.

**Common mistake**

- Treating the read error as a failure. `io.EOF` from a connection means the peer
  closed cleanly, which is the normal end of a conversation — reporting it as an
  error makes every well-behaved client look like a fault.

**Key detail:** `bufio.Scanner` is the other option and caps a token at
`bufio.MaxScanTokenSize` (64 KiB) by default, silently stopping on anything
longer. `ReadString` has no such limit, which is better here and worse when the
peer is hostile.

**See also:** tcp1 (one exchange) · tcp3 (many connections) · binary3
(length-prefixed framing) · the [chapter](../README.md)

**References**

- pkg.go.dev — bufio.Reader: https://pkg.go.dev/bufio#Reader
- RFC 9293 §3.1 (TCP is a stream): https://www.rfc-editor.org/rfc/rfc9293

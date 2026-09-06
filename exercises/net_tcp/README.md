# Net (TCP)

`net/http` hides a socket. This chapter opens one. A TCP server in Go is four
calls — `Listen`, `Accept`, `Read`, `Write` — and everything else is a decision
about *framing* (where does one message end?) and *concurrency* (who waits for
whom?).

Those two decisions are the whole difference between a toy and a server, and
they are the same whether the protocol on top is Redis, HTTP, Kafka or one you
invent this afternoon.

## 1. Listen and accept

```go
ln, err := net.Listen("tcp", "127.0.0.1:0")   // :0 = any free port
defer ln.Close()

conn, err := ln.Accept()                      // blocks until someone connects
defer conn.Close()
```

`tcp1`. `net.Conn` is an `io.ReadWriteCloser` with two addresses attached, so
every reader and writer idiom you already know applies to it unchanged.

Two details the tests lean on:

- **Port 0** asks the kernel for a free port, and `ln.Addr()` reports which one.
  Hardcoding 8080 in a test is how a suite fails on a machine where something
  else is already listening.
- **Closing your side is a protocol message.** It is what the peer sees as EOF.
  A handler that returns without closing leaves the client blocked until its
  deadline.

```
   client                     server
     │   connect ───────────►  Accept() returns
     │   "PING\n" ──────────►  Read
     │   ◄────────── "PONG\n"  Write
     │   ◄────────── (EOF)     Close
```

## 2. Framing: TCP is a stream, not messages

```go
r := bufio.NewReader(conn)
for {
    line, err := r.ReadString('\n')
    if err != nil {
        return nil            // io.EOF: the peer closed cleanly
    }
    …
}
```

`tcp2`. This is the single most important idea in the chapter. TCP guarantees
**order** and **delivery**, and guarantees nothing about *boundaries*: one
`Write` of three commands can arrive as one `Read` or as seven. The protocol
must say where a message ends, and there are only two answers in practice —

| Framing | Looks like | Used by |
|---|---|---|
| **Delimiter** | `…\r\n` | HTTP headers, Redis RESP, SMTP |
| **Length prefix** | `[4-byte length][payload]` | Kafka, BitTorrent, gRPC/HTTP2 |

`bufio.Reader` is what makes delimiter framing correct: it keeps the bytes past
the delimiter for the next call, so a split message is reassembled rather than
lost. Length-prefixed framing is `binary_encoding`'s `binary3`.

## 3. The accept loop and its ending

```go
for {
    conn, err := ln.Accept()
    if err != nil {
        if errors.Is(err, net.ErrClosed) {
            return nil        // ln was closed: this is a shutdown
        }
        return err
    }
    handle(conn)
}
```

`tcp3`. Closing the listener is how a server stops, and the `Accept` error that
follows is a shutdown signal wearing an error's clothes. Match it with
`errors.Is(err, net.ErrClosed)` — never by comparing the message string, which
has been reworded before.

## 4. A goroutine per connection

```go
for {
    conn, err := ln.Accept()
    …
    wg.Add(1)
    go func() {
        defer wg.Done()
        defer conn.Close()
        handle(conn)
    }()
}
```

`tcp4`. The accept loop stays single-threaded; the work moves off it. A slow
client then costs one goroutine — a few kilobytes — instead of the whole server.
This is the shape of `http.Server`, and it is why blocking code in Go scales.

Two obligations come with it:

- **Shared state needs a lock.** Every handler touches the same counter here,
  and an unguarded `n++` is a data race. golings runs `go test -race` on every
  exercise, so this fails loudly rather than producing plausible numbers.
- **A `WaitGroup`, so shutdown means shutdown.** Without it `Serve` returns
  while handlers are still writing.

In production you would also *bound* the goroutines — a semaphore channel or a
worker pool — because "one goroutine per connection" is also "one goroutine per
attacker".

## Gotchas

- **Never assume one `Read` is one message.** Frame explicitly.
- **`io.EOF` from a connection is not a failure** — it is the peer saying it is
  finished.
- **Close the connection**, or the peer waits for its deadline.
- **`errors.Is(err, net.ErrClosed)`** for shutdown, not string matching.
- **Deadlines, not timeouts.** `conn.SetDeadline(time.Now().Add(d))` is
  absolute, applies to reads and writes already in flight, and must be reset per
  operation. Without one, a peer that stops sending pins a goroutine forever.
- **`bufio.Scanner` caps a token at 64 KiB** and stops silently past that.
- **Listen on `127.0.0.1`, not `:`** unless you mean to be reachable.

## The exercises

- **tcp1** — accept one connection, answer one line, close it.
- **tcp2** — keep the connection open and frame many commands out of the stream.
- **tcp3** — an accept loop that stops cleanly when the listener closes.
- **tcp4** — a goroutine per connection, and a lock around what they share.

## Where this goes next

This chapter is the floor under three byox courses. Redis (115 stages) opens
with Listen/Accept and RESP, a delimiter-framed protocol; the HTTP server course
(14 stages) is this loop plus request parsing; Kafka (25 stages) swaps the
delimiter for a length prefix. All three assume, from stage 1, that a byte
stream is not a message queue.

## Source references

- [pkg.go.dev: net](https://pkg.go.dev/net) ·
  [net.Conn](https://pkg.go.dev/net#Conn) ·
  [bufio.Reader](https://pkg.go.dev/bufio#Reader)
- [RFC 9293 — TCP](https://www.rfc-editor.org/rfc/rfc9293) — §3.1 on the stream
  abstraction
- [Redis protocol (RESP)](https://redis.io/docs/latest/develop/reference/protocol-spec/)
  — a delimiter-framed protocol worth reading end to end; it is short

**Next: [binary_encoding](../binary_encoding/) →** — what to put *inside* those
frames when the payload is not text.

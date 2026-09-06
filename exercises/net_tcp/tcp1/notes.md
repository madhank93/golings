## tcp1 — accept one connection and answer it

```go
conn, err := ln.Accept()
if err != nil {
    return err
}
defer conn.Close()

line, err := bufio.NewReader(conn).ReadString('\n')
if err != nil {
    return err
}
if line == "PING\n" {
    _, err = conn.Write([]byte("PONG\n"))
}
return err
```

**Why it works**

- A `net.Listener` is a queue of pending connections. `Accept` blocks until one
  is there and hands back a `net.Conn`, which is an `io.ReadWriteCloser` with an
  address — so everything you know about readers and writers applies unchanged.

**Under the hood**

- The listener was created with `net.Listen("tcp", "127.0.0.1:0")`. Port **0**
  means "any free port", and `ln.Addr()` reports which one the kernel picked.
  That is how a test gets a server without hardcoding a port that might be busy.

**Common mistake**

- Not closing the connection. The client's next read then blocks until its
  deadline instead of seeing EOF, because on TCP the *only* way to say "I am
  finished" is to close your side. `defer conn.Close()` is not tidiness here,
  it is part of the protocol.

**Key detail:** `Write` returns how many bytes it wrote. For a short reply it is
always all of them; for a large one it may not be, which is why `io.Copy` and
`Fprintf` exist and hand-rolled write loops usually do not.

**See also:** tcp2 (many messages on one connection) · files1 · the
[chapter](../README.md)

**References**

- pkg.go.dev — net.Listener: https://pkg.go.dev/net#Listener
- pkg.go.dev — net.Conn: https://pkg.go.dev/net#Conn

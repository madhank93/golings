# HTTP client

`net/http` is a production HTTP client in the standard library — connection
pooling, keep-alives, HTTP/2, redirects, and proxy support included. One
exercise covers the shape; this chapter covers what a service needs beyond it.

## 1. The basic request

```go
resp, err := http.Get(url)
if err != nil {
    return "", err
}
defer resp.Body.Close()          // ALWAYS

body, err := io.ReadAll(resp.Body)
```

`http1`. Three things are load-bearing here:

- **`err != nil` covers transport failures only** — DNS, connection refused,
  timeout. A 404 or a 500 is a *successful* request; you have to check
  `resp.StatusCode` yourself. This is the most common HTTP bug in Go.
- **Always close the body**, even when you do not read it. An unclosed body
  holds its connection out of the pool forever — a leak that looks like
  "eventually the service stops making requests".
- **Drain before closing** if you want the connection reused:
  `io.Copy(io.Discard, resp.Body)`. A body closed unread cannot be kept alive.

## 2. Do not use `http.DefaultClient` in production

`http.Get` uses `http.DefaultClient`, which has **no timeout**. A server that
accepts your connection and never answers will hang that goroutine forever.

```go
var client = &http.Client{
    Timeout: 10 * time.Second,     // whole request: dial + write + read body
}
```

Build one client and **share it** — `http.Client` is safe for concurrent use and
its transport is where the connection pool lives. Creating a client per request
throws the pool away and, in a loop, exhausts file descriptors.

When tuning matters, the knobs are on the transport:

```go
&http.Transport{
    MaxIdleConnsPerHost: 100,      // default 2 — often the real bottleneck
    IdleConnTimeout:     90 * time.Second,
}
```

## 3. Context beats `Timeout` for real work

`Client.Timeout` is a blunt per-request cap. A request built with a context
participates in the caller's cancellation, so an abandoned HTTP handler tears
down the upstream call it was waiting on:

```go
req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
if err != nil {
    return err
}
resp, err := client.Do(req)
```

That is the form to default to. `errors.Is(err, context.DeadlineExceeded)`
distinguishes "we ran out of time" from "the server refused".

## 4. Sending and receiving JSON

```go
buf, _ := json.Marshal(payload)
req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(buf))
req.Header.Set("Content-Type", "application/json")

resp, err := client.Do(req)
…
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("upstream returned %s", resp.Status)
}
var out Response
if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
    return err
}
```

`json.NewDecoder(resp.Body).Decode` streams; `io.ReadAll` then `json.Unmarshal`
buffers the whole response first. Prefer the decoder, and cap untrusted
responses with `io.LimitReader` or `http.MaxBytesReader`.

## 5. Testing without the network

```go
srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "pong")
}))
defer srv.Close()

body, err := fetch(srv.URL)
```

This is what `http1`'s test does, and it is the right pattern: a **real** server
on a real loopback port, so the request you send is the request that gets
parsed. Faking `http.Client` with an interface tests your mock instead of your
HTTP.

The corollary for testable code: take the URL (and ideally the `*http.Client`) as
a parameter rather than baking in a constant — the `dependency_injection`
chapter, applied.

## Gotchas

- **`http.DefaultClient` has no timeout.** Set one.
- **`err == nil` does not mean success.** Check `resp.StatusCode`.
- **Always `defer resp.Body.Close()`**, and drain it if you want the connection
  back.
- **Reuse one client**; do not construct one per call.
- **`MaxIdleConnsPerHost` defaults to 2**, which throttles high-concurrency
  clients into constant reconnects.
- **`io.ReadAll` on an untrusted response** is unbounded memory.
- **A request body is an `io.Reader` and can only be read once** — retries need
  `bytes.NewReader`, not a consumed stream.
- **Retry only idempotent requests**, and only with backoff and a cap.

## The exercises

- **http1** — GET a URL, close the body, and read the response, against an
  `httptest` server.

## Source references

- [pkg.go.dev: net/http](https://pkg.go.dev/net/http) ·
  [http.Client](https://pkg.go.dev/net/http#Client) ·
  [http.Transport](https://pkg.go.dev/net/http#Transport)
- [pkg.go.dev: net/http/httptest](https://pkg.go.dev/net/http/httptest)
- [Go blog: Contexts](https://go.dev/blog/context) — cancellation across the
  request boundary

**End of the Stdlib & I/O tier.** Next: [http_server](../http_server/) — the
other end of the same connection.

## http1 — GET, close, read

```go
func fetch(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}
```

**Why it works**

- `http.Get` performs the request and returns a response whose `Body` is an
  `io.ReadCloser`. Reading it gives the payload; closing it returns the
  connection to the pool.

**Under the hood**

- `err != nil` covers **transport** failures only — DNS, refused, timeout. A 404
  or 500 is a *successful* request, so `resp.StatusCode` must be checked
  separately. This is the most common HTTP bug in Go.

**Common mistake**

- Skipping `defer resp.Body.Close()`, or closing without draining. An unclosed
  body holds its connection out of the pool forever; a body closed unread cannot
  be reused. `io.Copy(io.Discard, resp.Body)` before closing keeps the connection
  alive.

**Key detail:** `http.Get` uses `http.DefaultClient`, which has **no timeout** — a
server that accepts and never answers hangs the goroutine forever. Production
code builds one shared `&http.Client{Timeout: …}` and uses
`http.NewRequestWithContext` so the call participates in the caller's
cancellation.

**See also:** stdlib2 (`io.ReadAll` vs `io.Copy`) · context2 (deadlines) ·
httpsrv1 (the other end) · testadv3 (`httptest`) · the [chapter](../README.md)

**References**

- pkg.go.dev — http.Client: https://pkg.go.dev/net/http#Client
- pkg.go.dev — net/http/httptest: https://pkg.go.dev/net/http/httptest

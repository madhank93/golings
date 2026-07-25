## http1 — an HTTP client

```go
resp, err := http.Get(url)
if err != nil {
    return "", err
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
```

**Why it works**

- `http.Get` performs a GET and returns a `*http.Response`. The body is an
  `io.ReadCloser`, so `io.ReadAll` drains it into bytes.

**Nuance:** **always `defer resp.Body.Close()`** — leaking bodies exhausts
connections and file descriptors. Do it *after* checking `err != nil` (on error
there's no body to close). Note the test spins up an `httptest.NewServer` — the
idiomatic way to test HTTP code without hitting the network. For real use, set
timeouts via a custom `http.Client` rather than the default.

**References**

- Go by Example — HTTP Client: https://gobyexample.com/http-client
- pkg.go.dev — net/http/httptest: https://pkg.go.dev/net/http/httptest

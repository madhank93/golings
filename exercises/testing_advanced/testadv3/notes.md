## testadv3 — testing a handler with httptest

```go
func pingHandler(w http.ResponseWriter, r *http.Request) {
    if _, err := w.Write([]byte("pong")); err != nil {
        return
    }
}
```

**Why it works**

- `httptest.NewRequest` builds a request with no network involved and
  `NewRecorder` captures the status, headers, and body the handler writes. The
  test calls the handler directly and asserts on the recorder.

**Under the hood**

- `w.Write` implicitly sends `200 OK` on the first byte, which is why the handler
  needs no explicit `WriteHeader`. It also means headers set after that first
  write are ignored — the ordering rule from httpsrv1.

**Common mistake**

- Reaching for `httptest.NewServer` for handler tests. That starts a real server
  on a loopback port — the right tool for testing a **client** (http1), and
  unnecessary overhead when the unit under test is the handler itself.

**Key detail:** `w.Write` returns `(int, error)`, and ignoring it trips
`errcheck` in this repo's lint step. In a handler the error usually means the
client disconnected — there is nothing to send an error response to, so
returning is the correct handling.

**See also:** httpsrv1 (writing handlers) · http1 (client side) · testadv1 ·
logingest7 · the [chapter](../README.md)

**References**

- pkg.go.dev — httptest.NewRecorder: https://pkg.go.dev/net/http/httptest#NewRecorder
- pkg.go.dev — httptest.NewRequest: https://pkg.go.dev/net/http/httptest#NewRequest

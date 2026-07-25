## testadv3 — httptest

```go
req := httptest.NewRequest(http.MethodGet, "/ping", nil)
rec := httptest.NewRecorder()
pingHandler(rec, req)
res := rec.Result()
```

**Why it works**

- `httptest.NewRequest` builds a request and `httptest.NewRecorder` a fake
  `ResponseWriter`. Calling the handler with them exercises it **without a real
  network**; `rec.Result()` returns the response to assert status and body on.

**Nuance:** two flavors — `NewRecorder` tests a handler **in-process** (fastest,
used here); `httptest.NewServer` spins up a real localhost server for testing full
**clients** or middleware chains end-to-end. Remember to `defer res.Body.Close()`.

**References**

- pkg.go.dev — net/http/httptest: https://pkg.go.dev/net/http/httptest

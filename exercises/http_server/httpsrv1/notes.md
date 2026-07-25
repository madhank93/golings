## httpsrv1 — a handler + httptest

```go
func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" { name = "world" }
    fmt.Fprintf(w, "Hello, %s!", name)
}
rec := httptest.NewRecorder()
greetHandler(rec, httptest.NewRequest("GET", "/greet?name=Go", nil))
```

**Why it works**

- A handler writes its response into the `http.ResponseWriter`. `r.URL.Query().Get`
  reads a query param. `httptest.NewRecorder` captures what the handler wrote so
  you can assert on it — **no real network** needed.

**Nuance:** `ResponseWriter` is a stream you write to (via `Fprintf`, `Write`), not
a value you return. `httptest.NewRecorder()` + `NewRequest()` is the standard way
to unit-test handlers directly; call the handler as a plain function.

**References**

- pkg.go.dev — net/http: https://pkg.go.dev/net/http
- pkg.go.dev — net/http/httptest: https://pkg.go.dev/net/http/httptest

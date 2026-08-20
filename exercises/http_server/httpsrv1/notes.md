## httpsrv1 — a handler and httptest

```go
func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "world"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}
```

**Why it works**

- A handler receives the response writer and the request. `r.URL.Query()` parses
  the query string, `Get` returns `""` for a missing key — so the default is a
  plain empty check — and `Fprintf` writes to `w`, which is an `io.Writer`.

**Under the hood**

- `http.HandlerFunc` is a **named func type with a `ServeHTTP` method**, so any
  function of this shape converts into an `http.Handler`. That is the
  `methods` chapter's "methods on non-struct types" doing real work.

**Common mistake**

- Setting headers after writing. The first `Write` sends the status line and
  headers, so `w.Header().Set(...)` afterwards is silently ignored. Order is:
  headers, then `WriteHeader`, then body.

**Key detail:** the test calls the handler **directly** with
`httptest.NewRequest` and `NewRecorder` — no network, no port, no server. That
is the fastest and most precise way to test a handler; `httptest.NewServer` is
for testing *clients*.

**See also:** httpsrv2 (routing) · httpsrv3 (JSON) · testadv3 · http1 (the
client side) · the [chapter](../README.md)

**References**

- pkg.go.dev — http.HandlerFunc: https://pkg.go.dev/net/http#HandlerFunc
- pkg.go.dev — httptest.NewRecorder: https://pkg.go.dev/net/http/httptest#NewRecorder

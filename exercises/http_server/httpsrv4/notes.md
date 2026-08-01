## httpsrv4 — middleware

```go
func withHeader(value string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("X-App", value)
            next.ServeHTTP(w, r) // call the wrapped handler
        })
    }
}
```

**Why it works**

- Middleware is a `func(http.Handler) http.Handler`: it takes the next handler and
  returns a new one that does something (set a header) then calls
  `next.ServeHTTP`. Wrapping `withHeader("golings")(final)` produces a handler that
  runs the header logic before `final`.

**Key detail:** this is the **decorator** pattern — chain middleware
(`a(b(c(handler)))`) for logging, auth, CORS, recovery, each a small wrapper. Do
work **before** `next.ServeHTTP` (request side) and/or **after** (response side).
Everything composes because they all share the `http.Handler` interface.

**References**

- pkg.go.dev — net/http.Handler: https://pkg.go.dev/net/http#Handler

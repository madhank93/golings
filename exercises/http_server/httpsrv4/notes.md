## httpsrv4 — middleware

```go
func withHeader(value string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("X-App", value)
            next.ServeHTTP(w, r) // the chain continues
        })
    }
}
```

**Why it works**

- Middleware is `func(http.Handler) http.Handler`: a closure that wraps a handler
  and returns a handler. This one sets a header and then calls the wrapped
  handler, so both run for every request.

**Under the hood**

- Because the type is standard, middleware from unrelated sources composes:
  `logging(auth(recovery(mux)))`. It reads outside-in — the outermost sees the
  request first and the response last, which is why logging goes outermost and
  recovery goes close to the handler.

**Common mistake**

- Forgetting `next.ServeHTTP(w, r)`. The chain stops silently: the handler never
  runs and the client gets an empty 200. Nothing errors, which makes it a
  five-minute debugging session the first time.

**Key detail:** to log or measure the **status code**, wrap the
`http.ResponseWriter` in a type that records what `WriteHeader` received — the
real writer will not tell you. That recorder is what `logingest6` builds, and the
reason its middleware can report `status=404`.

**See also:** httpsrv2 · anonymous_functions3 (closures) · slog2 ·
logingest6 · the [chapter](../README.md)

**References**

- pkg.go.dev — http.Handler: https://pkg.go.dev/net/http#Handler
- Go blog — Writing Web Applications: https://go.dev/doc/articles/wiki/

## httpsrv2 — method + path routing

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /users/{id}", userHandler)
// in the handler:
r.PathValue("id") // the {id} segment
```

**Why it works**

- Go 1.22's `ServeMux` understands patterns with a **method** and **wildcards**:
  `"GET /users/{id}"` only matches GET requests, and `r.PathValue("id")` pulls out
  the `{id}` segment. A POST to the same path automatically gets `405 Method Not
  Allowed`.

**Nuance:** before 1.22 you needed a third-party router (chi, gorilla) for this —
now the stdlib mux does method matching, path wildcards, and precedence. The
wildcard name in the pattern must exactly match the `PathValue` argument.

**References**

- The Go Blog — Routing enhancements: https://go.dev/blog/routing-enhancements
- pkg.go.dev — net/http.ServeMux: https://pkg.go.dev/net/http#ServeMux

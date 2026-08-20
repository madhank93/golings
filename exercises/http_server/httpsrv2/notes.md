## httpsrv2 — ServeMux method and path patterns

```go
mux.HandleFunc("GET /users/{id}", userHandler)

func userHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "user %s", r.PathValue("id"))
}
```

**Why it works**

- Since Go 1.22 a `ServeMux` pattern can carry a method and wildcards. `GET
  /users/{id}` matches only GETs on that shape, and `r.PathValue("id")` returns
  the captured segment.

**Under the hood**

- Matching is by **specificity**, not registration order: `/users/{id}` beats
  `/users/`, and a pattern with a method beats one without. A trailing slash
  still means "subtree". Two patterns that overlap without one being more
  specific panic at registration — a conflict you hear about at startup.

**Common mistake**

- Registering on `http.DefaultServeMux` (via `http.HandleFunc`). It is a package
  global that any imported package can add to — which is precisely how
  `net/http/pprof` publishes `/debug/pprof/` (pprof3). Build your own mux.

**Key detail:** before 1.22 this needed a third-party router, so most Go tutorials
still show one. `{name...}` captures the rest of the path, and `{$}` anchors the
end — `"GET /{$}"` matches only `/`.

**See also:** httpsrv1 · httpadv2 (exact wildcard names) · pprof3 ·
logingest3 · the [chapter](../README.md)

**References**

- pkg.go.dev — http.ServeMux: https://pkg.go.dev/net/http#ServeMux
- Go blog — Routing enhancements for Go 1.22: https://go.dev/blog/routing-enhancements

## httpadv2 — PathValue and wildcard names

```go
mux.HandleFunc("GET /users/{id}/posts/{postID}", func(w, r) {
    postID := r.PathValue("postID") // name must match the {wildcard} exactly
})
```

**Why it works**

- `r.PathValue(name)` returns the path segment captured by `{name}` in the pattern.
  The name must match a wildcard **exactly** — `PathValue("post_id")` returns `""`
  because the pattern declares `{postID}`, not `{post_id}`.

**Key detail:** a mismatched or non-existent wildcard name silently returns the empty
string, not an error — a quiet bug. Patterns can carry several wildcards
(`/users/{id}/posts/{postID}`), each addressable by its exact name.

**References**

- pkg.go.dev — net/http.ServeMux: https://pkg.go.dev/net/http#ServeMux

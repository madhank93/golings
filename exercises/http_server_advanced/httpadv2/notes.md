## httpadv2 — PathValue takes the exact wildcard name

```go
mux.HandleFunc("GET /users/{id}/posts/{postID}", func(w http.ResponseWriter, r *http.Request) {
    postID := r.PathValue("postID") // matches the pattern exactly
    ...
})
```

**Why it works**

- `r.PathValue` looks up a wildcard **by the name written in the pattern**. The
  pattern says `{postID}`, so `PathValue("post_id")` finds nothing.

**Under the hood**

- A miss returns `""` — no error, no panic. The empty string then flows into the
  handler as if the client had sent it, which is why a typo here produces a
  puzzling 200 with an empty body rather than a failure.

**Common mistake**

- Assuming snake_case works because the URL looks like that. Wildcard names are
  case-sensitive identifiers in the pattern and nothing else; the path segment's
  content is unrelated to the name.

**Key detail:** `{name}` matches exactly one segment, `{name...}` (trailing) the
rest of the path including slashes, and `{$}` anchors the end. Registering the
**same wildcard name twice** in one pattern panics — the one mistake in this area
that fails loudly. Treat an empty `PathValue` as a 400 rather than passing it on.

**See also:** httpsrv2 (routing basics) · httpsrv3 (validating input) ·
logingest3 · the [chapter](../README.md)

**References**

- pkg.go.dev — Request.PathValue: https://pkg.go.dev/net/http#Request.PathValue
- pkg.go.dev — ServeMux patterns: https://pkg.go.dev/net/http#ServeMux

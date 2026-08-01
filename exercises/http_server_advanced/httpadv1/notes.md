## httpadv1 — CSRF defense with CrossOriginProtection

```go
func protect(h http.Handler) http.Handler {
    return http.NewCrossOriginProtection().Handler(h)
}
```

**Why it works**

- `http.CrossOriginProtection` (Go 1.25) blocks cross-site, **state-changing**
  requests using the browser's `Sec-Fetch-Site` metadata. Wrapping a handler with
  it rejects a cross-site POST — CSRF defense with **no tokens or cookies**.

**Key detail:** it leans on modern browsers sending `Sec-Fetch-*` headers; safe
methods (GET/HEAD) and same-origin requests pass through. This is a stdlib answer
to a problem that historically needed a CSRF-token library. Opt in by wrapping the
handlers that mutate state.

**References**

- pkg.go.dev — http.CrossOriginProtection: https://pkg.go.dev/net/http#CrossOriginProtection

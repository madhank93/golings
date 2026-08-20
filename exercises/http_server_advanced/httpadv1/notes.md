## httpadv1 — CrossOriginProtection

```go
func protect(h http.Handler) http.Handler {
    return http.NewCrossOriginProtection().Handler(h)
}
```

**Why it works**

- `http.CrossOriginProtection` (Go 1.25) wraps a handler and rejects **unsafe
  methods** — POST, PUT, DELETE, PATCH — whose request originated cross-site.
  The bare handler accepts them, which is what the test catches.

**Under the hood**

- The check reads the browser's **`Sec-Fetch-Site`** header, which describes
  where the request came from. That means CSRF defence with no tokens, no session
  storage, and no per-form plumbing. Safe methods (GET, HEAD, OPTIONS) pass
  through untouched — a GET that changes state is the actual bug.

**Common mistake**

- Treating it as authentication. It relies on a real browser sending truthful
  metadata; any client that sets its own headers is unaffected. It defends
  against a malicious *page*, which is exactly what CSRF is, and nothing else.

**Key detail:** the exceptions are explicit — `AddTrustedOrigin` for a known
partner site, `AddInsecureBypassPattern` for endpoints called by non-browser
clients (webhooks). Wrap the whole mux, once, rather than per handler.

**See also:** httpsrv4 (middleware shape) · httpadv3 (proxy header handling) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — http.CrossOriginProtection: https://pkg.go.dev/net/http#CrossOriginProtection
- MDN — Sec-Fetch-Site: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site

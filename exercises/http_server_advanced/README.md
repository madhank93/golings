# HTTP servers — advanced

Three things a real service needs that the basics chapter left out: protection
against cross-site requests, exact handling of path wildcards, and proxying to a
backend. All three are in `net/http` and its `httputil` subpackage — still no
framework.

## 1. CSRF protection without tokens (`http.CrossOriginProtection`, Go 1.25)

```go
func protect(h http.Handler) http.Handler {
    return http.NewCrossOriginProtection().Handler(h)
}
```

`httpadv1`. Classic CSRF defence means generating a token, storing it in the
session, embedding it in every form, and comparing on submit. Go 1.25 offers a
simpler mechanism: browsers now send **`Sec-Fetch-Site`** describing where a
request originated, and `CrossOriginProtection` rejects **unsafe methods**
(POST, PUT, DELETE, PATCH) whose origin is cross-site.

Safe methods — GET, HEAD, OPTIONS — pass through untouched, which is correct:
they are not supposed to change state, and a GET that does is the actual bug.

Configure the exceptions explicitly:

```go
p := http.NewCrossOriginProtection()
p.AddTrustedOrigin("https://admin.example.com")   // allow a known partner
p.AddInsecureBypassPattern("/api/webhook")        // non-browser callers
```

The limitation to understand: this relies on the **browser** sending truthful
`Sec-Fetch-Site` metadata. It defends against a malicious page in a real
browser — which is what CSRF is — and does nothing for a client that sets its
own headers, which is what authentication is for.

## 2. Wildcards must match by name

```go
mux.HandleFunc("GET /users/{id}/posts/{postID}", func(w http.ResponseWriter, r *http.Request) {
    postID := r.PathValue("postID")     // exact wildcard name
})
```

`httpadv2`. `r.PathValue` takes the wildcard's name **as written in the
pattern**. Ask for `"post_id"` when the pattern says `{postID}` and you get
`""` — no error, no panic, just an empty string flowing into your handler.

The rules worth knowing:

- Names are case-sensitive and must match exactly.
- `{name}` matches exactly one path segment; `{name...}` (trailing) matches the
  rest of the path including slashes.
- `{$}` at the end anchors: `"GET /{$}"` matches only `/`, not everything.
- Duplicate wildcard names in one pattern are a **panic at registration**, which
  is the one mistake here that fails loudly.

Since a typo is silent, validate: treat an empty `PathValue` as a 400 rather
than passing it downstream.

## 3. Reverse proxying with `Rewrite`

```go
proxy := &httputil.ReverseProxy{
    Rewrite: func(pr *httputil.ProxyRequest) {
        pr.SetURL(target)          // point the outbound request at the backend
        pr.Out.Host = pr.In.Host   // optionally preserve the inbound Host
    },
}
```

`httpadv3`. `ReverseProxy` copies a request to a backend and streams the
response back. The `Rewrite` hook (Go 1.20) replaced the older `Director` field
for a security reason worth knowing:

`Director` mutated the request in place, so **inbound hop-by-hop and forwarding
headers could leak through** to the backend — a client could forge
`X-Forwarded-For`. `Rewrite` receives a `ProxyRequest` with separate `In` and
`Out` requests, and `SetURL` strips the inbound forwarding headers and sets
`X-Forwarded-For` itself. Same code shape, safer default.

`SetURL` also **joins paths**: with a target of `http://backend/api`, a request
for `/users` goes to `/api/users`. And `pr.Out.Host` is deliberately *not*
preserved by default — set it explicitly when the backend does virtual hosting.

## 4. Graceful shutdown, since a real server needs it

```go
srv := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second}

go func() {
    if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}()

<-ctx.Done()                        // SIGINT/SIGTERM via signal.NotifyContext

shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
if err := srv.Shutdown(shutdownCtx); err != nil {
    log.Printf("forced shutdown: %v", err)
}
```

`Shutdown` stops accepting new connections and waits for in-flight handlers to
finish, up to the deadline. Two details: `ListenAndServe` returns
`http.ErrServerClosed` on a clean shutdown, so that is not an error to report;
and `Shutdown` does **not** wait for hijacked or WebSocket connections — track
those yourself. The capstone (`logingest6`) builds exactly this.

## Gotchas

- **`PathValue` with a wrong name returns `""`** silently.
- **Duplicate wildcard names panic** at registration time.
- **`CrossOriginProtection` trusts browser metadata** — not a substitute for
  authentication.
- **`Director` is deprecated**; use `Rewrite` to avoid header leakage.
- **`SetURL` joins paths** — a target with a path prefixes every route.
- **`http.ErrServerClosed` after `Shutdown` is success**, not failure.
- **Proxies need timeouts on both sides**, or a slow backend pins your
  goroutines.

## The exercises

- **httpadv1** — wrap a handler in `CrossOriginProtection` so cross-site POSTs
  are blocked.
- **httpadv2** — read the path value with the exact wildcard name.
- **httpadv3** — point the proxy at its backend with `pr.SetURL`.

## Source references

- [pkg.go.dev: http.CrossOriginProtection](https://pkg.go.dev/net/http#CrossOriginProtection)
  (Go 1.25)
- [pkg.go.dev: ServeMux patterns](https://pkg.go.dev/net/http#ServeMux) ·
  [Request.PathValue](https://pkg.go.dev/net/http#Request.PathValue)
- [pkg.go.dev: httputil.ReverseProxy](https://pkg.go.dev/net/http/httputil#ReverseProxy) ·
  [ProxyRequest.SetURL](https://pkg.go.dev/net/http/httputil#ProxyRequest.SetURL)
- [pkg.go.dev: Server.Shutdown](https://pkg.go.dev/net/http#Server.Shutdown)
- [MDN: Sec-Fetch-Site](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site)

**Next: [cli](../cli/) →** — the other kind of program, and the flag parsing that
makes it testable.

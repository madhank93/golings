## httpadv3 — ReverseProxy.Rewrite

```go
&httputil.ReverseProxy{
    Rewrite: func(pr *httputil.ProxyRequest) {
        pr.SetURL(target) // point the outbound request at the backend
    },
}
```

**Why it works**

- `ReverseProxy` forwards requests to a backend. The modern `Rewrite` hook
  (replacing the deprecated `Director`) receives a `*ProxyRequest`; `pr.SetURL(target)`
  retargets the outbound request at the backend while preserving the inbound path.

**Key detail:** `Rewrite` supersedes `Director` because it exposes **both** the inbound
(`pr.In`) and outbound (`pr.Out`) requests, and `SetURL` correctly sets the
`X-Forwarded-*` headers — things easy to get wrong by hand. A reverse proxy in ~10
lines of stdlib.

**References**

- pkg.go.dev — httputil.ReverseProxy: https://pkg.go.dev/net/http/httputil#ReverseProxy

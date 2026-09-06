## httpadv3 — ReverseProxy.Rewrite

```go
&httputil.ReverseProxy{
    Rewrite: func(pr *httputil.ProxyRequest) {
        pr.SetURL(target) // point the outbound request at the backend
    },
}
```

**Why it works**

- `Rewrite` runs for every proxied request and is where the outbound request gets
  its destination. Without `SetURL` the outbound URL is never retargeted, so the
  proxy forwards nowhere.

**Under the hood**

- `Rewrite` (Go 1.20) replaced the deprecated `Director` for a security reason:
  `Director` mutated the request in place, so inbound hop-by-hop and
  `X-Forwarded-*` headers could leak through to the backend — a client could
  forge `X-Forwarded-For`. `ProxyRequest` separates `In` from `Out`, and `SetURL`
  strips those inbound headers and sets forwarding headers itself.

**Common mistake**

- Expecting the inbound `Host` to be preserved. It is not, deliberately — set
  `pr.Out.Host = pr.In.Host` when the backend does virtual hosting.

**Key detail:** `SetURL` **joins paths**: with a target of `http://backend/api`, a
request for `/users` is forwarded to `/api/users`. And a proxy needs timeouts on
both sides, or one slow backend pins your goroutines.

**See also:** httpadv1 (header trust) · http1 (client timeouts) · httpsrv4 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — httputil.ReverseProxy: https://pkg.go.dev/net/http/httputil#ReverseProxy ·
  ProxyRequest.SetURL: https://pkg.go.dev/net/http/httputil#ProxyRequest.SetURL

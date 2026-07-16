# HTTP Servers — Advanced

Production HTTP servers need more than routing. These exercises cover built-in
CSRF protection, precise path-wildcard extraction, and reverse proxying with the
modern `Rewrite` hook — patterns that show up in real services built on
`net/http`.

- **httpadv1** — `http.CrossOriginProtection` (Go 1.25) blocks cross-site POSTs
- **httpadv2** — `r.PathValue` and matching `{wildcard}` names exactly
- **httpadv3** — `httputil.ReverseProxy.Rewrite` with `ProxyRequest.SetURL`

## Resources

- [http.CrossOriginProtection](https://pkg.go.dev/net/http#CrossOriginProtection)
- [ServeMux patterns (Go 1.22)](https://pkg.go.dev/net/http#ServeMux)
- [httputil.ReverseProxy](https://pkg.go.dev/net/http/httputil#ReverseProxy)

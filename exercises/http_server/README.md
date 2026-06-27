# HTTP Server

Build HTTP services with the standard library — no framework. `net/http` gives
you handlers and routing; since Go 1.22 `http.ServeMux` supports method + path
patterns like `"GET /users/{id}"`. Test handlers without a network using
`net/http/httptest`.

- **httpsrv1** — write an `http.HandlerFunc` and test it with `httptest`
- **httpsrv2** — route with `http.ServeMux` and Go 1.22 `"GET /users/{id}"` patterns
- **httpsrv3** — decode a JSON request and encode a JSON response
- **httpsrv4** — wrap a handler with middleware (`func(http.Handler) http.Handler`)

Pairs with `testing_advanced/testadv3` (handler testing) and `http_client` (the
client side).

## Resources

- [net/http](https://pkg.go.dev/net/http)
- [Routing enhancements (Go 1.22)](https://go.dev/blog/routing-enhancements)
- [Learn Go with Tests — HTTP server](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server)

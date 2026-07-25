## context3 — request-scoped values

```go
type ctxKey string
const userKey ctxKey = "user"

ctx := context.WithValue(context.Background(), userKey, "alice")
u, ok := ctx.Value(userKey).(string) // "alice", true
```

**Why it works**

- `context.WithValue` attaches a key/value that any function down the call chain
  can read with `ctx.Value(key)`. A missing key returns `nil`, so the comma-ok
  assertion falls back to `"anonymous"`.

**Nuance:** use a **private, custom key type** (`type ctxKey string`) — not a bare
`string` — so your key can't collide with another package's. Reserve context
values for **request-scoped** data (request IDs, auth user), never for passing
optional function parameters; those belong in the signature.

**References**

- pkg.go.dev — context.WithValue: https://pkg.go.dev/context#WithValue
- The Go Blog — Go Concurrency Patterns: Context: https://go.dev/blog/context

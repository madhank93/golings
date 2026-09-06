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

**Under the hood**

- `WithValue` does not mutate anything; it prepends one immutable node to a
  linked chain. `Value` walks that chain toward the root comparing keys with
  `==`, so lookup is linear in the number of values attached — fine for a handful
  of request-scoped items, wrong as a general store.

**Common mistake**

- Using a bare `string` as the key. Key comparison includes the **type**, so an
  unexported `type ctxKey string` cannot collide with another package that also
  picked `"user"` — while two packages using plain strings silently overwrite
  each other. Unexported key types make collision impossible.

**Key detail:** put request-scoped data that crosses API boundaries in a context —
trace ids, the authenticated user, a request logger — supplied by middleware.
Never put a function's actual inputs there: a function whose behaviour depends on
values invisible in its signature cannot be called correctly. Always assert with
comma-ok; the value may legitimately be absent.

**See also:** context1 (cancellation) · httpsrv4 (middleware attaching values) ·
logingest7 (request-scoped `slog` logger) · the [context chapter](../README.md)

**References**

- pkg.go.dev — context.WithValue: https://pkg.go.dev/context#WithValue
- Go blog — Contexts and structs: https://go.dev/blog/context-and-structs
- Go blog — Context: https://go.dev/blog/context

## modules8 — blank imports and init

```go
import (
    "fmt"

    "example.com/notify/sender"
    _ "example.com/notify/sender/email"
)
```

**Why it works**

- `email` registers itself with `sender` from its `init` function, and `init`
  only runs for packages that are part of the program. The blank import `_`
  links the package in for that side effect alone — no name is bound, so the
  "imported and not used" rule does not fire.

**Under the hood**

- Initialization goes dependencies first: `sender`'s package-level `registry`
  map is set up, then `email.init` registers into it, then `notify` — all before
  the test calls `Send`. Within a package, variables initialize before any
  `init`, and multiple `init` functions run in source order.

**Common mistake**

- Expecting a package to self-register just because it sits in the module.
  Being in the module tree is not enough; if nothing imports it, the linker
  never sees it and the lookup quietly fails (`no sender registered for "email"`).

**Key detail:** this is the `database/sql` driver pattern —
`import _ "github.com/lib/pq"` makes `sql.Open("postgres", …)` work, because
the driver calls `sql.Register` in its `init`. Put the blank import in `main`
(or the one package that chooses the driver), not in a reusable library, so
callers decide what gets linked.

**See also:** modules2 (import paths) · modules3 (exported names — email exports none) · the
[chapter](../README.md)

**References**

- Go spec — Package initialization: https://go.dev/ref/spec#Package_initialization
- Effective Go — Imports for side effect: https://go.dev/doc/effective_go#blank_import
- pkg.go.dev — database/sql `Register`: https://pkg.go.dev/database/sql#Register

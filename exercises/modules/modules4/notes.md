## modules4 — internal packages

```go
import "example.com/bank/ledger"

func Statement(cents int) string {
    return "Balance: " + ledger.Format(cents)
}
```

**Why it works**

- `ledger/internal/money` can be imported only by code in the tree rooted at
  `ledger/`. `bank` is outside that tree, so it goes through `ledger.Format` —
  the API `ledger` chose to offer, which also prefixes the `-` for overdrafts
  that `money.Dollars` never handled.

**Under the hood**

- The rule is enforced by the go command, not the compiler: an import of a
  path containing an `internal` element is refused unless the importer lives
  under the parent of that `internal` directory. The error is
  `use of internal package ... not allowed`. It holds across modules too — no
  other module can import your `internal` packages.

**Common mistake**

- Moving the package out of `internal/` to make the error go away. That turns
  an implementation detail into public API that you then have to keep stable.
  The check is the design working: call the exported wrapper, or export what
  is missing from the parent package.

**Key detail:** the root is internal's **parent**, not the module root. A
top-level `internal/` is shared by the whole module, while `ledger/internal/`
is private to `ledger` and its subpackages — so placement sets how widely a
helper is shared.

**See also:** modules2 (import paths) · modules3 (exported names, the
per-identifier version of the same idea) · the [chapter](../README.md)

**References**

- go command — Internal packages: https://pkg.go.dev/cmd/go#hdr-Internal_packages
- Go 1.4 release notes — internal packages: https://go.dev/doc/go1.4#internalpackages
- Organizing a Go module: https://go.dev/doc/modules/layout

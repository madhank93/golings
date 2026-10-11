## modules2 — import path = module path + directory

```go
// module example.com/shop, package in ./cart
import "example.com/shop/cart"
```

**Why it works**

- A package's import path is the module path from `go.mod` joined with the
  package's directory relative to the module root. `cart.go` sits in `./cart`
  of module `example.com/shop`, so it is `example.com/shop/cart`. The package
  is still referred to as `cart` in code — that comes from its `package` clause.

**Under the hood**

- An import path whose first element has no dot, like `cart` or `fmt`, is
  treated as a standard-library path. The go command looks for it under
  `$GOROOT/src`, finds nothing, and reports
  `package cart is not in std` — it never searches your module for it.

**Common mistake**

- Reaching for a relative import, `import "./cart"`. Relative paths are not
  allowed in module mode; every import is the full path, even between packages
  of the same module.

**Key detail:** directories are the unit, not files. Every `.go` file in one
directory belongs to one package (plus its `_test` package), and a
subdirectory is always a separate package with its own import path — nesting
on disk implies nothing about access.

**See also:** modules1 (the module path) · the [chapter](../README.md)

**References**

- How to Write Go Code — importing packages from your module: https://go.dev/doc/code#ImportingLocal
- Go Modules Reference — module paths: https://go.dev/ref/mod#module-path

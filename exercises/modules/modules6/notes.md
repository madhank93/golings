## modules6 — major version suffix /v2

```go
import "example.com/units/v2"

f, unit := units.CToF(celsius) // package name is still units
```

**Why it works**

- From v2 on, the major version is part of the module path:
  `example.com/units/v2`, as its `go.mod` declares and the main module
  requires. Import paths start with the module path, so the import must carry
  `/v2` as well. The package clause still says `package units`, so code calls
  `units.CToF` unchanged.

**Under the hood**

- This is semantic import versioning: different major versions have different
  paths, so the go command treats `example.com/units` and
  `example.com/units/v2` as unrelated modules. Plain `example.com/units` is not
  required anywhere, hence `cannot find module providing package
  example.com/units`. It also means v1 and v2 can be linked into one program
  side by side.

**Common mistake**

- Tagging a `v2.0.0` release without changing the `module` line. The go
  command rejects a v2+ version whose `go.mod` path lacks the `/v2` suffix, so
  publishing a major version means updating `go.mod` and every internal import
  in the module.

**Key detail:** v0 and v1 never take a suffix. The rule exists because a new
major version is allowed to break compatibility (here `CToF` grew a second
result), and a distinct path keeps existing importers on the old API until
they opt in.

**See also:** modules5 (require + replace, used here to point v2 at
`./units`) · modules2 (import paths) · the [chapter](../README.md)

**References**

- Go Modules Reference — major version suffixes: https://go.dev/ref/mod#major-version-suffixes
- Go blog — Go Modules: v2 and Beyond: https://go.dev/blog/v2-go-modules
- Developing a major version update: https://go.dev/doc/modules/major-version

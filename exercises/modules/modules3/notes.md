## modules3 — exported identifiers

```go
type Reading struct {
    Celsius float64
}

func Fahrenheit(r Reading) float64 {
    return r.Celsius*9/5 + 32
}
```

**Why it works**

- An identifier is exported — visible to importers — when its first character
  is an upper-case letter. That applies to everything at package level and to
  struct fields and methods too, so both the function and the `Celsius` field
  need capitalising for `temp.Reading{Celsius: 100}` to compile outside the
  package.

**Under the hood**

- A test file declaring `package temp_test` is compiled as a separate package
  that imports `temp` like any other client. It sees only the exported API —
  which is the point: it tests what users get. Files declaring `package temp`
  in the same directory would see everything.

**Common mistake**

- Renaming the declaration and missing a use. Visibility is the name itself,
  so `celsius` → `Celsius` is a rename across every reference in the package;
  the compiler lists the stragglers as undefined.

**Key detail:** "upper case" means Unicode class Lu. A name starting with `_`
or with a letter that has no case (say, a CJK character) can never be exported.
Exporting a type does not export its fields: a struct can be public while
keeping every field private, and callers then construct it only through
functions you provide.

**See also:** modules2 · functions1 (case as access control) · structs1 · the
[chapter](../README.md)

**References**

- Go spec — Exported identifiers: https://go.dev/ref/spec#Exported_identifiers
- go command — test packages: https://pkg.go.dev/cmd/go#hdr-Test_packages
- Effective Go — Names: https://go.dev/doc/effective_go#names

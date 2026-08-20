## stdlib1 — JSON struct tags

```go
type User struct {
    Name string `json:"full_name"`
    Age  int    `json:"user_age"`
}
```

**Why it works**

- `encoding/json` uses the field name by default; a tag renames the key it maps
  to. The incoming JSON uses `full_name`/`user_age`, so without tags neither
  field is populated.

**Under the hood**

- Marshaling walks the struct with reflection, which is why **only exported
  fields are visible**. A lowercase field is silently skipped — no error, no key.
  Matching on decode is case-*insensitive*, so `Name` would already accept
  `"name"`; tags are needed here because the keys differ by more than case.

**Common mistake**

- Passing a value to `Unmarshal`. It needs a pointer (`&u`) to write into, and
  passing a value fails at run time with
  `json: Unmarshal(non-pointer main.User)`.

**Key detail:** other tag options: `json:"-"` never encodes the field,
`json:"name,omitempty"` drops it when empty, and `json:",string"` encodes a
number as a JSON string. A malformed tag fails silently — `go vet`'s `structtag`
check catches it.

**See also:** stdlib7 (`omitzero`) · structs1 (field basics) · reflect2 (how
tags are read) · the [chapter](../README.md)

**References**

- pkg.go.dev — encoding/json: https://pkg.go.dev/encoding/json
- Go blog — JSON and Go: https://go.dev/blog/json

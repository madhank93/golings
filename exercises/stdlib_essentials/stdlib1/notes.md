## stdlib1 — encoding/json struct tags

```go
type User struct {
    Name string `json:"full_name"`
    Age  int    `json:"user_age"`
}
json.Unmarshal(data, &u)
```

**Why it works**

- A **struct tag** maps a Go field to a JSON key: `json:"full_name"` binds the
  exported `Name` field to the `"full_name"` key, so `Unmarshal` fills it from
  that key.

**Key detail:** only **exported** (capitalized) fields are (un)marshaled — `json`
can't see unexported ones. Unmarshal takes a **pointer** (`&u`) so it can write
into your struct. Without a tag, matching is case-insensitive on the field name;
add `,omitempty` to skip zero values when marshaling.

**References**

- pkg.go.dev — encoding/json: https://pkg.go.dev/encoding/json
- The Go Blog — JSON and Go: https://go.dev/blog/json

## maps3 — a missing key returns the zero value

```go
phone := phoneBook["Ana"]        // was "Anna" — a typo, not an error
phoneBook["Laura"] = "+11 99 98 97"
phone := phoneBook["Laura"]
```

**Why it works**

- Reading a key that is not there never fails: it yields the value type's zero
  value, so the mistyped `"Anna"` quietly produced `""` and only the test noticed.
  Fixing the key — and inserting `"Laura"` before reading her — makes both
  lookups return real numbers.

**Under the hood**

- The lookup hashes the key, finds the bucket, and compares keys with `==`. A
  miss has no failure channel to report through, so the runtime returns the zero
  value of `V` — the same value a stored zero would give. The two-result form
  exists precisely to tell them apart.

**Common mistake**

- Testing for absence with `if m[k] == ""`. That is also true for a key stored
  with an empty value. The correct form is comma-ok:

```go
if phone, ok := phoneBook["Ana"]; ok { … }
```

**Key detail:** the same asymmetry makes `map[string]bool` a clean set — absent
and `false` mean the same thing there, so `if set[k]` reads correctly without
comma-ok.

**See also:** maps1 · variables3 (zero values) · safety2 (`sync.Map`'s
`Load` returns the same pair) · the [chapter](../README.md)

**References**

- Go blog — Go maps in action: https://go.dev/blog/maps
- Go spec — Index expressions: https://go.dev/ref/spec#Index_expressions

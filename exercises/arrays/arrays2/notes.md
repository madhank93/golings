## arrays2 — every element shares one type

```go
names := [4]string{"John", "Maria", "Carl", "Peter"}
```

**Why it works**

- `[4]string` says every element is a `string`. The literal `10` is a number, and
  Go performs no coercion, so the mixed version fails with
  `cannot use 10 (untyped int constant) as string value`.

**Under the hood**

- An array is a contiguous block of `N` elements of identical size — that is what
  makes `a[i]` a single address computation. Mixed types would have no fixed
  stride and no fixed element size, so the restriction is structural rather than
  stylistic.

**Common mistake**

- Reaching for `[]any` to hold mixed values. It compiles, but every read then
  needs a type assertion and the compiler stops helping you. A struct with named
  fields is almost always what the data actually is.

**Key detail:** arrays are **values**. `b := a` copies all four strings, and
passing one to a function copies it too — unlike a slice, which shares its
backing array.

**See also:** arrays1 (indexing) · structs1 (mixed fields, named) ·
slices1 · the [chapter](../README.md)

**References**

- Go spec — Array types: https://go.dev/ref/spec#Array_types
- Go by Example — Arrays: https://gobyexample.com/arrays

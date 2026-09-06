## strings2 — bytes are not characters

```go
func charCount(s string) int {
    return utf8.RuneCountInString(s)
}
```

**Why it works**

- `len(s)` counts **bytes**. In UTF-8 one character can be several bytes — `é`
  is 2, `世` is 3 — so `len("héllo")` is 6 and `len("世界")` is 6 too.
  `utf8.RuneCountInString` decodes and counts code points instead: 5 and 2.

**Under the hood**

- A string value is a pointer plus a length, over immutable bytes. `s[i]` hands
  back one byte with no decoding (`"héllo"[1]` is 195, half of `é`), while
  `for i, r := range s` decodes as it walks — which is why its index jumps
  0, 1, 3, 4, 5 rather than counting up by one.

**Common mistake**

- Converting with `[]rune(s)` just to call `len`. It works, but it decodes and
  copies the entire string; `RuneCountInString` answers the same question
  without allocating. Reach for `[]rune` only when you need random access by
  character.

**Key detail:** "how long is this text?" has three answers — bytes, runes, and
user-visible glyphs. Even runes are not the last word: an emoji with a skin-tone
modifier is several runes that render as one character.

**See also:** strings1 (the package) · primitive_types4 (`byte` vs `rune`) ·
range1 (ranging a collection) · the [chapter](../README.md)

**References**

- Go blog — Strings, bytes, runes and characters: https://go.dev/blog/strings
- pkg.go.dev — unicode/utf8: https://pkg.go.dev/unicode/utf8

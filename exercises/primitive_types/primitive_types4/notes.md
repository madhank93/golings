## primitive_types4 — a byte holds a character's code

```go
var b2 byte = 'a' // 'a' is the rune literal for code point 97
fmt.Println("representation for b2:", b2) // prints 97
```

**Why it works**

- `''` (empty) is an illegal rune literal. A single-quoted character like `'a'`
  is a **rune constant** whose numeric value (97) fits in a `byte`.

**Under the hood**

- `byte` is an alias for `uint8` and `rune` is an alias for `int32` — not
  distinct types, just names that say what the number means. A rune literal is
  an untyped constant, so `'a'` fits in a `byte` while `'世'` (19990) does not
  and fails to compile in the same slot.

**Common mistake**

- Expecting `fmt.Println(b2)` to print `a`. It prints `97`, because the value
  *is* a number. Use `string(b2)` or `%c` to see the character.

**Key detail:** single quotes make one rune, double quotes make a string, and
backticks make a raw string that ignores escapes. Three different things.

**See also:** strings2 (bytes vs runes in a whole string) ·
primitive_types5 (type names) · the [chapter](../README.md)

**References**

- Go spec — Rune literals: https://go.dev/ref/spec#Rune_literals
- Go blog — Strings, bytes, runes and characters: https://go.dev/blog/strings

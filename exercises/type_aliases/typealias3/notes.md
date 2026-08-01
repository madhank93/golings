## typealias3 — alias vs defined type

```go
type MyByte = byte // ALIAS (note the =): a second name for byte
func asByte(b byte) byte { return b }

var m MyByte = 65
asByte(m) // works with NO conversion — MyByte IS byte
```

**Why it works**

- `type MyByte = byte` (with `=`) is a **type alias** — not a new type, just
  another spelling of `byte`. So `MyByte` and `byte` are fully interchangeable and
  `asByte(m)` needs no conversion.

**Key detail:** the single `=` is the whole difference. `type MyByte byte` (no `=`)
would be a **distinct** defined type requiring `byte(m)` — that's typealias1's
world. Aliases exist mainly for **gradual refactoring** (moving a type between
packages without breaking callers); `byte` and `rune` are themselves aliases for
`uint8` and `int32`.

**References**

- Go spec — Alias declarations: https://go.dev/ref/spec#Alias_declarations

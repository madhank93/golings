## typealias3 — an alias is the same type

```go
type MyByte = byte // note the = : an ALIAS, not a new type

var m MyByte = 65
asByte(m) // passes as a byte with no conversion
```

**Why it works**

- With the `=`, `MyByte` is a second **name** for `byte`, not a new type. The two
  are interchangeable everywhere — arguments, assignments, method sets, interface
  satisfaction. Without the `=` it is a defined type and the call needs
  `byte(m)`.

**Under the hood**

- The alias is resolved during type-checking and does not survive into the type
  system: `reflect.TypeOf(m)` reports `uint8`. You already use several —
  `byte` is itself an alias for `uint8`, `rune` for `int32`, and `any` for
  `interface{}`.

**Common mistake**

- Reaching for an alias to add safety. It adds none: an alias cannot be
  distinguished from what it names and cannot carry new methods. When you want
  the compiler to stop a mix-up, define a type (drop the `=`).

**Key detail:** aliases exist mainly to **move a type between packages** without
breaking callers — leave `type OldName = newpkg.NewName` behind during a
migration. Since Go 1.24 they may also take type parameters, which is handy for
naming a long generic type.

**See also:** typealias1 (defined types) · typealias2 (methods) ·
primitive_types4 (`byte`/`rune` are aliases) · the [chapter](../README.md)

**References**

- Go spec — Alias declarations: https://go.dev/ref/spec#Alias_declarations
- Go blog — Alias names: https://go.dev/blog/alias-names

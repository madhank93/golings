# Type aliases & defined types

Two declarations that look almost identical do completely different things:

```go
type Celsius float64    // DEFINED TYPE — a brand-new type
type MyByte  = byte     // ALIAS — a second name for an existing type
```

The `=` is the whole difference. A defined type shares its underlying
representation but is a **distinct type** the compiler will not mix with
anything else. An alias is not a type at all — it is a name, resolved away, and
completely interchangeable with what it names.

The three exercises make each half concrete: convert across a defined type, hang
a method on one, and see an alias pass with no conversion at all.

## 1. Defined types are distinct

```go
type Celsius float64

var boiling Celsius = 100
func toFahrenheit(c float64) float64 { … }

toFahrenheit(boiling)            // does NOT compile
toFahrenheit(float64(boiling))   // explicit conversion — typealias1
```

`Celsius` and `float64` have the same layout and the same arithmetic, and Go
still refuses to substitute one for the other. That refusal is the feature: it
is how a codebase stops a temperature being passed where a length was expected,
or a `UserID` where an `OrderID` belongs — the class of bug that plain `int`
parameters invite.

Conversion between them is free at run time (no code is generated) and explicit
in the source, which is exactly the trade Go wants: zero cost, visible intent.

Note what a defined type does *not* inherit: methods. `type MyTime time.Time`
starts with no methods at all, because the new type's method set begins empty —
only the underlying representation carries over.

## 2. Methods are why you define a type

```go
type Celsius float64

func (c Celsius) String() string {
    return fmt.Sprintf("%g°C", float64(c))
}

fmt.Sprint(Celsius(25))   // "25°C" — typealias2
```

A defined type gives you somewhere to attach behaviour, and satisfying
`fmt.Stringer` makes every `fmt` call format it your way. Note the
`float64(c)` inside: formatting `c` directly with `%v` would call `String()`
again, recursing until the stack overflows. Convert to the underlying type
first — the single most common self-inflicted infinite loop in Go.

This is also why you cannot attach a method to an alias of a foreign type:
`type MyTime = time.Time` names `time.Time` itself, and methods may only be
declared on types defined in the current package.

## 3. Aliases are the same type

```go
type MyByte = byte      // note the =

var m MyByte = 65
func asByte(b byte) byte { return b }

asByte(m)               // compiles: MyByte IS byte — typealias3
```

No conversion, no distinct identity, and `MyByte` and `byte` are
interchangeable everywhere including method sets and interface satisfaction.
`reflect.TypeOf(m)` reports `uint8` — the alias has vanished by then.

You already use several: `byte` is itself an alias for `uint8`, `rune` for
`int32`, and `any` for `interface{}`.

Aliases exist for one main purpose: **moving a type between packages without
breaking anyone.** Leave `type OldName = newpkg.NewName` behind and existing
code keeps compiling against either name during a gradual migration. Since Go
1.24 aliases can also take type parameters (`type Set[T comparable] =
map[T]struct{}`), which makes them useful for naming long generic types.

Outside those cases, an alias is usually the wrong tool — it adds a name without
adding safety.

```ascii
type Celsius float64        type MyByte = byte

  Celsius  ─┐ distinct        MyByte ──┐
  float64  ─┘ types           byte   ──┴─► the SAME type
  conversion required          no conversion; interchangeable
  can carry methods            cannot carry new methods
```

## 4. Choosing between them

| Want | Use |
|---|---|
| Prevent mixing two values with the same representation | defined type |
| Attach methods, satisfy an interface | defined type |
| Rename or relocate a type without breaking callers | alias |
| Shorten a long generic type name (1.24+) | alias |
| A domain concept (`UserID`, `Celsius`, `Weekday`) | defined type |

## Gotchas

- **The `=` is easy to miss** when reading, and changes everything.
- **A defined type starts with no methods**, even when its underlying type has
  many.
- **`String()` that formats the receiver with `%v` recurses forever.** Convert
  first.
- **You cannot declare methods on an alias of a type from another package** —
  it is that package's type.
- **Untyped constants still convert freely**: `var c Celsius = 100` works
  without a conversion, so defined types do not stop literal mistakes.
- **Conversion is not free of meaning**, only of cost: `float64(celsius)`
  discards the unit the type was carrying.

## The exercises

- **typealias1** — a `Celsius` will not pass as a `float64`; convert at the call
  site.
- **typealias2** — attach `String()` to `Celsius`, converting inside to avoid
  recursion.
- **typealias3** — make `MyByte` a true alias so it passes as a `byte` untouched.

## Source references

- [Go spec: Type definitions](https://go.dev/ref/spec#Type_definitions) ·
  [Alias declarations](https://go.dev/ref/spec#Alias_declarations)
- [Go blog: Alias names](https://go.dev/blog/alias-names) — the package-migration
  use case
- [Go 1.24 release notes: generic type aliases](https://go.dev/doc/go1.24#language)
- [pkg.go.dev: fmt.Stringer](https://pkg.go.dev/fmt#Stringer)

**End of the Types & Methods tier.** Next:
[anonymous_functions](../anonymous_functions/) and [errors](../errors/) — values
that carry behaviour, and the interface Go returns instead of throwing.

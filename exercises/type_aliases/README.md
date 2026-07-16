# Type Aliases & Defined Types

Go distinguishes a *defined type* (`type Celsius float64`), which is a brand-new
type that needs explicit conversion, from a *type alias* (`type Byte = byte`),
which is merely a second name for an existing type. Knowing the difference
prevents accidental unit mix-ups and explains why some conversions are required
and others are not.

- **typealias1** — defined types are distinct; convert with `float64(x)`
- **typealias2** — attaching a `String()` method to a defined type
- **typealias3** — a true alias (`= T`) is the same type, no conversion needed

## Resources

- [Go spec: Type definitions](https://go.dev/ref/spec#Type_definitions)
- [Go spec: Alias declarations](https://go.dev/ref/spec#Alias_declarations)
- [Go blog: Type aliases](https://go.dev/blog/alias-names)

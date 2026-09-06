# Structs

A struct is a fixed set of named fields laid out in one block of memory. That is
the whole idea, and it is the main way Go models data — there are no classes,
no inheritance, and no constructors. A type is its fields; behaviour is attached
separately with methods; and reuse comes from **embedding** one struct in
another rather than from a hierarchy.

These three exercises build a `Person`: declare its fields, embed a second
struct into it, and attach a method.

## 1. Declaring, constructing, comparing

```go
type Person struct {
    name string
    age  int
}

p := Person{name: "John", age: 42}   // keyed — always prefer this
q := Person{"John", 42}              // positional — brittle, avoid
var z Person                          // zero value: name "", age 0
```

Keyed literals survive a field being added or reordered; positional ones break
silently or loudly. The zero value is every field's zero value, which is why a
`Person` needs no constructor to be usable — Go's convention is a `NewX`
function only when construction has real work to do.

Field names follow the same export rule as everything else: `name` is package
private, `Name` is visible to other packages. That single letter is what decides
whether `encoding/json` can see the field, too.

Structs are **comparable** with `==` when all their fields are, and they are
**values**: assigning or passing one copies every field.

## 2. What the memory looks like

Fields are laid out in declaration order, with padding inserted to satisfy each
field's alignment:

```ascii
type S struct {          bytes
    a bool               [a][pad pad pad pad pad pad pad]
    b int64              [........ b ........]
    c bool               [c][pad pad pad pad pad pad pad]
}                        = 24 bytes

type T struct {          reordered by size
    b int64              [........ b ........]
    a bool               [a][c][pad pad pad pad pad pad]
    c bool
}                        = 16 bytes
```

Same three fields, same types, different size — the compiler does not reorder
fields for you. This rarely matters, and occasionally matters a lot (millions of
instances, cache pressure). `unsafe.Sizeof` and `fieldalignment` in
`go vet -vettool` will tell you when it does.

## 3. Embedding: composition without inheritance

```go
type ContactDetails struct {
    phone string
}

type Person struct {
    name string
    age  int
    ContactDetails          // embedded: no field name
}

p := Person{name: "John", ContactDetails: ContactDetails{phone: "+01"}}
p.phone                     // promoted — shorthand for p.ContactDetails.phone
```

An embedded field is a real field whose **name is its type name**. Its fields
and methods are *promoted*: reachable on the outer struct as if they were
declared there. `structs2` is exactly this — note that the literal still
initialises it by type name, because that is the field's name.

Promotion is a lookup rule, not inheritance. There is no subtyping: a `Person`
is not a `ContactDetails`, cannot be assigned to one, and nothing is
overridden. If the outer struct declares a `phone` field of its own, the shallow
one wins and the inner is still reachable at `p.ContactDetails.phone`. Ambiguous
promotions at the same depth are simply not promoted — you get a compile error
only if you use the ambiguous name.

Embedding *interfaces* is how types get default behaviour cheaply (a struct
embedding `io.Reader` satisfies `io.Reader` by delegating), which the
`interfaces` chapter picks up.

## 4. Methods make it a type worth having

```go
func (p Person) FullName() string {
    return p.firstName + " " + p.lastName
}
```

The receiver `(p Person)` is what makes this a method rather than a function.
Methods live outside the struct definition, can be spread across files in the
package, and are the reason a struct can satisfy an interface. The choice
between a value and a pointer receiver is the whole subject of the `methods`
chapter — and the first thing to get right about any type you design.

## 5. Tags, and the reflection-facing side

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}
```

A tag is a string literal attached to a field, invisible to the compiler and
read at run time through reflection. `encoding/json`, database mappers, and
validators all use them. They are opaque text: a typo (`jsonn:"name"`) fails
silently, so `go vet`'s `structtag` check is worth having on.

## Gotchas

- **Positional literals break** when fields change. Use keyed literals.
- **A struct copies on assignment and on every call.** For large structs, or any
  struct with a mutex, pass a pointer.
- **Never copy a struct containing a `sync.Mutex`** — the copy gets its own lock.
- **Unexported fields are invisible to `encoding/json`** and to other packages;
  the capital letter is the API.
- **Comparability is all-or-nothing**: one slice, map, or func field makes the
  whole struct uncomparable, and `==` stops compiling.
- **Embedding is not inheritance.** No overriding, no subtyping, no virtual
  dispatch — just promoted names.
- **`struct{}{}` is the zero-size value**, the idiomatic set element
  (`map[string]struct{}`) and channel signal.

## The exercises

- **structs1** — declare the fields the test constructs.
- **structs2** — embed `ContactDetails` so `person.phone` is promoted rather
  than nested.
- **structs3** — attach a `FullName()` method returning the two names with a
  single space.

## Source references

- [Go spec: Struct types](https://go.dev/ref/spec#Struct_types) ·
  [Selectors](https://go.dev/ref/spec#Selectors) — the promotion rules in full
- [Effective Go: Embedding](https://go.dev/doc/effective_go#embedding)
- [Go blog: JSON and Go](https://go.dev/blog/json) — what tags are for
- [pkg.go.dev: unsafe.Sizeof](https://pkg.go.dev/unsafe#Sizeof) — measuring
  layout and padding
- [A Tour of Go: Structs](https://go.dev/tour/moretypes/2)

**Next: [pointers](../pointers/) →** — how a function reaches the struct the
caller is holding instead of a copy of it.

# Enums (iota)

Go has no `enum` keyword. What it has instead is a defined integer type plus a
`const` block using `iota`, and the combination gets you most of what an enum
provides: a distinct type the compiler will not confuse with a plain `int`, a
set of named values, and — with a `String()` method — readable output.

What it does **not** get you is exhaustiveness. Nothing stops `Weekday(42)`, and
nothing warns when a `switch` misses a case. Knowing where the guarantee stops
is the point of this chapter.

## 1. `iota`, precisely

```go
type Weekday int

const (
    Sunday Weekday = iota   // 0
    Monday                  // 1  — the expression repeats
    Tuesday                 // 2
    // …
    Saturday                // 6
)
```

`iota` is the index of the **ConstSpec** (the line) within its `const` block,
starting at 0. A line with no `= …` repeats the previous expression, so
`Monday` is implicitly `Weekday = iota` evaluated at line index 1. That
repetition is what makes `enums1` work with a single `= iota` at the top.

Two rules to keep straight:

- **`iota` resets to 0 in every `const` block**, not every file or package.
- **It counts lines, not constants.** `A, B = iota, iota` puts both on line 0.

Skipping and shifting fall out of the same rule:

```go
const (
    _  = iota             // skip 0 with the blank identifier
    KB = 1 << (10 * iota) // 1 << 10
    MB                    // 1 << 20
    GB                    // 1 << 30
)

const (
    Read Permission = 1 << iota   // 1
    Write                         // 2
    Execute                       // 4  — a bit flag set
)
```

## 2. Why the named type matters

```go
type Weekday int
func schedule(d Weekday) { … }

schedule(3)          // compiles — 3 is an untyped constant
schedule(someInt)    // does NOT compile — int is not Weekday
```

Declaring `type Weekday int` makes a **distinct type**: a plain `int` variable
cannot be passed where a `Weekday` is expected, so unit mix-ups are caught at
compile time. Untyped constants still convert freely, which is why the literal
`3` slips through — a real limit, not a bug.

The type also gives you somewhere to hang methods, which is the next step.

## 3. `String()` for readable values

```go
func (c Color) String() string {
    switch c {
    case Red:   return "Red"
    case Green: return "Green"
    case Blue:  return "Blue"
    }
    return fmt.Sprintf("Color(%d)", int(c))
}
```

This satisfies `fmt.Stringer`, so `fmt.Println(Red)` prints `Red` instead of
`0` — `enums2`. Two details worth copying: handle the **unknown** value rather
than returning `""`, and convert the receiver (`int(c)`) inside the fallback,
because formatting `c` itself with `%d`… is fine, but formatting it with `%v`
would call `String()` again and recurse until the stack dies.

Writing these by hand is fine for three values. For a long list, generate it:

```sh
go install golang.org/x/tools/cmd/stringer@latest
//go:generate stringer -type=Color
```

## 4. What the compiler will not do for you

- **No range checking.** `Color(99)` is a valid `Color`. Validate at the
  boundary — when parsing input — with a `switch` or a lookup table.
- **No exhaustiveness checking.** Adding `Yellow` will not make any existing
  `switch` fail to compile. Linters such as `exhaustive` add that check if you
  want it.
- **The zero value is the first constant**, so whichever name you list first is
  what an unset field means. Many codebases reserve index 0 for `Unknown`
  precisely so an uninitialised value is not silently valid.

For values that cross a wire or a database, prefer explicit numbers over
positional `iota` — reordering the block silently changes every stored value.

## Gotchas

- **`iota` counts ConstSpec lines**, so a blank line or a comment does not
  advance it, but a line with two names still counts once.
- **It resets in each `const` block.**
- **The zero value is meaningful whether you meant it to be or not.**
- **`String()` implemented with `%v` on the receiver recurses forever.**
- **Reordering an `iota` block renumbers everything** — dangerous once values
  are persisted.
- **A `String()` method on the value type is not in `*T`'s way**, but a pointer
  receiver here would mean `fmt.Println(Red)` does not use it. Use a value
  receiver.

## The exercises

- **enums1** — complete the weekday block; one `= iota` at the top numbers the
  rest.
- **enums2** — add `String()` so each colour prints its name.

## Source references

- [Go spec: Iota](https://go.dev/ref/spec#Iota) ·
  [Constant declarations](https://go.dev/ref/spec#Constant_declarations)
- [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
- [pkg.go.dev: stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) —
  generating the `String()` method
- [Go by Example: Enums](https://gobyexample.com/enums)

**Next: [type_aliases](../type_aliases/) →** — why `type Celsius float64` is a
new type at all, and what the `=` form does instead.

## enums2 — give an enum readable names

```go
type Color int
const ( Red Color = iota; Green; Blue )

func (c Color) String() string {
    switch c {
    case Red:   return "Red"
    case Green: return "Green"
    case Blue:  return "Blue"
    default:    return "Unknown"
    }
}
```

**Why it works**

- The constants are just `int`s (0, 1, 2). Adding a `String()` method (satisfying
  `fmt.Stringer`) makes them print as names instead of numbers.

**Key detail:** always include a `default` — an out-of-range `Color(99)` should return
`"Unknown"`, not fall through to nothing. For real projects the `stringer` tool
(`go generate`) writes this method for you from the const block.

**References**

- Go by Example — Enums / Iota: https://gobyexample.com/enums
- Effective Go — Constants: https://go.dev/doc/effective_go#constants

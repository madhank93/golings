## switch3 — map values to results

```go
func weekDay(day int) string {
    switch day {
    case 0:
        return "Sunday"
    case 1:
        return "Monday"
    // ...
    default:
        return ""
    }
}
```

**Why it works**

- One value, seven outcomes: each case compares `day` with `==` and returns. The
  same logic as seven `if`s, with the branching structure visible at a glance.

**Common mistake**

- Leaving out `default` in a function that must return something. Every path
  needs a `return`, and "the switch covers 0–6" is not something the compiler
  can verify — it will reject the function as missing a return.

**Key detail:** for a dense integer-to-string mapping like this one, a package-level
slice (`var days = [...]string{"Sunday", …}`) with a range check is often
clearer and faster. Reach for `switch` when the cases are irregular, and for a
table when they are an index.

**See also:** switch1 · enums2 (`iota` plus a `String()` method — the typed
version of this) · if2 · the [chapter](../README.md)

**References**

- Go spec — Expression switches: https://go.dev/ref/spec#Expression_switches
- Go by Example — Switch: https://gobyexample.com/switch

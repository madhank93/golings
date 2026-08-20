## switch2 — the conditionless switch needs `default`

```go
switch {
case 0 > 1:
    fmt.Println("zero is greater than one")
default:
    fmt.Println("one is greater than zero")
}
```

**Why it works**

- With no expression after `switch`, every case is a boolean condition. `case:`
  with nothing after it has nothing to evaluate; the catch-all branch is spelled
  `default`.

**Under the hood**

- The conditionless form is shorthand for `switch true`, so each case is
  compared against `true` in order. That makes it the direct replacement for an
  `if / else if` ladder — with the conditions in a column instead of drifting
  right one indent at a time.

**Common mistake**

- Assuming `default` has to come last. It can appear anywhere in the case list
  and still runs only when nothing else matched — though last is the
  conventional place.

**Key detail:** a switch may also run a statement first —
`switch v := compute(); { case v > 100: … }` — scoping `v` to the switch.

**See also:** switch1 (expression form) · if2 (the ladder this replaces) ·
select1 (a switch-shaped statement that picks at random) ·
the [chapter](../README.md)

**References**

- Go spec — Expression switches: https://go.dev/ref/spec#Expression_switches
- Effective Go — Switch: https://go.dev/doc/effective_go#switch

## switch1 — switch on a value

```go
status := "open"
switch status {
case "open":
    ...
case "closed":
    ...
}
```

**Why it works**

- A `switch status` compares `status` against each `case` value. The broken code
  wrote `switch { case "open": ... }` — a conditionless switch expects each case
  to be a **boolean**, not a bare string, so `case "open"` was a type error.

**Key detail:** Go cases **do not fall through** by default — the matching case runs
and the switch ends (no `break` needed). Use the `fallthrough` keyword only when
you explicitly want the next case to run too.

**References**

- Go by Example — Switch: https://gobyexample.com/switch

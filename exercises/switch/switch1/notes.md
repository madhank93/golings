## switch1 — switch on a value

```go
status := "open"
switch status {
case "open":
    fmt.Println("status is open")
case "closed":
    fmt.Println("status is closed")
}
```

**Why it works**

- Naming the value after `switch` makes each case a value to compare it with.
  The broken version omitted `status`, which turns it into a conditionless
  switch — and `"open"` is not a boolean, so it cannot compile.

**Under the hood**

- A conditionless `switch` is `switch true`: every case must be a boolean
  expression. With an expression present, cases are compared using `==` in
  source order and the first match wins — so any comparable type works, strings
  and structs included.

**Common mistake**

- Adding `break` at the end of each case out of C habit. Go cases never fall
  through; the `break` does nothing, and inside a loop it breaks the *switch*
  rather than the loop.

**Key detail:** one case can list several values — `case "closed", "archived":` —
and a `switch` with no matching case and no `default` simply does nothing.

**See also:** switch2 (the conditionless form) · switch3 (mapping values) ·
if2 · the [chapter](../README.md)

**References**

- Go spec — Expression switches: https://go.dev/ref/spec#Expression_switches
- Go by Example — Switch: https://gobyexample.com/switch

## switch3 — map a number to a name, with a default

```go
switch day {
case 0:
    return "Sunday"
// ...
default:
    return "Unknown"
}
```

**Why it works**

- Each `case` returns the matching weekday; `default` catches any number outside
  0–6, so the function always returns something.

**Key detail:** always give a `switch` a `default` when the input isn't guaranteed to
be one of the listed cases — it's your total-coverage safety net. Cases can also
list several values at once (`case 6, 0:`) when they share a result.

**References**

- Go by Example — Switch: https://gobyexample.com/switch

## cli1 — parse flags with a FlagSet

```go
fs := flag.NewFlagSet("greet", flag.ContinueOnError)
name := fs.String("name", "world", "who to greet")
count := fs.Int("count", 1, "how many times")
if err := fs.Parse(args); err != nil {
    return "", 0
}
return *name, *count
```

**Why it works**

- A `FlagSet` parses an **explicit slice**, so the function can be called from a
  test with any arguments. The definers return pointers because the values do not
  exist until `Parse` runs.

**Under the hood**

- The package-level `flag.String`/`flag.Parse` register on `flag.CommandLine`, a
  global reading `os.Args`. That global cannot be given test input, and
  registering the same flag name twice **panics** — so a helper using it cannot
  be called from two tests.

**Common mistake**

- Choosing `flag.ExitOnError` (the default for `flag.CommandLine`). A bad-input
  test then calls `os.Exit(2)` and kills the test binary. `ContinueOnError`
  returns the error instead, which is what makes this testable.

**Key detail:** `fs.StringVar(&cfg.Name, "name", ...)` binds into a struct field
instead of returning a pointer — usually tidier for a real tool. Custom types
implement `flag.Value` (`String`/`Set`) and register with `fs.Var`.

**See also:** cli2 (positional args) · di1 (injecting output) · errors1 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — flag.FlagSet: https://pkg.go.dev/flag#FlagSet
- Go by Example — Command-Line Flags: https://gobyexample.com/command-line-flags

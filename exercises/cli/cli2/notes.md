## cli2 — positional arguments

```go
if err := fs.Parse(args); err != nil {
    return ""
}
return fs.Arg(0) // "" when there is none
```

**Why it works**

- After `Parse`, the leftovers are the positional arguments. `fs.Args()` returns
  them all and `fs.Arg(0)` returns the first — or `""` when there is none, so no
  bounds check is needed.

**Under the hood**

- Parsing **stops at the first non-flag argument**. In `["-v", "build", "./..."]`,
  `build` ends the flag section, so anything flag-shaped after it is treated as a
  positional. That is deliberate — it is what lets a subcommand have its own
  flags — but it means flags must precede positionals unless you parse in stages.

**Common mistake**

- Writing `-v true` for a boolean. Booleans are `-v` or `-v=true`; the spaced
  form would make `true` a positional argument and silently leave the flag false.

**Key detail:** subcommands are just nested FlagSets — switch on `fs.Arg(0)`, then
parse `fs.Args()[1:]` with a fresh `FlagSet` for that command. That is how `go`,
`docker`, and `kubectl` are shaped, with no dependency.

**See also:** cli1 (defining flags) · di1 (testable `run(args, stdout)`) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — FlagSet.Args: https://pkg.go.dev/flag#FlagSet.Args
- Go by Example — Subcommands: https://gobyexample.com/command-line-subcommands

## cli2 — positional arguments

```go
fs := flag.NewFlagSet("cmd", flag.ContinueOnError)
fs.Bool("v", false, "verbose output")
fs.Parse(args)
fs.Arg(0) // first NON-flag argument, e.g. "build"
```

**Why it works**

- After `Parse`, the leftover **positional** arguments (everything that isn't a
  flag) are available via `fs.Args()` / `fs.Arg(i)`. Given `["-v", "build", "./..."]`,
  `Arg(0)` is `"build"`.

**Nuance:** `Arg(i)` returns `""` for an out-of-range index (no panic), so
`firstArg(["-v"])` safely yields `""`. Flag parsing **stops** at the first
non-flag argument — this is how CLIs separate global flags from a subcommand and
its args (`tool -v build ./...`).

**References**

- pkg.go.dev — flag.FlagSet.Args: https://pkg.go.dev/flag#FlagSet.Args

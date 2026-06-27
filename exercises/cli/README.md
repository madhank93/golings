# CLI (flag)

Go's standard `flag` package parses command-line flags — no third-party library
needed. Using a `flag.NewFlagSet` (instead of the package-global functions) lets
you parse an explicit argument slice, which is exactly what makes a CLI testable.

- **cli1** — define and parse flags on a `FlagSet`
- **cli2** — read the positional arguments left after the flags with `fs.Args()`/`fs.Arg(0)`

## Resources

- [flag package](https://pkg.go.dev/flag)
- [Command-line arguments](https://gobyexample.com/command-line-flags)

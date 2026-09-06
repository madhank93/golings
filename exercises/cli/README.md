# CLI (flag)

Go's `flag` package parses command-line flags, and for most tools it is all you
need — no Cobra, no urfave/cli. What makes it worth a chapter is a structural
point that generalises well beyond flags: using a **`FlagSet`** instead of the
package-level functions turns argument parsing from untestable global state into
an ordinary function you can call with a slice.

## 1. `FlagSet` over the globals

```go
func parseGreeting(args []string) (string, int) {
    fs := flag.NewFlagSet("greet", flag.ContinueOnError)
    name := fs.String("name", "world", "who to greet")
    count := fs.Int("count", 1, "how many times")

    if err := fs.Parse(args); err != nil {
        return "", 0
    }
    return *name, *count
}
```

`cli1`. Compare with `flag.String(...)` + `flag.Parse()`, which register on
`flag.CommandLine` — a package global reading `os.Args`. That version cannot be
called twice (duplicate flag registration panics), cannot be given test input,
and pollutes state other tests will see.

The definers return **pointers**, because the value does not exist until
`Parse` runs. The `fs.StringVar(&cfg.Name, "name", …)` form binds into a struct
field instead, which is usually tidier for a real tool.

The error behaviour is chosen at construction:

| Mode | On a bad flag |
|---|---|
| `flag.ContinueOnError` | `Parse` returns an error — the testable choice |
| `flag.ExitOnError` | prints usage and calls `os.Exit(2)` |
| `flag.PanicOnError` | panics |

`ExitOnError` is why some CLIs are impossible to test: a bad-input test kills
the test binary.

## 2. Positional arguments

```go
fs.Parse(args)
fs.Args()     // everything after the flags
fs.Arg(0)     // "" when there is none — no bounds check needed
fs.NArg()     // how many
```

`cli2`. Parsing **stops at the first non-flag argument**, which is the rule that
surprises people: in `["-v", "build", "./..."]`, `build` ends the flag section,
so a `-x` after it is a positional argument, not a flag. That is deliberate — it
is what lets `go test -v ./... -run X` work — but it means flags must precede
positionals unless you parse in stages.

`fs.Arg(0)` returning `""` rather than panicking is what makes the "did the user
give me a subcommand?" check a one-liner.

## 3. Subcommands are just nested FlagSets

```go
switch fs.Arg(0) {
case "serve":
    serveFS := flag.NewFlagSet("serve", flag.ContinueOnError)
    port := serveFS.Int("port", 8080, "listen port")
    serveFS.Parse(fs.Args()[1:])       // the rest, after the subcommand
    return serve(*port)
case "migrate":
    …
default:
    fs.Usage()
    return errUsage
}
```

This is how `go`, `docker`, and `kubectl` are structured, and the standard
library covers it without a dependency. Reach for Cobra when you want generated
completions, deep nesting, and man pages — not for two subcommands.

## 4. What `flag` supports, and what it does not

Supported: `-flag`, `--flag`, `-flag=value`, `-flag value` (for non-booleans),
`-help`/`-h`, and custom types via the `flag.Value` interface:

```go
type levels []string
func (l *levels) String() string     { return strings.Join(*l, ",") }
func (l *levels) Set(s string) error { *l = append(*l, s); return nil }

fs.Var(&lv, "level", "may be repeated")
```

Not supported: single-dash grouping (`-abc` for `-a -b -c`), required flags, or
mutually exclusive sets. Booleans are the one syntax trap — `-v true` does *not*
work; it is `-v` or `-v=true`, because `-v true` would make `true` a positional.

## 5. Making a whole CLI testable

The same move as the `FlagSet` scales up to `main`:

```go
func main() {
    if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func run(args []string, stdout, stderr io.Writer) error { … }
```

`main` does nothing but wire and exit; `run` takes its arguments and its output
streams (the `dependency_injection` chapter, again) and returns an error. Now a
test calls `run([]string{"-name", "Go"}, &buf, io.Discard)` and asserts on the
buffer — the entire tool under test, no subprocess.

Exit codes belong in `main` alone: `os.Exit` skips every deferred call, so
calling it from inside your logic loses cleanup.

## Gotchas

- **`flag.Parse()` uses globals and `os.Args`** — untestable. Use a `FlagSet`.
- **Definers return pointers**; dereference after `Parse`.
- **Parsing stops at the first positional**, so flags must come first.
- **Boolean flags need `-v` or `-v=true`**, never `-v true`.
- **`ExitOnError` terminates the test binary** on bad input.
- **`os.Exit` skips defers** — only call it from `main`.
- **Registering the same flag name twice panics**, which is why a global-flag
  helper cannot be called from two tests.

## The exercises

- **cli1** — define flags on a `FlagSet`, parse an explicit slice, return the
  values.
- **cli2** — read the first positional argument after the flags.

## Source references

- [pkg.go.dev: flag](https://pkg.go.dev/flag) ·
  [flag.Value](https://pkg.go.dev/flag#Value) ·
  [FlagSet.Args](https://pkg.go.dev/flag#FlagSet.Args)
- [Go by Example: Command-Line Flags](https://gobyexample.com/command-line-flags) ·
  [Subcommands](https://gobyexample.com/command-line-subcommands)
- [Mat Ryer: How I write HTTP services / testable main](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/)

**End of the Building Applications tier.** Next:
[testing_advanced](../testing_advanced/) — the tools that keep all of this
honest.

## cli1 — parse flags

```go
fs := flag.NewFlagSet("greet", flag.ContinueOnError)
name := fs.String("name", "world", "")
count := fs.Int("count", 1, "")
fs.Parse(args)
// use *name, *count
```

**Why it works**

- Each `fs.String`/`fs.Int` registers a flag with a **default** and returns a
  **pointer**; `fs.Parse(args)` fills those pointers from the arguments. Missing
  flags keep their defaults (`world`, `1`).

**Key detail:** the flag values are pointers, so you dereference (`*name`) after
`Parse`. Using a `flag.NewFlagSet` (instead of the global `flag.CommandLine`) makes
parsing testable — you feed it an explicit `args` slice. `ContinueOnError` returns
an error instead of calling `os.Exit` on a bad flag.

**References**

- pkg.go.dev — flag: https://pkg.go.dev/flag
- Go by Example — Command-Line Flags: https://gobyexample.com/command-line-flags

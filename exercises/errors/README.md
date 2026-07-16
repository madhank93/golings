# Errors

Go treats errors as ordinary values returned from functions, not exceptions.
These exercises cover the idiomatic toolkit:

- returning and checking `error` values (`errors.New`)
- wrapping with `fmt.Errorf("...: %w", err)` and matching with `errors.Is`
- custom error types and extracting them with `errors.As`
- `panic`/`recover` and why it is a last resort, not flow control
- deep wrap chains: `%w` keeps `errors.Is` working, `%v` breaks it

## Resources

- [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
- [Go blog: Errors are values](https://go.dev/blog/errors-are-values)
- [Go blog: Working with Errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [Go by Example: Errors](https://gobyexample.com/errors)

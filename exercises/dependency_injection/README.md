# Dependency Injection

Dependency injection (DI) means giving a function or type its collaborators
instead of hard-coding them. In Go this needs no framework — just **interfaces**
and plain parameters. The payoff is testability: inject a fake collaborator and
assert on what your code does.

- **di1** — inject an `io.Writer` instead of printing to stdout; test with `bytes.Buffer`
- **di2** — inject behaviour through a small interface (a `Clock`) to make time deterministic
- **di3** — constructor/field injection: a struct holds a dependency interface and delegates to it

## Resources

- [Learn Go with Tests — Dependency Injection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)
- [Effective Go — Interfaces](https://go.dev/doc/effective_go#interfaces)

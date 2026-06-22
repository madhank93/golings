# Advanced testing

Beyond a basic `TestXxx`, Go's testing toolkit covers subtests, property-based
fuzzing, and HTTP handler testing.

- **testadv1** — table-driven tests with `t.Run` subtests
- **testadv2** — fuzzing with `func FuzzXxx(*testing.F)` (seed corpus runs under `go test`; try `go test -fuzz=Fuzz`)
- **testadv3** — handler testing with `net/http/httptest`

## Resources

- [Go blog: Using subtests and sub-benchmarks](https://go.dev/blog/subtests)
- [Go: Fuzzing tutorial](https://go.dev/doc/tutorial/fuzz)
- [pkg.go.dev: net/http/httptest](https://pkg.go.dev/net/http/httptest)

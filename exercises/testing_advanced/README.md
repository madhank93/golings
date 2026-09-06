# Advanced testing

Every exercise in this repo has been run by `go test`, so the basics are already
familiar: a `TestXxx(*testing.T)` in a `_test.go` file, `t.Errorf` to fail and
continue, `t.Fatalf` to fail and stop. This chapter is the rest of the toolkit —
the four things that separate a test suite from a pile of tests.

## 1. Table-driven tests and subtests

```go
cases := []struct {
    name string
    in   int
    want string
}{
    {"multiple of three", 3, "Fizz"},
    {"multiple of five", 5, "Buzz"},
    {"multiple of fifteen", 15, "FizzBuzz"},
    {"plain number", 7, "7"},
}

for _, tc := range cases {
    t.Run(tc.name, func(t *testing.T) {
        if got := fizzbuzz(tc.in); got != tc.want {
            t.Errorf("fizzbuzz(%d) = %q, want %q", tc.in, got, tc.want)
        }
    })
}
```

`testadv1`. This is *the* Go testing idiom. Adding a case is one line, each
subtest reports independently, and `go test -run 'TestFizzBuzz/multiple_of_five'`
runs exactly one.

The `name` field is not decoration — without `t.Run`, a failure says "test
failed" and you go hunting; with it, the failing case names itself. And a good
failure message includes **input, got, and want**, in that order, so the message
alone is enough to act on.

`t.Parallel()` inside the subtest runs cases concurrently. Since Go 1.22 the
loop variable is per-iteration, so the old `tc := tc` line is no longer needed.

## 2. Fuzzing

```go
func FuzzReverse(f *testing.F) {
    for _, seed := range []string{"hello", "", "abc", "Hello, 世界"} {
        f.Add(seed)
    }
    f.Fuzz(func(t *testing.T, s string) {
        rev := reverse(reverse(s))
        if rev != s {
            t.Errorf("round trip: got %q, want %q", rev, s)
        }
        if !utf8.ValidString(reverse(s)) {
            t.Errorf("reverse produced invalid UTF-8")
        }
    })
}
```

`testadv2`. A fuzz test states a **property** that must hold for *every* input,
and the toolchain generates inputs trying to break it. The seed corpus (`f.Add`)
runs during an ordinary `go test`; `go test -fuzz=FuzzReverse` starts generating.

The properties above are the classic pair: a round trip that must return the
original, and an invariant (still valid UTF-8) that the byte-reversing
implementation violates on `"Hello, 世界"` — the bug this exercise is built
around.

When fuzzing finds a failure it writes the input to `testdata/fuzz/…`, which
becomes a permanent regression test. Fuzzing suits parsers, encoders,
normalisers — anything taking untrusted input.

## 3. Testing HTTP handlers

```go
req := httptest.NewRequest(http.MethodGet, "/ping", nil)
rec := httptest.NewRecorder()

pingHandler(rec, req)

res := rec.Result()
if res.StatusCode != http.StatusOK { … }
body, _ := io.ReadAll(res.Body)
```

`testadv3`. `httptest.NewRequest` builds a request without a network;
`NewRecorder` captures what the handler wrote. Calling the handler directly is
faster than a server and tests exactly the unit you mean.

The sibling is `httptest.NewServer`, which starts a **real** server on a
loopback port — the right tool when testing a *client* (`http1`), because it
exercises real request serialisation.

## 4. Benchmarks

```go
func BenchmarkSumSlice(b *testing.B) {
    data := make([]int, 1000)
    for i := range data {
        data[i] = i
    }
    for b.Loop() {          // Go 1.24+ form
        sumSlice(data)
    }
}
```

`testadv4`. Run with `go test -bench=. -benchmem`. `b.Loop()` (Go 1.24) replaced
`for i := 0; i < b.N; i++` and fixes two long-standing problems: setup before
the loop is excluded automatically without `b.ResetTimer()`, and the compiler
cannot optimise away the call, so the old `result = …` sink is unnecessary.

Read the output as ns/op, B/op, allocs/op — and compare runs with
`benchstat`, because a single number is noise.

## 5. The helpers worth knowing

| Call | Does |
|---|---|
| `t.Helper()` | failures point at the caller, not the helper |
| `t.Cleanup(fn)` | runs after the test, including subtests — better than `defer` |
| `t.TempDir()` | temp directory, auto-removed |
| `t.Setenv(k, v)` | env var restored afterwards (disables parallel) |
| `t.Skip()` / `testing.Short()` | skip slow tests under `-short` |
| `TestMain(m)` | one-time setup/teardown for the package |

Two things Go deliberately does not have: assertion libraries (the standard idiom
is `if got != want { t.Errorf(...) }`) and mocking frameworks (`mock1`). Both
are choices, not omissions.

Always run concurrent code with `-race` (`mise run test` does), and note that
`go test` caches results — a passing test that "did not run" was cached; `-count=1`
forces it.

## Gotchas

- **`t.Fatal` from a non-test goroutine does not work** — it calls
  `runtime.Goexit` on the wrong goroutine. Send the failure back over a channel.
- **Missing `t.Helper()`** makes every failure point at your assertion helper.
- **`t.Parallel()` subtests run after the parent returns**, so anything the
  parent deferred has already happened — use `t.Cleanup`.
- **A fuzz target must be deterministic**; randomness inside it makes failures
  unreproducible.
- **`-race` roughly doubles runtime** and is still worth it in CI.
- **Benchmarks without `b.Loop`** need `b.ResetTimer()` after setup.
- **Test caching** can hide flakiness; `-count=1` when in doubt.

## The exercises

- **testadv1** — implement FizzBuzz against a table of `t.Run` subtests.
- **testadv2** — fix a byte-reversing function so the fuzz properties hold for
  multibyte input.
- **testadv3** — write a handler and assert on it with `httptest`.
- **testadv4** — implement the function a `b.Loop()` benchmark measures.

## Source references

- [Go blog: Using subtests and sub-benchmarks](https://go.dev/blog/subtests) ·
  [Fuzzing tutorial](https://go.dev/doc/tutorial/fuzz) ·
  [Testable Examples](https://go.dev/blog/examples)
- [pkg.go.dev: testing](https://pkg.go.dev/testing) ·
  [testing.B.Loop](https://pkg.go.dev/testing#B.Loop) ·
  [net/http/httptest](https://pkg.go.dev/net/http/httptest)
- [Go Wiki: TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
- [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat)

**Next: [profiling](../profiling/) →** — when the benchmark says it is slow and
you need to know why.

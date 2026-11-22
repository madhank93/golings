# Testing

Most exercises so far shipped a test that you made pass. This chapter flips it:
you write the test's side. Go's testing kit is small. `go test` finds `_test.go`
files, runs every `TestXxx(t *testing.T)` in them, and reports a test as failed
when it calls `t.Errorf` or `t.Fatalf`. There are no assertion helpers, so a
check is just an `if` and a message.

## 1. What `go test` runs

```go
// file: math_test.go
package math_test

import "testing"

func TestDouble(t *testing.T) {
    if got := double(4); got != 8 {
        t.Errorf("double(4) = %d, want 8", got)
    }
}
```

Three rules decide whether a function is a test:

- the file name ends in `_test.go`;
- the name starts with `Test` and the next letter is not lowercase (`TestDouble`, not `Testdouble`);
- it takes exactly one `*testing.T` and returns nothing.

Get the signature wrong and `go test` refuses to build the package. Get the name
wrong (`testDouble`) and it silently treats it as an ordinary function that
never runs.

## 2. Failing: `Errorf` vs `Fatalf`

| Call | Marks failed | Keeps running the test |
|---|---|---|
| `t.Errorf(...)` | yes | yes, and later checks still report |
| `t.Fatalf(...)` | yes | no, it stops this test immediately |

Use `Errorf` by default so one run reports every wrong value. Use `Fatalf` when
continuing makes no sense, usually because an error means the value you were
about to inspect is unusable:

```go
u, err := parseUser("ana:30")
if err != nil {
    t.Fatalf("parseUser error: %v", err) // u is nil; checking u.Name would panic
}
if u.Name != "ana" {
    t.Errorf("Name = %q, want %q", u.Name, "ana")
}
```

## 3. Messages: input, got, want

A failure message should let you act without opening the file. The convention
is `f(input) = got, want expected`:

```text
abs(-3) = -3, want 3
```

Put `got` before `want`. Every Go test reads that way, so readers know where to
look.

## 4. Testing the failure path

An error is often the correct result. A test for it checks that the expected
error came back, not that no error did:

```go
_, err := greet("")
if !errors.Is(err, ErrEmptyName) {
    t.Fatalf("greet(\"\") error = %v, want ErrEmptyName", err)
}
```

`errors.Is` (from the [errors](../errors/) chapter) also matches a wrapped
error, so the test keeps passing if `greet` later adds context with `%w`.

## Running tests

```sh
go test ./...                 # every package below here
go test -v ./...              # also print each test name and its result
go test -run TestGreet ./...  # only tests whose name matches the regexp
```

## Gotchas

- **A test that never fails is not a test.** Break the code once on purpose and
  watch the test go red before you trust it.
- **`t.Fatalf` inside a goroutine** does not stop the test. Call it only from
  the test's own goroutine.
- **`want` from the code under test.** Writing `want := abs(-3)` makes the
  test agree with whatever the code does. Write expected values out literally.

## The exercises

- **testing1**: fix a test signature so `go test` will run it.
- **testing2**: turn an unconditional failure into a real check.
- **testing3**: test an error path with `errors.Is` and `t.Fatalf`.

## Source references

- [pkg.go.dev: testing](https://pkg.go.dev/testing): `T`, `Errorf`, `Fatalf`, `Run`
- [Go Code Review Comments: Useful Test Failures](https://go.dev/wiki/CodeReviewComments#useful-test-failures)
- [Go Wiki: TestComments](https://go.dev/wiki/TestComments): got-before-want and other message conventions
- [`go help test`](https://pkg.go.dev/cmd/go#hdr-Test_packages) and [`go help testflag`](https://pkg.go.dev/cmd/go#hdr-Testing_flags)

**End of the Functions & Errors tier.** Next: [generics](../generics/), which
lets one function work across types without giving up the checking these
chapters relied on. Table tests, fuzzing and benchmarks come later in
[advanced testing](../testing_advanced/).

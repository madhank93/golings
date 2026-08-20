## mock2 — a stub drives both paths

```go
func WelcomeMessage(f Fetcher, id int) string {
    name, err := f.Fetch(id)
    if err != nil {
        return "Welcome, guest!"
    }
    return fmt.Sprintf("Welcome, %s!", name)
}
```

**Why it works**

- `Fetch` returns `(string, error)`, so the function has two paths. The fake
  returns a canned name for known ids and an error for anything else — one
  double, both branches covered.

**Under the hood**

- The error branch is the reason this pattern earns its keep. Reaching it against
  a real database would mean breaking the database on purpose; with a stub it is
  one unknown id. This is where injection pays for itself.

**Common mistake**

- Using `name` before checking `err`. When the error is non-nil the other results
  are meaningless — the value is whatever zero value the implementation
  returned, and formatting it produces `"Welcome, !"`.

**Key detail:** for a stub whose behaviour changes per test, hold a **function**
in the double — `type stubFetcher struct{ fetch func(int) (string, error) }` —
so each test sets its own without declaring a new type. And drop
`var _ Fetcher = fakeFetcher{}` in the test file so the double fails to compile
when the interface changes.

**See also:** mock1 (spies) · di3 (injection) · errors1 (the `(T, error)`
contract) · testadv3 (`httptest` instead of faking HTTP) ·
the [chapter](../README.md)

**References**

- Learn Go with Tests — Mocking: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
- pkg.go.dev — errors: https://pkg.go.dev/errors

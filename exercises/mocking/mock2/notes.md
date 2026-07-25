## mock2 — a stub / fake

```go
type fakeFetcher struct{ names map[int]string }
func (f fakeFetcher) Fetch(id int) (string, error) {
    if name, ok := f.names[id]; ok { return name, nil }
    return "", errors.New("not found")
}
```

**Why it works**

- A **stub/fake** returns **canned data** so you can test logic — including the
  **error path**. The fake returns a name for known ids and an error otherwise, so
  both `WelcomeMessage` branches ("Welcome, Go!" and "Welcome, guest!") are
  covered.

**Nuance:** spy (mock1) vs stub — a spy records *how it was called*; a stub
*controls what it returns*. Stubs make error handling testable without needing a
real failing dependency (a down database, a 500 response). A `fake` is a stub with
a bit of real behaviour (here, a map lookup).

**References**

- Learn Go with Tests — Mocking: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

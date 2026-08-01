## di2 — inject behaviour through an interface

```go
type Clock interface{ Now() time.Time }

func Greeting(c Clock) string {
    if c.Now().Hour() < 12 { return "Good morning" }
    return "Good evening"
}
type fixedClock struct{ t time.Time }
func (f fixedClock) Now() time.Time { return f.t }
```

**Why it works**

- Calling `time.Now()` directly makes `Greeting` untestable — its result depends on
  the wall clock. Injecting a `Clock` interface lets the test pass a `fixedClock`
  that returns any moment, so both the morning and evening branches are reachable.

**Key detail:** wrap **non-deterministic** dependencies (time, randomness, network,
filesystem) behind a small interface so tests can substitute a predictable fake.
Define the interface with only the method you actually use (`Now`), not the whole
`time` API.

**References**

- Learn Go with Tests — Dependency Injection: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

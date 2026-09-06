## di2 — inject a Clock

```go
func Greeting(c Clock) string {
    if c.Now().Hour() < 12 {
        return "Good morning"
    }
    return "Good evening"
}
```

**Why it works**

- The function reads the time from the injected `Clock` rather than calling
  `time.Now()` itself, so the test can supply `fixedClock{t: …}` and pin the
  hour. Production passes a clock that delegates to `time.Now`.

**Under the hood**

- `Clock` has one method, and the test double is three lines with a **value**
  receiver — nothing to record, so no pointer needed. The production
  implementation is equally small: `type realClock struct{}` with
  `func (realClock) Now() time.Time { return time.Now() }`.

**Common mistake**

- Calling `time.Now()` inside business logic and testing around it — a suite
  that passes all morning and fails after lunch, or one that computes the
  expected value with the same call it is testing and therefore asserts nothing.

**Key detail:** the same wrapper trick covers every non-deterministic dependency:
`rand`, UUIDs, `os.Getenv`, the filesystem. For *concurrent* time-dependent code
there is now a better tool — `testing/synctest` (synctest1) runs the real `time`
package against a fake clock, no injection required.

**See also:** di1 · di3 · synctest2 (virtual time) · mock2 (canned results) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — time.Time: https://pkg.go.dev/time#Time
- Learn Go with Tests — Dependency Injection: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection

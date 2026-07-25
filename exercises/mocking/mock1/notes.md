## mock1 — a spy

```go
type spyNotifier struct{ messages []string }
func (s *spyNotifier) Notify(msg string) { s.messages = append(s.messages, msg) }

Alert(spy, 30, 25)
// assert: spy.messages == ["too hot"]
```

**Why it works**

- A **spy** is a test double that **records** how it was called. `spyNotifier`
  captures every `Notify` message, so the test can assert that `Alert` fired
  exactly one "too hot" (and stayed silent when cool).

**Nuance:** use a spy to verify **interactions** — that a side effect happened —
when there's no return value to check. Because `Alert` depends on the `Notifier`
interface, swapping the real notifier for the spy needs no change to `Alert`
itself. In Go you usually hand-write these small doubles rather than use a
mocking framework.

**References**

- Learn Go with Tests — Mocking: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking

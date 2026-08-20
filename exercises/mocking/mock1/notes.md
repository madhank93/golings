## mock1 — a spy records the interaction

```go
func Alert(n Notifier, temp, threshold int) {
    if temp > threshold {
        n.Notify("too hot")
    }
}

type spyNotifier struct{ messages []string }

func (s *spyNotifier) Notify(msg string) { s.messages = append(s.messages, msg) }
```

**Why it works**

- `Alert` has no return value — the notification **is** the behaviour. A spy
  records every call, so the test can assert that the alert fired when the
  temperature was above the threshold and stayed silent when it was not.

**Under the hood**

- The spy needs a **pointer receiver** and must be passed as `&spyNotifier{}`.
  With a value receiver, `append` would write into a copy of the struct and the
  test would see an empty slice — a double that silently records nothing.

**Common mistake**

- Testing only the firing case. The silent case is what catches an inverted
  comparison (`<` instead of `>`) or a missing guard, and it costs three lines.

**Key detail:** hand-written doubles need no library. Ten lines in the `_test.go`
file give you exactly the recording you want — and because over-specifying is
annoying to write by hand, these tests tend to assert behaviour rather than call
sequences.

**See also:** mock2 (a stub returning data) · di3 (what gets injected) ·
testadv1 (table-driven cases) · the [chapter](../README.md)

**References**

- Learn Go with Tests — Mocking: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking
- pkg.go.dev — testing: https://pkg.go.dev/testing

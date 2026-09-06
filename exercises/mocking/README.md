# Mocking (test doubles)

Once a dependency arrives through an interface, a test can pass something else
in its place. In Go that something else is usually **ten lines of ordinary code
in the test file** — no mocking library, no code generation, no expectation DSL.

The two shapes cover almost everything: a **spy** records how it was called so
you can assert on the interaction, and a **stub** (or fake) returns canned
results so you can drive the logic, including the error paths that are otherwise
hard to reach.

## 1. The vocabulary, briefly

| Double | Does | Assert on |
|---|---|---|
| **Stub** | returns canned values | the result your code produced |
| **Spy** | records calls, may also stub | the calls that were made |
| **Fake** | a real but simplified implementation (in-memory store) | either |
| **Mock** | a double with pre-programmed *expectations* | verified automatically |

Go idiom leans on the first three. Full mocks — "expect `Save` to be called once
with these arguments, in this order" — couple a test to the implementation's
internals, and Go's tooling gives you no reason to prefer them.

## 2. A spy: assert on the interaction

Some behaviour has no return value; the effect *is* the call. Alerting is the
classic case:

```go
func Alert(n Notifier, temp, threshold int) {
    if temp > threshold {
        n.Notify("too hot")
    }
}

type spyNotifier struct{ messages []string }

func (s *spyNotifier) Notify(msg string) {
    s.messages = append(s.messages, msg)
}
```

`mock1`. The test passes `&spyNotifier{}` and then checks two things that matter
equally: that the alert **fired** when it should, and that it **stayed silent**
when it should not. The negative case is the one people skip, and the one that
catches an inverted comparison.

Note the pointer receiver — a spy has to record into itself, so it must be
passed as `&spyNotifier{}`. A value receiver would append to a copy and the test
would see nothing.

## 3. A stub: drive the logic, including failures

```go
type Fetcher interface {
    Fetch(id int) (string, error)
}

type fakeFetcher struct{ names map[int]string }

func (f fakeFetcher) Fetch(id int) (string, error) {
    if name, ok := f.names[id]; ok {
        return name, nil
    }
    return "", errors.New("not found")
}
```

`mock2`. The canned map lets one double serve both tests: a known id exercises
the happy path, an unknown id exercises the error branch. Reaching that error
branch against a real database would mean breaking the database on purpose —
which is exactly why the injection was worth doing.

For a function-shaped dependency, the stub can be a field holding a function,
which lets each test set its own behaviour without a new type:

```go
type stubFetcher struct {
    fetch func(id int) (string, error)
}
func (s stubFetcher) Fetch(id int) (string, error) { return s.fetch(id) }
```

## 4. Faking a big interface without implementing it

When the interface is large and a test only touches one method, **embed** it:

```go
type partialStore struct {
    Store              // embedded interface — nil
    saveFunc func(string)
}
func (p partialStore) Save(n string) { p.saveFunc(n) }
```

The embedded `Store` supplies the rest of the method set to satisfy the
compiler; calling any of them panics with a nil-pointer dereference, which is
the right outcome — the test asserted that only `Save` was used, and it will say
so loudly if that changes.

## 5. What to assert, and what not to

Test the **behaviour your code is responsible for**, not the shape of its
collaboration:

```go
// good — the outcome
if got := WelcomeMessage(f, 1); got != "Welcome, Go!" { … }

// good — the effect, when the effect IS the point
if len(spy.messages) != 1 { … }

// fragile — the internals
if spy.callCount != 3 || spy.lastArgs[0] != "…" { … }
```

A test that pins the exact call sequence fails on every refactor that keeps the
behaviour identical. That is the failure mode mocking libraries make easy, and
the reason hand-written doubles — which make over-specification annoying to
write — tend to age better.

Also worth knowing: for HTTP, do not fake the client. `net/http/httptest` gives
you a **real** server and a real recorder, which tests the request you actually
send (`testadv3`).

## Gotchas

- **A spy needs a pointer receiver** and must be passed as `&spy{}`, or its
  recording goes into a copy.
- **Test the silent path too** — "did not notify" is as much a behaviour as
  "notified".
- **Doubles belong in `_test.go` files.** A fake shipped in the package is API
  nobody asked for.
- **Don't assert on call counts and argument order** unless that ordering is the
  contract.
- **`var _ Notifier = (*spyNotifier)(nil)`** in the test file catches a double
  that has drifted out of sync with the interface.
- **Concurrency**: a spy called from several goroutines needs a mutex, or
  `-race` will fail the test for you.
- **If a double is hard to write, the interface is too big.** That is feedback
  about the design, not the test.

## The exercises

- **mock1** — a spy that records notifications; assert both the firing and the
  silent case.
- **mock2** — a fake fetcher returning canned data and an error, covering both
  branches.

## Source references

- [pkg.go.dev: testing](https://pkg.go.dev/testing) ·
  [net/http/httptest](https://pkg.go.dev/net/http/httptest)
- [Go Code Review Comments: interfaces](https://go.dev/wiki/CodeReviewComments#interfaces)
- [Learn Go with Tests: Mocking](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking)
- [Go blog: Testable Examples](https://go.dev/blog/examples)

**End of the Intermediate block.** Next:
[stdlib_essentials](../stdlib_essentials/) — the packages the rest of the
curriculum builds on.

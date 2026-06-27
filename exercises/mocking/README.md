# Mocking (test doubles)

When code depends on an interface, a test can pass a hand-written **double** in
place of the real thing — no mocking library required:

- a **spy** records how it was called, so you can assert on the interaction
- a **stub/fake** returns canned data, so you can drive the logic (including
  error paths)

- **mock1** — a spy that records notifications
- **mock2** — a fake fetcher returning canned data and an error

## Resources

- [Learn Go with Tests — Mocking](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking)

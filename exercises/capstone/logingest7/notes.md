## logingest7 — SplitN, and fuzzing an invariant instead of an answer

```go
func ParseLine(line string, at time.Time) (Event, error) {
	parts := strings.SplitN(line, LineSeparator, 3)
	if len(parts) != 3 {
		return Event{}, fmt.Errorf("%w: want %d fields separated by %q, got %d",
			ErrInvalidEvent, 3, LineSeparator, len(parts))
	}

	e := Event{
		Source:  strings.TrimSpace(parts[0]),
		Level:   strings.TrimSpace(parts[1]),
		Message: strings.TrimSpace(parts[2]),
		At:      at,
	}
	if err := e.Validate(); err != nil {
		return Event{}, err
	}
	return e, nil
}
```

**Why it works**

- `SplitN` with `n = 3` does two jobs at once. It stops splitting after the
  second separator, so a message containing `|` survives intact — and it gives
  a bounded length to check before indexing.
- Delegating the field rules to `Event.Validate` means the text path and the
  JSON path from stage three cannot drift apart. One definition of "valid
  event", two ways in.
- Returning `Event{}` alongside a non-nil error, rather than a
  partially-filled event, removes the question of whether a caller who ignores
  the error gets something half-usable. They get the zero value, which
  `Validate` also rejects.

**Key detail:** `strings.Split` here would be a panic waiting for input.
`Split("api", "|")` returns one element, and `parts[2]` on it is
`index out of range`. In a log ingester, "input" means whatever an agent sends,
which eventually means anything at all. The length check is not defensive
programming — it is the parse.

**Key detail:** a fuzz target cannot assert on an expected value, because the
fuzzer invents the input and nobody knows the answer in advance. It asserts on
**invariants** — properties true for every input:

1. it never panics (implicit: a panic fails the target), and
2. when it reports success, the event really is valid.

That second one is what fails against the stub instantly. Invariants are also
what make fuzzing worth running on a parser at all: `go test -fuzz=FuzzParseLine`
explores millions of inputs a minute, and any input that breaks a property is
written to `testdata/fuzz/` and replayed as a normal unit test forever after.

**Key detail:** the benchmark uses `for b.Loop()` (Go 1.24) rather than
`for i := 0; i < b.N; i++`. `b.Loop` keeps the loop variable and the benchmarked
values alive across iterations, so the compiler cannot optimise the call away,
and it handles the timer around setup for you. The old form still works; the new
one is harder to write a misleading benchmark with.

**References**

- pkg.go.dev — strings.SplitN: https://pkg.go.dev/strings#SplitN
- Go — Fuzzing tutorial: https://go.dev/doc/tutorial/fuzz
- Go blog — Fuzzing is Beta Ready: https://go.dev/blog/fuzz-beta
- pkg.go.dev — testing.B.Loop: https://pkg.go.dev/testing#B.Loop

## iter4 — drive a sequence yourself with `iter.Pull`

```go
func zip(a, b iter.Seq[string]) []string {
	var out []string
	next, stop := iter.Pull(b)
	defer stop()
	for x := range a {
		y, ok := next()
		if !ok {
			break
		}
		out = append(out, x+y)
	}
	return out
}
```

**Why it works**

- `iter.Pull(seq)` returns `next func() (V, bool)` and `stop func()`. It runs the
  push iterator on a separate goroutine and hands you a **cursor**: each `next()`
  call resumes the sequence, takes one value, and pauses it again.
- That inverts control. A `range` loop is driven *by* the iterator; `next()` is
  driven *by you* — which is the only way to advance two sequences in lockstep.
- `ok` reports whether a value came out. When `b` is exhausted `next()` returns
  the zero value and `false`, so `break` ends the pairing at the shorter
  sequence — `["a","b","c"]` and `["1","2"]` give `["a1","b2"]`.
- The broken version nested `range b` inside `range a`. Ranging a push iterator
  always starts it **from the beginning**, so every outer pass paired with `"1"`
  and nothing ever detected that `b` had run out.

**Key detail:** `defer stop()` is not optional bookkeeping. If you abandon a
pulled sequence without calling `stop`, the goroutine parked inside the iterator
stays blocked forever — a leak. Calling `stop` more than once, or `next` after
the end, is safe by design, so the `defer` is always the right shape.

**References**

- iter package — Pull: https://pkg.go.dev/iter#Pull
- Go blog — Range Over Function Types: https://go.dev/blog/range-functions
- slices.Values: https://pkg.go.dev/slices#Values

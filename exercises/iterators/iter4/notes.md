## iter4 — drive a sequence with iter.Pull

```go
func zip(a, b iter.Seq[string]) []string {
    next, stop := iter.Pull(b)
    defer stop()

    var out []string
    for x := range a {
        y, ok := next()
        if !ok {
            break // b is exhausted
        }
        out = append(out, x+y)
    }
    return out
}
```

**Why it works**

- `iter.Pull` converts a push iterator (`iter.Seq`) into a `next`/`stop` pair you
  call yourself. That is what lockstep work needs: one value from `b` per element
  of `a`, with `ok == false` telling you when `b` ran out.

**Under the hood**

- A `range` loop drives an iterator start to finish — you cannot pause it. The
  broken version ranges `b` **inside** the loop over `a`, restarting it each
  pass, so every pair uses `b`'s first element and nothing notices the end.
  `Pull` runs the producer as a coroutine and resumes it one value at a time.

**Common mistake**

- Skipping `defer stop()`. The paused producer stays alive until it is stopped or
  collected, so leaving it out leaks the coroutine. After `stop`, `next` returns
  `false` forever.

**Key detail:** `iter.Pull2` does the same for `Seq2`. Reach for `Pull` when the
consumer needs control — zipping, merging sorted sequences, lookahead — and stay
with `range` everywhere else, because it is simpler and cannot leak.

**See also:** iter1 · iter3 · modern3 (the push side) · select1 (the channel
version of merging) · the [chapter](../README.md)

**References**

- pkg.go.dev — iter.Pull: https://pkg.go.dev/iter#Pull
- Go blog — Range Over Function Types: https://go.dev/blog/range-functions

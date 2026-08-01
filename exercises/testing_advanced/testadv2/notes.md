## testadv2 — fuzzing

```go
func FuzzReverse(f *testing.F) {
    for _, seed := range []string{"hello", "", "Hello, 世界"} {
        f.Add(seed) // seed corpus
    }
    f.Fuzz(func(t *testing.T, s string) {
        if reverse(reverse(s)) != s { t.Errorf(...) }        // property
        if utf8.ValidString(s) && !utf8.ValidString(reverse(s)) { ... }
    })
}
```

**Why it works**

- A fuzz test feeds **generated** inputs to find edge cases. You assert
  **properties** that must always hold (reversing twice returns the original;
  reversing valid UTF-8 stays valid) rather than exact outputs. `reverse` works on
  `[]rune` so multibyte characters survive.

**Key detail:** fuzzing shines for **round-trip / invariant** properties. Run it with
`go test -fuzz=FuzzReverse`; on a failure the offending input is saved to
`testdata/fuzz/` as a permanent regression test. The classic bug it catches:
reversing by **byte** instead of **rune** corrupts UTF-8.

**References**

- Go — Fuzzing tutorial: https://go.dev/doc/tutorial/fuzz

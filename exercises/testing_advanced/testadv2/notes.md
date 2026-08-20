## testadv2 — fuzzing finds the UTF-8 bug

```go
func reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

**Why it works**

- Reversing **bytes** splits multibyte UTF-8 characters apart, producing invalid
  strings — `"Hello, 世界"` becomes mojibake. Converting to `[]rune` first
  reverses whole code points, so the result round-trips and stays valid.

**Under the hood**

- A fuzz test states a **property** that must hold for every input, and the
  toolchain generates inputs trying to break it. The seed corpus (`f.Add`) runs
  during an ordinary `go test`; `go test -fuzz=FuzzReverse` starts generating new
  inputs. A failing input is written to `testdata/fuzz/`, where it becomes a
  permanent regression test.

**Common mistake**

- Writing a fuzz target that is not deterministic — randomness or time inside it
  makes a discovered failure unreproducible, which defeats the corpus.

**Key detail:** the two properties here are the classic pair: a **round trip**
(`reverse(reverse(s)) == s`) and an **invariant** (`utf8.ValidString`). Both are
checkable without knowing the answer, which is what makes property-based testing
work where table tests need you to write the expected output.

**See also:** strings2 (bytes vs runes) · testadv1 (table tests) ·
logingest7 (fuzzing a parser) · the [chapter](../README.md)

**References**

- Go — Fuzzing tutorial: https://go.dev/doc/tutorial/fuzz
- pkg.go.dev — testing.F: https://pkg.go.dev/testing#F

## testadv4 — benchmarks

```go
func BenchmarkSumSlice(b *testing.B) {
    data := make([]int, 1000)
    for i := range data { data[i] = i }
    for b.Loop() { // Go 1.24+: runs the body b.N times, timed
        sumSlice(data)
    }
}
```

**Why it works**

- A `Benchmark` function measures performance. `b.Loop()` (Go 1.24+) runs the body
  the right number of times to get a stable timing — setup **before** the loop is
  not counted. Run it with `go test -bench=.`.

**Nuance:** `for b.Loop()` replaces the old `for i := 0; i < b.N; i++` and, crucially,
keeps the compiler from optimizing the benchmarked call away (a common flaw in
hand-written benchmarks). Put all setup outside the loop so you time only the code
under test.

**References**

- pkg.go.dev — testing.B.Loop: https://pkg.go.dev/testing#B.Loop
- The Go Blog — Using subtests and sub-benchmarks: https://go.dev/blog/subtests

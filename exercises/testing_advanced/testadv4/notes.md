## testadv4 — benchmarks and b.Loop

```go
func sumSlice(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

**Why it works**

- The benchmark measures whatever the function does; the exercise is implementing
  it so the correctness test passes. A plain `go test` runs the `Test`, not the
  `Benchmark` — benchmarks need `go test -bench=.`.

**Under the hood**

- `for b.Loop()` (Go 1.24) replaced `for i := 0; i < b.N; i++` and fixes two
  long-standing problems: setup before the loop is **excluded from the timing
  automatically**, so `b.ResetTimer()` is unnecessary, and the compiler cannot
  optimise the call away, so the old "assign to a package-level sink" trick is
  too.

**Common mistake**

- Reading a single benchmark run as a measurement. Numbers move with CPU
  frequency, other load, and alignment — compare runs with `benchstat`, and use
  `-benchmem` to see B/op and allocs/op, which are usually more stable and more
  actionable than ns/op.

**Key detail:** benchmark first, then profile (`pprof1`) to find out **why**, then
change one thing and re-measure. Optimising without that loop is a change of
unknown sign.

**See also:** testadv1 · pprof1 (CPU profiles) · pprof2 (allocation) ·
slices1 (`make` with capacity) · the [chapter](../README.md)

**References**

- pkg.go.dev — testing.B.Loop: https://pkg.go.dev/testing#B.Loop
- benchstat: https://pkg.go.dev/golang.org/x/perf/cmd/benchstat

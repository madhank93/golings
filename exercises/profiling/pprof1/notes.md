## pprof1 — StopCPUProfile is the flush

```go
if err := pprof.StartCPUProfile(w); err != nil {
    return err
}
defer pprof.StopCPUProfile()
work()
```

**Why it works**

- `StartCPUProfile` begins sampling, but the samples are buffered — they are
  **written to `w` only when `StopCPUProfile` runs**. Without it the writer stays
  empty and nothing reports an error.

**Under the hood**

- The profiler interrupts the process about 100 times a second and records every
  goroutine's stack. Sampling means cheap (a few percent overhead) but
  statistical: a rare-yet-slow path can be invisible, and a very short workload
  produces too few samples to mean anything.

**Common mistake**

- Calling `Stop` at the end of the function body instead of deferring it. Any
  early return then leaves profiling **on**, and the next `StartCPUProfile`
  fails with "cpu profiling already in use".

**Key detail:** in practice you rarely write this by hand — `go test -cpuprofile
cpu.out -bench .` does it for you, then `go tool pprof cpu.out` opens the viewer.
Read `top` as flat (time in this function) vs cum (including callees): high cum,
low flat means the cost is downstream.

**See also:** pprof2 (memory) · pprof3 (live server) · testadv4 (benchmarks) ·
defer2 · the [chapter](../README.md)

**References**

- pkg.go.dev — runtime/pprof: https://pkg.go.dev/runtime/pprof
- Go blog — Profiling Go Programs: https://go.dev/blog/pprof

## pprof1 — capture a CPU profile

```go
if err := pprof.StartCPUProfile(w); err != nil {
    return err
}
defer pprof.StopCPUProfile()
work()
```

**Why it works**

- `pprof.StartCPUProfile(w)` begins sampling the CPU; the profile is only
  **flushed** to `w` when you call `StopCPUProfile`. `defer` guarantees the stop
  (and flush) runs, so `w` ends up non-empty.

**Nuance:** forget `StopCPUProfile` and the writer stays empty — the classic
mistake. Inspect the result with `go tool pprof <file>`. In practice you rarely
hand-roll this: `go test -cpuprofile cpu.out` and `import _ "net/http/pprof"`
(see pprof3) are the usual entry points.

**References**

- pkg.go.dev — runtime/pprof: https://pkg.go.dev/runtime/pprof
- The Go Blog — Profiling Go Programs: https://go.dev/blog/pprof

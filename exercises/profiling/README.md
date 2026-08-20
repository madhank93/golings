# Profiling

A benchmark tells you *that* something is slow. A profile tells you **where**.
Go ships the whole toolchain: sampling profilers in the runtime, a `go tool
pprof` viewer, and an HTTP endpoint that exposes both from a live process.

The rule that comes before any of it: **measure, do not guess**. The bottleneck
is almost never where it feels like it should be, and an optimisation applied
without a profile is a change of unknown sign.

## 1. CPU profiles

```go
func captureCPUProfile(w io.Writer, work func()) error {
    if err := pprof.StartCPUProfile(w); err != nil {
        return err
    }
    defer pprof.StopCPUProfile()      // the flush — not optional
    work()
    return nil
}
```

`pprof1`. The profiler interrupts the program ~100 times a second and records
every goroutine's stack. Those samples are buffered and only **written to `w`
when `StopCPUProfile` runs** — forget it and you get an empty file with no
error, which is the bug the exercise is built on.

`defer` it immediately after starting, for the same reason you defer every other
cleanup.

In practice you rarely write this by hand:

```sh
go test -cpuprofile cpu.out -bench .
go tool pprof cpu.out            # then: top, list <func>, web
```

Read `top` as **flat** (time in this function) versus **cum** (time including
what it called). A high cum with low flat means the cost is downstream.

## 2. Memory: `ReadMemStats` and the heap profile

```go
var before, after runtime.MemStats
runtime.GC()
runtime.ReadMemStats(&before)

keep := alloc()

runtime.GC()
runtime.ReadMemStats(&after)
runtime.KeepAlive(keep)

return after.HeapAlloc - before.HeapAlloc
```

`pprof2`. Two calls make this measurement meaningful and both are easy to skip:

- **`runtime.GC()` before each read**, so previously-freed garbage is not
  counted as live. Without it you measure allocation churn, not retention.
- **`runtime.KeepAlive(keep)`**, so the compiler cannot decide the value is dead
  before the second reading — which would report zero retention for an object
  you are still holding.

`ReadMemStats` **stops the world**, so it is a diagnostic, not something to call
per request. The useful fields: `HeapAlloc` (live now), `TotalAlloc` (cumulative,
never decreases), `Sys` (from the OS), `NumGC`, and `PauseTotalNs`.

For allocation *sites* rather than totals, use the heap profile —
`go test -memprofile mem.out`, then `go tool pprof -alloc_space` (everything
allocated) or `-inuse_space` (live at the sample).

## 3. `net/http/pprof` on a live server

```go
import _ "net/http/pprof"        // side-effect import: registers on DefaultServeMux

mux := http.NewServeMux()
mux.HandleFunc("/healthz", healthz)
mux.Handle("/debug/pprof/", http.DefaultServeMux)    // forward the prefix
```

`pprof3`. The import registers its handlers on `http.DefaultServeMux`. A server
using its **own** mux therefore exposes nothing — you have to forward the prefix
explicitly, which is what the exercise fixes.

That is also the security point: **the blank import alone can publish debug
endpoints**, so any package in your dependency tree importing it silently adds
`/debug/pprof/` to a default-mux server. Never expose it on a public listener.
Put it on a separate internal port, or behind auth.

Once wired, everything is a URL:

```sh
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30   # CPU
go tool pprof http://localhost:6060/debug/pprof/heap                 # heap
curl  http://localhost:6060/debug/pprof/goroutine?debug=2            # every stack
```

The goroutine profile is the one that finds **leaks**: a count that climbs and
never falls, with every stack parked on the same channel send.

## 4. The other profiles, and the workflow

| Profile | Answers |
|---|---|
| `profile` (CPU) | where is time spent |
| `heap` | what is allocated, and what is retained |
| `goroutine` | how many, and blocked where — leak detection |
| `block` | where goroutines wait on sync primitives (needs `SetBlockProfileRate`) |
| `mutex` | lock contention (needs `SetMutexProfileFraction`) |
| `allocs` | allocation sites since start |
| `trace` | scheduler, GC, and syscall timeline — `go tool trace` |

The loop that works: **benchmark → profile → change one thing → benchmark
again**, comparing with `benchstat`. And check the cheap wins first — an
allocation inside a loop, a `[]byte`→`string` conversion in a hot path, a
missing `make([]T, 0, n)` — before anything clever.

## Gotchas

- **No `StopCPUProfile`, no profile.** Silent empty output.
- **`ReadMemStats` stops the world**; do not call it in a request path.
- **Without `runtime.GC()` first**, heap deltas measure garbage, not retention.
- **Without `KeepAlive`**, the compiler may free the value before you measure it.
- **Importing `net/http/pprof` publishes endpoints** on the default mux —
  never on a public port.
- **CPU profiling has overhead** (a few percent) and samples: rare-but-slow
  paths can be invisible.
- **Profile a realistic workload.** A microbenchmark with a warm cache tells you
  about the cache.

## The exercises

- **pprof1** — defer `StopCPUProfile` so the profile is flushed.
- **pprof2** — return the heap delta around a forced GC.
- **pprof3** — forward `/debug/pprof/` from a custom mux to the default one.

## Source references

- [Go blog: Profiling Go Programs](https://go.dev/blog/pprof)
- [pkg.go.dev: runtime/pprof](https://pkg.go.dev/runtime/pprof) ·
  [net/http/pprof](https://pkg.go.dev/net/http/pprof) ·
  [runtime.MemStats](https://pkg.go.dev/runtime#MemStats)
- [Diagnostics guide](https://go.dev/doc/diagnostics) — profiling, tracing, and
  debugging in one place
- [go tool trace](https://pkg.go.dev/cmd/trace)

**Next: [applied](../applied/) →** — several chapters at once, in the shape of
real code.

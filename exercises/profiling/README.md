# Profiling

Profiling is how you find where a Go program spends time and memory. These
exercises cover capturing a CPU profile, reading heap statistics with
`runtime.ReadMemStats`, and exposing the `net/http/pprof` endpoints on a custom
mux — the everyday tooling for diagnosing production services.

- **pprof1** — `pprof.StartCPUProfile` / `StopCPUProfile` (flush on stop!)
- **pprof2** — `runtime.ReadMemStats` around `runtime.GC` to measure retention
- **pprof3** — wiring `net/http/pprof` onto a non-default `ServeMux`

## Resources

- [runtime/pprof package](https://pkg.go.dev/runtime/pprof)
- [net/http/pprof package](https://pkg.go.dev/net/http/pprof)
- [Go blog: Profiling Go Programs](https://go.dev/blog/pprof)

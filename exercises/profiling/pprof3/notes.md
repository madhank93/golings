## pprof3 — pprof endpoints on a custom mux

```go
import _ "net/http/pprof" // registers handlers on http.DefaultServeMux

mux := http.NewServeMux()
mux.HandleFunc("/healthz", healthz)
mux.Handle("/debug/pprof/", http.DefaultServeMux) // forward the prefix
```

**Why it works**

- The blank import registers `/debug/pprof/*` on `http.DefaultServeMux`. A server
  running its **own** mux therefore exposes none of it until the prefix is
  forwarded to the default mux.

**Under the hood**

- That is also the security story: a package anywhere in your dependency tree can
  blank-import `net/http/pprof`, and any server using the default mux then serves
  heap dumps and goroutine stacks. Owning your mux means the exposure is a
  deliberate line of code — this one.

**Common mistake**

- Registering `"/debug/pprof"` without the trailing slash. The pprof handlers
  live under the prefix (`/debug/pprof/heap`, `/goroutine`, …), and only a
  trailing-slash pattern matches a subtree.

**Key detail:** never expose this on a public listener — put it on a separate
internal port or behind auth. Once wired, everything is a URL:
`go tool pprof http://host/debug/pprof/profile?seconds=30` for CPU, `/heap` for
memory, and `/goroutine?debug=2` for every stack — the one that finds leaks.

**See also:** pprof1 · pprof2 · httpsrv2 (why your own mux) · logingest8 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — net/http/pprof: https://pkg.go.dev/net/http/pprof
- Go — Diagnostics: https://go.dev/doc/diagnostics

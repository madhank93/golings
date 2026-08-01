## pprof3 — expose net/http/pprof on a custom mux

```go
import _ "net/http/pprof" // side-effect: registers handlers on DefaultServeMux

mux := http.NewServeMux()
mux.Handle("/debug/pprof/", http.DefaultServeMux) // forward to the default mux
```

**Why it works**

- Importing `net/http/pprof` for its **side effect** (`_`) registers the
  `/debug/pprof/` handlers on `http.DefaultServeMux`. A server using its **own**
  mux won't expose them unless it forwards that prefix to the default mux.

**Key detail:** the blank import runs the package's `init()`, which does the
registration — you never call anything from it directly. If your app uses a
custom mux (most do), you must explicitly `Handle("/debug/pprof/", ...)` to wire
the profiler in. Never expose pprof on a public interface — it's debug-only.

**References**

- pkg.go.dev — net/http/pprof: https://pkg.go.dev/net/http/pprof

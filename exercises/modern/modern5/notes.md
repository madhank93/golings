## modern5 — per-iteration loop variables

The fix is to let the loop declare the variable:

```go
func labelers(items []string) []func() string {
	var fns []func() string
	for _, item := range items {
		fns = append(fns, func() string { return item })
	}
	return fns
}
```

**Why it works**

- Before Go 1.22 a `for ... range` loop created **one** variable and overwrote it
  each pass. Every closure captured that same variable, so after the loop they
  all read whatever it held last — `gamma`, three times.
- Go 1.22 changed the spec: a loop variable declared **by the loop** (`:=` in
  the `range` clause or the `for i := 0; ...` init) is a **new variable each
  iteration**. Each closure now captures its own, so the three functions report
  `alpha`, `beta`, `gamma`.
- The broken version dodged the fix by declaring `item` **outside** the loop and
  assigning to it with `=`. That variable belongs to the enclosing scope, not to
  the loop, so the old sharing behaviour is exactly what you get. The semantics
  follow the **declaration**, not the loop.

**Key detail:** this change is gated on the `go` line in your `go.mod`, not on
the toolchain. A module that still says `go 1.21` keeps the old per-loop
behaviour even when built with a current compiler — which is why the same source
can behave differently in two repositories.

**References**

- Go blog — Fixing For Loops in Go 1.22: https://go.dev/blog/loopvar-preview
- Go Wiki — LoopvarExperiment: https://go.dev/wiki/LoopvarExperiment
- Go 1.22 release notes — language: https://go.dev/doc/go1.22#language

## modern2 — range over an integer

```go
for i := range n { // i goes 0, 1, ..., n-1
    total += i
}
```

**Why it works**

- Go 1.22 lets `for range` take an **integer** directly: `for i := range n` runs
  `n` times with `i` from `0` to `n-1` — no `i := 0; i < n; i++` needed.

**Nuance:** the values are `0 .. n-1` (so `sumTo(5)` adds 0+1+2+3+4 = 10). Go 1.22
also fixed the classic loop-variable capture footgun — each iteration now gets a
**fresh** `i`, so `go func(){ use(i) }()` inside a loop is finally safe without
copying.

**References**

- Go 1.22 release notes — range over int: https://go.dev/doc/go1.22#language

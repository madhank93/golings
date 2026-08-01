## mapspkg2 — maps.DeleteFunc

```go
maps.DeleteFunc(m, func(k string, v int) bool {
    return v <= 0 // remove every entry whose TTL has expired
})
```

**Why it works**

- `maps.DeleteFunc` (Go 1.21+) removes every entry for which the predicate returns
  `true`, in a single pass — replacing the collect-keys-then-delete two-step.

**Key detail:** it's safe because the standard library handles the deletion correctly
internally. Doing it by hand, you can't both `range` a map and `delete` from it
across separate passes without care — `DeleteFunc` encapsulates that. Pairs with
`slices.DeleteFunc` for the slice equivalent.

**References**

- pkg.go.dev — maps.DeleteFunc: https://pkg.go.dev/maps#DeleteFunc

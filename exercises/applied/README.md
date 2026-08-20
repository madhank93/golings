# Applied

Two exercises that stop drilling one feature and instead combine several, the
way real code does. Implementing an interface the standard library defines, and
building a small concurrency-safe type out of a map, a mutex, methods, and a
sentinel error.

## 1. Implementing `sort.Interface`

```go
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

sort.Sort(ByAge(people))
```

`applied1`. Three methods on a **defined slice type**, and `sort.Sort` can order
it — the `methods` chapter (methods on non-struct types), the `interfaces`
chapter (implicit satisfaction), and the `type_aliases` chapter (defined types)
arriving at once.

The receivers are values, not pointers, and `Swap` still works: `a` is a slice,
so the copy shares the same backing array. That is the `pointers` chapter's
"copied header, shared data" rule doing something useful.

`Less` must be a **strict weak ordering** — `Less(i, i)` false, and consistent
across elements. A comparison like `a[i].Age <= a[j].Age` violates it and can
make the sort misbehave rather than merely order oddly.

Worth knowing what has changed around this API: modern Go usually writes

```go
slices.SortFunc(people, func(a, b Person) int { return cmp.Compare(a.Age, b.Age) })
```

which needs no type declaration and is faster (no interface dispatch per
comparison). `sort.Interface` still matters because you will read it constantly,
and because it is the clearest small example of an interface with more than one
method. Note the flipped convention: `sort.Interface` wants a `bool` "less",
`slices.SortFunc` wants a three-way `int`.

## 2. A concurrency-safe store

```go
var ErrNotFound = errors.New("key not found")

type Store struct {
    mu sync.Mutex
    m  map[string]int
}

func NewStore() *Store { return &Store{m: make(map[string]int)} }

func (s *Store) Set(key string, val int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[key] = val
}

func (s *Store) Get(key string) (int, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    v, ok := s.m[key]
    if !ok {
        return 0, ErrNotFound
    }
    return v, nil
}
```

`applied2`. Six ideas in twenty lines:

- **The mutex sits next to the data it guards**, first field, and both methods
  take it (`sync`).
- **Pointer receivers**, because a value receiver would copy the mutex — and a
  copied lock guards nothing (`methods`).
- **`NewStore` exists** because the zero value is not usable: a `nil` map panics
  on write (`maps`).
- **The map is unexported**, so the lock cannot be bypassed. This is the whole
  reason the type exists rather than passing a bare map around (`structs`).
- **Comma-ok distinguishes absent from zero** — without it, a stored `0` and a
  missing key are the same answer (`maps3`).
- **A sentinel error** lets callers write `errors.Is(err, ErrNotFound)` instead
  of comparing strings (`errors2`).

The design generalises: this is the shape of every cache, registry, and
in-memory index in Go. Scale it up and the next decisions are an `RWMutex` when
reads dominate (`safety1`), sharding when one lock becomes the bottleneck, and
`context` on the methods once anything can block.

```ascii
Store (exported type, unexported map)
   │
   ├─ mu   guards ──┐
   └─ m  map[K]V ◄──┘
        ▲
   Set / Get  — the only doors in; the lock cannot be skipped
```

## 3. What "applied" means in practice

Both exercises share a shape worth naming: **a small type that owns its
invariant**. `ByAge` owns its ordering; `Store` owns its locking. Callers cannot
get it wrong because there is no way in that bypasses the methods.

That is the payoff of the whole intermediate block — unexported fields, methods,
interfaces, and sentinel errors exist so that correctness lives in one place
instead of being every caller's responsibility.

## Gotchas

- **A value receiver on a type holding a mutex copies the lock.** Pointer
  receivers, always.
- **`Less` must be strict** (`<`, not `<=`).
- **`sort.Sort` is not stable**; `sort.SliceStable` / `slices.SortStableFunc`
  are.
- **A `nil` map panics on write** — hence the constructor.
- **Returning the internal map** from a getter leaks the thing the lock protects.
- **`defer mu.Unlock()` immediately after `Lock()`**, so every return path
  unlocks.
- **Run it with `-race`** — an unlocked path is invisible until it is not.

## The exercises

- **applied1** — implement `Less` and `Swap` so `sort.Sort` can order by age.
- **applied2** — implement `Set` on a mutex-guarded store whose `Get` returns a
  sentinel error.

## Source references

- [pkg.go.dev: sort.Interface](https://pkg.go.dev/sort#Interface) ·
  [slices.SortFunc](https://pkg.go.dev/slices#SortFunc) ·
  [cmp.Compare](https://pkg.go.dev/cmp#Compare)
- [pkg.go.dev: sync.Mutex](https://pkg.go.dev/sync#Mutex) ·
  [errors.Is](https://pkg.go.dev/errors#Is)
- [Effective Go: Interfaces and methods](https://go.dev/doc/effective_go#interfaces_and_types)
- [Data Race Detector](https://go.dev/doc/articles/race_detector)

**Next: [capstone](../capstone/) →** — the same integration, eight stages long,
as one running service.

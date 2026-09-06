## methods1 — value receiver vs pointer receiver

```go
func (r Rectangle) Area() float64 { return r.Width * r.Height }

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

**Why it works**

- `Area` only reads, so a **value receiver** — it works on a copy. `Scale` must
  change the caller's rectangle, so a **pointer receiver**, which holds the
  address.

**Under the hood**

- `r.Scale(2)` on a variable compiles to `(&r).Scale(2)`: the compiler takes the
  address for you when the value is **addressable**. It cannot do that for a map
  entry, a function's return value, or a literal — `Rectangle{3,4}.Scale(2)`
  does not compile.

**Common mistake**

- Assuming a value satisfies an interface whose method has a pointer receiver.
  It does not: pointer-receiver methods are in the method set of `*T` only, so
  `var s Scaler = Rectangle{}` fails with *"method Scale has pointer receiver"*.
  Pass `&Rectangle{}`.

**Key detail:** pick one receiver form per type and stay with it. Mixed receivers
make that method-set rule bite in code far from the declaration — and a value
receiver on a type holding a `sync.Mutex` copies the lock.

**See also:** methods2 (methods on non-structs) · pointers3 (copy vs address) ·
interfaces1 (what method sets are checked against) · the
[chapter](../README.md)

**References**

- Go spec — Method sets: https://go.dev/ref/spec#Method_sets
- Go Code Review Comments — receiver type: https://go.dev/wiki/CodeReviewComments#receiver-type

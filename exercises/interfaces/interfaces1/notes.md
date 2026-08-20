## interfaces1 — satisfaction is implicit

```go
func (r Rectangle) Area() float64 { return r.W * r.H }
func (c Circle) Area() float64    { return math.Pi * c.R * c.R }

// both now satisfy: type Shape interface { Area() float64 }
```

**Why it works**

- `Shape` lists one method. Any type with an `Area() float64` method satisfies it
  — no `implements` keyword, no declared relationship. `totalArea` then calls
  `Area` on each element without knowing which concrete type it holds.

**Under the hood**

- An interface value is two words: a pointer to type information (which carries
  the method table) and a pointer to the data. `s.Area()` looks the method up in
  that table and calls it with the data pointer — dynamic dispatch, one
  indirection.

**Common mistake**

- Defining the interface next to the implementations. Go's convention is the
  opposite: the **consumer** declares the interface it needs, and implementations
  return concrete types. "Accept interfaces, return structs" — it keeps
  interfaces small and lets types written before yours satisfy it.

**Key detail:** small is the point. `io.Reader` has one method and is the most
reused type in Go; an interface with six methods is hard to satisfy and harder
to fake in a test.

**See also:** interfaces2 (`Stringer`) · interfaces3 (getting the type back) ·
methods1 (method sets) · mock1 (why small interfaces are testable) ·
the [chapter](../README.md)

**References**

- Go spec — Interface types: https://go.dev/ref/spec#Interface_types
- Go Code Review Comments — interfaces: https://go.dev/wiki/CodeReviewComments#interfaces

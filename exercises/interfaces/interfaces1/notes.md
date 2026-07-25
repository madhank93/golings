## interfaces1 — implicit satisfaction

```go
type Shape interface{ Area() float64 }

func (r Rectangle) Area() float64 { return r.W * r.H }
func (c Circle) Area() float64    { return math.Pi * c.R * c.R }
```

**Why it works**

- Give `Rectangle` and `Circle` an `Area() float64` method and each **implicitly**
  satisfies `Shape` — no `implements` keyword. That lets `[]Shape` hold both and
  `totalArea` call `Area()` polymorphically.

**Nuance:** satisfaction is **structural** — if the methods match, the type fits
the interface, even one defined in another package. This is what makes Go
interfaces so composable: you can satisfy an interface you didn't know existed.

**References**

- A Tour of Go — Interfaces: https://go.dev/tour/methods/9
- Go by Example — Interfaces: https://gobyexample.com/interfaces

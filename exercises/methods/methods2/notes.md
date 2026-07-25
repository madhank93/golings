## methods2 — methods on any named type

```go
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
    return float64(c)*9/5 + 32
}
```

**Why it works**

- A receiver doesn't have to be a struct. Because `Celsius` is a **named type you
  declared**, you can hang methods on it — `Celsius(100).Fahrenheit()`.

**Nuance:** you may define methods only on types declared **in the same package**.
You can't add methods to `float64` or to types from other packages — that's why
`type Celsius float64` (a new named type) exists here. Note the explicit
`float64(c)` conversion: `Celsius` and `float64` are distinct types.

**References**

- Go by Example — Methods: https://gobyexample.com/methods
- A Tour of Go — Methods: https://go.dev/tour/methods/1

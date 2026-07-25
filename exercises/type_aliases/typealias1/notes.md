## typealias1 — a defined type needs explicit conversion

```go
type Celsius float64
func toFahrenheit(c float64) float64 { ... }

var boiling Celsius = 100
toFahrenheit(float64(boiling)) // explicit conversion required
```

**Why it works**

- `type Celsius float64` creates a **new, distinct** type. Even though its
  underlying type is `float64`, Go won't pass a `Celsius` where a `float64` is
  wanted without `float64(boiling)`.

**Nuance:** that friction is a **feature** — distinct types stop you from
accidentally adding a `Celsius` to a `Meter`, catching unit mix-ups at compile
time. Contrast with a type **alias** (typealias3), where no conversion is needed.

**References**

- Go spec — Type definitions: https://go.dev/ref/spec#Type_definitions

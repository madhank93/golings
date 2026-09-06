# Methods

A method is a function with a **receiver** — an extra parameter written before
the name, binding the function to a type. That is all the syntax there is, and
it buys three things: dot-call notation, interface satisfaction, and a natural
home for behaviour next to the data it operates on.

The one decision that matters is the receiver's form. `(r T)` receives a copy;
`(r *T)` receives an address. That choice is not just about mutation — it
decides which values can satisfy an interface, which is where it surprises
people.

## 1. Value receiver vs pointer receiver

```go
func (r Rectangle) Area() float64 {         // copy: read-only
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) { // address: mutates in place
    r.Width *= factor
    r.Height *= factor
}
```

`methods1` implements both on one type. `Area` cannot change the caller's
rectangle even if it tries — it holds a copy. `Scale` can, because it holds the
address.

Choose a pointer receiver when the method **mutates** the receiver, when the
struct is **large** enough that copying costs, or when it holds something that
must not be copied (a `sync.Mutex`, a `strings.Builder`). Choose a value
receiver for small immutable values — coordinates, durations, enums.

And then apply it **consistently**: if any method of a type needs a pointer
receiver, give them all pointer receivers. Mixed receivers make the method-set
rule below bite at a distance.

## 2. Method sets — the rule behind the surprise

Go inserts `&` and `*` for you when calling a method directly, so both of these
work on an addressable variable:

```go
r := Rectangle{3, 4}
r.Area()      // value receiver on a value
r.Scale(2)    // pointer receiver: compiler rewrites to (&r).Scale(2)
```

Interfaces get no such help. What satisfies an interface is the **method set**:

| Receiver | In the method set of `T` | In the method set of `*T` |
|---|---|---|
| `func (t T)` | yes | yes |
| `func (t *T)` | **no** | yes |

```go
type Scaler interface{ Scale(float64) }

var s Scaler = Rectangle{3, 4}    // does NOT compile
var s Scaler = &Rectangle{3, 4}   // fine
```

The error reads *"Rectangle does not implement Scaler (method Scale has pointer
receiver)"* — one of the messages worth recognising instantly. The reason is
addressability: an interface holds a copy of the value, and a copy inside an
interface has no address to take, so a pointer-receiver method has nothing to
mutate.

The same rule explains why `r.Scale(2)` fails on a value you cannot address,
such as a map entry (`m["k"].Scale(2)`) or a function's return value.

## 3. Methods on any named type

```go
type Celsius float64

func (c Celsius) Fahrenheit() float64 {
    return float64(c)*9/5 + 32
}

Celsius(100).Fahrenheit()   // 212
```

Receivers are not limited to structs — any type **defined in your package** can
have methods, including one whose underlying type is a number, a string, a
slice, or a func. `methods2` is the temperature case; `sort.Interface`
implementations on a named slice type (`applied1`) and `http.HandlerFunc` (a
method on a func type) are the same trick at work.

Note the conversion inside: `float64(c)` is required because `Celsius` is a
*distinct* type from `float64`. That distinction is the `type_aliases` chapter.

Two limits: you cannot define a method on a type from another package
(`func (t time.Time) Foo()` is rejected), and you cannot define one on a pointer
type or an interface type. Wrap the foreign type in your own defined type when
you need to attach behaviour.

## 4. Methods are functions, and can be used as values

```go
f := r.Area          // method value: r is captured now
f()                  // 12

g := Rectangle.Area  // method expression: receiver becomes the first parameter
g(r)                 // 12
```

A method value binds the receiver at the moment it is created — a common source
of "why did it use the old value?" when the receiver is later mutated. Method
expressions turn a method into an ordinary function, which is how a method can
be passed where a `func(T) U` is expected.

## Gotchas

- **Pointer-receiver methods are not in `T`'s method set**, so `T` does not
  satisfy an interface those methods provide. Pass `&t`.
- **Mixed receivers on one type** make the above unpredictable to a reader. Pick
  one.
- **A value receiver copies on every call.** For a large struct in a hot loop
  that is real cost.
- **Methods on `nil` pointers are legal** and only panic if the body
  dereferences — a `*Node` tree method handling `nil` is idiomatic.
- **You cannot add methods to types from other packages**; define your own type
  around it.
- **A `String()` method that formats the receiver with `%v` recurses forever**
  and blows the stack. Convert first: `float64(c)`.

## The exercises

- **methods1** — implement `Area` (value receiver, reads) and `Scale` (pointer
  receiver, mutates) on the same struct.
- **methods2** — attach a method to a named non-struct type, converting the
  receiver to do the arithmetic.

## Source references

- [Go spec: Method declarations](https://go.dev/ref/spec#Method_declarations) ·
  [Method sets](https://go.dev/ref/spec#Method_sets) ·
  [Method values](https://go.dev/ref/spec#Method_values)
- [Effective Go: Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go Code Review Comments: receiver type](https://go.dev/wiki/CodeReviewComments#receiver-type)
- [A Tour of Go: Methods](https://go.dev/tour/methods/1)

**Next: [interfaces](../interfaces/) →** — what those method sets are checked
against, and how Go does polymorphism without declaring it.

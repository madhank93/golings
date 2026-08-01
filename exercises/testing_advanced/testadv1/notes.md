## testadv1 — table-driven tests

```go
cases := []struct {
    name string
    in   int
    want string
}{
    {"multiple of three", 3, "Fizz"},
    // ...
}
for _, c := range cases {
    t.Run(c.name, func(t *testing.T) {
        if got := fizzbuzz(c.in); got != c.want { t.Errorf(...) }
    })
}
```

**Why it works**

- A slice of cases pairs each input with its expected output; the loop runs
  `t.Run(name, ...)` per case. Each `t.Run` is a **named subtest** that fails
  independently, so one bad case doesn't hide the others.

**Key detail:** this is the dominant Go testing idiom — add a case by adding a struct
literal, not a new function. Subtest names show in output (`TestFizzBuzz/multiple_of_five`)
and can be run selectively with `go test -run TestFizzBuzz/five`.

**References**

- The Go Blog — Using subtests and sub-benchmarks: https://go.dev/blog/subtests

## testadv1 — table-driven subtests

```go
func fizzbuzz(n int) string {
    switch {
    case n%15 == 0:
        return "FizzBuzz"
    case n%3 == 0:
        return "Fizz"
    case n%5 == 0:
        return "Buzz"
    default:
        return strconv.Itoa(n)
    }
}
```

**Why it works**

- The table pairs each input with its expected output and `t.Run` gives every
  case its own subtest. The `%15` case must come **first**: 15 is divisible by
  both 3 and 5, so a later check would never be reached.

**Under the hood**

- Each `t.Run` is a real subtest with its own name, so a failure identifies the
  case and `go test -run 'TestFizzBuzz/multiple_of_five'` runs exactly one.
  Spaces in the name become underscores in the path.

**Common mistake**

- Failing with `t.Errorf("wrong answer")`. A good message names **input, got, and
  want** — `fizzbuzz(5) = "Fizz", want "Buzz"` — so the message alone is enough
  to act on without opening the file.

**Key detail:** this is *the* Go testing idiom: adding a case is one line, and the
structure resists the temptation to test one thing per function. `t.Parallel()`
inside the subtest runs cases concurrently, and since Go 1.22 the old `tc := tc`
copy is no longer needed.

**See also:** testadv2 (fuzzing) · testadv3 (`httptest`) · switch2 ·
mock1 · the [chapter](../README.md)

**References**

- Go blog — Using subtests and sub-benchmarks: https://go.dev/blog/subtests
- Go Wiki — TableDrivenTests: https://go.dev/wiki/TableDrivenTests

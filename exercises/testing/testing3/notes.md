## testing3: testing the error path

```go
func TestGreetEmpty(t *testing.T) {
    _, err := greet("")
    if !errors.Is(err, ErrEmptyName) {
        t.Fatalf("greet(%q) error = %v, want ErrEmptyName", "", err)
    }
}
```

**Why it works**

- For an empty name the error is the result. The test fails when the error is
  missing (`nil`) or is a different error, and `errors.Is` covers both cases.

**Common mistake**

- `if err == nil`. That passes for any error at all, including one from a bug
  that has nothing to do with empty names.

**Key detail:** `TestGreet` uses `t.Fatalf` on the error because the string
check after it is meaningless if `greet` failed. Stop with `Fatalf` when later
checks depend on this one, and use `Errorf` otherwise.

**See also:** errors2 (`errors.Is`) · testing2 · mock1 · the [chapter](../README.md)

**References**

- pkg.go.dev: testing.T.Fatalf: https://pkg.go.dev/testing#T.Fatalf
- pkg.go.dev: errors.Is: https://pkg.go.dev/errors#Is

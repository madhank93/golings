# Loops

Go has exactly one loop keyword: `for`. It comes in three shapes, and between
them they cover everything `while`, `do`/`until` and C-style loops do in other
languages.

## 1. The three-clause loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i) // 0 1 2 3 4
}
```

Three parts, separated by semicolons:

| Part | Runs | Here |
|---|---|---|
| init | once, before the first iteration | `i := 0` |
| condition | before every iteration; the loop stops when it is false | `i < 5` |
| post | after every iteration | `i++` |

`i` exists only inside the loop. The condition is the part to read carefully:
`i < n` runs `n` times starting from 0, and `i <= n` includes `n` itself.
Getting it wrong by one is the most common loop bug there is.

## 2. The condition-only loop (Go's `while`)

```go
for n > 1 {
    n /= 2
}
```

Leave out init and post and you have a `while` loop. Something inside the body
must eventually make the condition false, or the loop never ends.

Leave out the condition too and you get a loop that runs forever, until a
`break` or `return` inside it:

```go
for {
    line, err := readLine()
    if err != nil {
        break
    }
    handle(line)
}
```

## 3. `break` and `continue`

```go
for i := 0; i < len(nums); i++ {
    if nums[i] == 0 {
        continue // skip the rest of this iteration, go to the next one
    }
    if nums[i] < 0 {
        break // leave the loop entirely
    }
    total += nums[i]
}
```

Both act on the innermost loop. To leave an outer loop from inside an inner one,
label the outer loop (`outer: for ...`) and write `break outer`.

## Gotchas

- **Off by one.** Check the first and last value of the loop variable, not just
  that the loop "looks right".
- **Infinite loops.** In a condition-only loop, make sure the body changes what
  the condition tests.
- **Indexing past the end.** `for i := 0; i <= len(s); i++` reads `s[len(s)]`
  on its last iteration and panics. Use `<`.

## The exercises

- **loops1**: fix a three-clause loop that stops one number early.
- **loops2**: write a condition-only loop.
- **loops3**: skip and stop with `continue` and `break`.

## Source references

- [A Tour of Go: For](https://go.dev/tour/flowcontrol/1) and the pages after it
- [Effective Go: For](https://go.dev/doc/effective_go#for)
- [The Go spec: For statements](https://go.dev/ref/spec#For_statements) ·
  [Break](https://go.dev/ref/spec#Break_statements) ·
  [Continue](https://go.dev/ref/spec#Continue_statements)

**Next: [maps](../maps/) →**, the other built-in collection. After it,
[range](../range/) gives a shorter loop over slices and maps.

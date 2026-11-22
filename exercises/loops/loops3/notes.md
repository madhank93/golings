## loops3: `continue` and `break`

```go
for i := 0; i < len(nums); i++ {
    if nums[i] == 0 {
        continue
    }
    if nums[i] < 0 {
        break
    }
    total += nums[i]
}
```

**Why it works**

- `continue` skips the rest of this iteration; the loop's post statement
  (`i++`) still runs. `break` leaves the loop, so nothing after the first
  negative number is added.

**Common mistake**

- Putting the checks after `total += nums[i]`. The negative number is added
  before the loop stops.

**Key detail:** zeros add nothing anyway, so `continue` here only makes the
intent explicit. In real code it usually skips something expensive, like a
blank line before parsing it.

**See also:** loops2 · range3 · the [chapter](../README.md)

**References**

- Go spec, Break statements: https://go.dev/ref/spec#Break_statements
- Go spec, Continue statements: https://go.dev/ref/spec#Continue_statements

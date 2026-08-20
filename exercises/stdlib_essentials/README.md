# Standard library essentials

Go's standard library is the reason so much production Go has a short
`go.mod`. JSON, HTTP, time, regular expressions, and file I/O all ship with the
toolchain, are covered by the compatibility promise, and are usually good
enough that reaching for a dependency needs a reason.

These seven exercises are a tour of the packages you will use in almost every
program: `encoding/json`, `io`, `slices`, `time`, `strconv`, `regexp`, and one
recent JSON addition worth knowing about.

## 1. `encoding/json`: tags are the schema

```go
type User struct {
    Name string `json:"full_name"`
    Age  int    `json:"user_age"`
}

json.Unmarshal(data, &u)     // note the pointer
out, err := json.Marshal(u)
```

Marshaling walks the struct with reflection, so two rules follow immediately:

- **Only exported fields are visible.** A lowercase field is invisible to
  `encoding/json` — no error, no key, just missing data. This is the single most
  common JSON bug in Go.
- **Tags rename**, and without one the field name is used verbatim. Matching on
  decode is case-**insensitive**, so `Name` already accepts `"name"` — `stdlib1`
  needs tags because its keys differ by more than case.

Useful tag options: `json:"-"` (never encode), `json:"name,omitempty"` (drop
when empty), and `json:",string"` (encode a number as a JSON string).

Decoding into `any` gives you `map[string]any` with every number as a
`float64` — fine for exploration, painful in real code. Decode into a struct.

## 2. `io.Reader` / `io.Writer`: the two interfaces everything speaks

```go
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }

n, err := io.Copy(dst, src)     // stream from any Reader to any Writer
```

One method each, and between them they cover files, network connections,
buffers, compressors, hashers, and HTTP bodies. `stdlib2` is `io.Copy` — the
call that streams without ever holding the whole input in memory.

The vocabulary worth memorising:

| Call | Does |
|---|---|
| `io.Copy(dst, src)` | stream everything, fixed-size buffer |
| `io.ReadAll(r)` | read into memory (only when you know it is small) |
| `strings.NewReader(s)` / `bytes.NewReader(b)` | turn data into a Reader |
| `io.LimitReader(r, n)` | cap how much can be read — use on untrusted input |
| `io.TeeReader(r, w)` | read and copy at the same time |
| `io.MultiWriter(a, b)` | write to several destinations |

`Read` returning `io.EOF` is not a failure; it is how a stream ends.

## 3. `slices` and the generic helpers

```go
out := slices.Clone(nums)
slices.Sort(out)                                   // ascending
slices.SortFunc(out, func(a, b int) int { return b - a })   // custom order
slices.Reverse(out)
slices.Contains(out, 3)
i, found := slices.BinarySearch(out, 3)
```

`stdlib3` sorts descending, which has two idiomatic answers: sort then
`slices.Reverse`, or `slices.SortFunc` with a comparison that returns negative,
zero, or positive — the same contract as `cmp.Compare` (and the opposite
convention from the old `sort.Slice`, which took a `less bool`).

`slices.Sort` is not stable; `slices.SortStableFunc` is. And `slices.Clone`
matters more than it looks — see the `maps_package` chapter for the same trap
with maps.

## 4. `time`: the reference layout

Go formats dates by **example**, not by `%Y-%m-%d` codes:

```
Mon Jan  2 15:04:05 MST 2006
 |    |   |  |  |  |   |  |
 |    |   |  |  |  |   |  +-- 2006  year
 |    |   |  |  |  |   +----- MST   timezone
 |    |   |  |  |  +--------- 05    second
 |    |   |  |  +------------ 04    minute
 |    |   |  +--------------- 15    hour (24h; 03 for 12h)
 |    |   +------------------ 2     day
 |    +---------------------- Jan   month
 +---------------------------- Mon  weekday
```

That reference instant is `01/02 03:04:05PM '06 -0700` — the digits 1 through 7
in order, which is the mnemonic. So `"2006-01-02"` means YYYY-MM-DD
(`stdlib4`), and `"15:04:05"` is a 24-hour clock.

Also worth knowing: `time.Duration` is a named `int64` of nanoseconds, so
`5*time.Second` is arithmetic, not a constructor; compare times with `Before`,
`After`, `Equal` — never `==`, which compares the monotonic clock reading and
location too.

## 5. `strconv`: strings ↔ numbers

```go
n, err := strconv.Atoi("21")            // string -> int
s := strconv.Itoa(21)                   // int -> string
f, err := strconv.ParseFloat("3.14", 64)
b, err := strconv.ParseBool("true")
q := strconv.Quote(s)                   // "with \"escapes\""
```

`stdlib5`. The error is the point: `Atoi("nope")` returns a `*strconv.NumError`,
which is how you tell "the user typed nonsense" from "the user typed 0".

Do **not** use `fmt.Sprintf("%d", n)` for this — `strconv.Itoa` is several times
faster and says what it means. And `string(65)` is not a conversion, it is the
rune `"A"`; `go vet` flags it.

## 6. `regexp`: compile once

```go
var emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

func isEmail(s string) bool { return emailRe.MatchString(s) }
```

`stdlib6`. Compiling is the expensive part, so compile at package level with
`MustCompile` (which panics on a bad pattern — correct at init time, since a
literal pattern that does not compile is a bug you want to hear about
immediately), and call `MatchString` in the hot path.

Go's `regexp` uses RE2: **linear time, no backtracking**, and therefore immune to
catastrophic-backtracking denial of service. The trade is no backreferences and
no lookahead. Note that `MatchString` searches anywhere in the string unless you
anchor with `^…$`.

For fixed substrings, `strings.Contains`/`HasPrefix` beat a regexp by a wide
margin.

## 7. `omitzero` vs `omitempty` (Go 1.24)

```go
type Event struct {
    Name string    `json:"name"`
    At   time.Time `json:"at,omitzero"`   // was omitempty
}
```

`omitempty` predates most of Go's type system. It drops empty strings, zero
numbers, `false`, and nil/empty collections — but **a struct is never "empty"**
to it, so a zero `time.Time` marshals as `"0001-01-01T00:00:00Z"`. `stdlib7` is
that surprise.

`omitzero` (Go 1.24) asks the type whether it is the zero value (using
`IsZero() bool` when present), so a zero `time.Time` disappears as intended. Use
`omitzero` for structs and anything with a meaningful zero; `omitempty` remains
right for "drop the empty slice".

## Gotchas

- **Unexported fields are invisible to `encoding/json`.**
- **`json.Unmarshal` needs a pointer** — passing a value fails at run time, not
  compile time.
- **`io.ReadAll` on an untrusted body is a memory bomb.** Wrap with
  `io.LimitReader` or use `http.MaxBytesReader`.
- **`io.EOF` is a normal end**, not an error to report.
- **The time layout is the reference date**, not a format string —
  `"YYYY-MM-DD"` silently produces garbage rather than failing.
- **Never compare `time.Time` with `==`**; use `Equal`.
- **`regexp.MustCompile` inside a function** recompiles on every call.
- **`omitempty` cannot drop a zero struct** — that is what `omitzero` is for.

## The exercises

- **stdlib1** — tag the struct so JSON keys map onto the fields.
- **stdlib2** — stream a reader into a buffer with `io.Copy`.
- **stdlib3** — sort descending with the `slices` helpers.
- **stdlib4** — parse a date with the reference layout.
- **stdlib5** — convert a string to a number and handle the failure.
- **stdlib6** — compile a pattern once and match against it.
- **stdlib7** — swap `omitempty` for `omitzero` so a zero time disappears.

## Source references

- [pkg.go.dev: encoding/json](https://pkg.go.dev/encoding/json) ·
  [io](https://pkg.go.dev/io) · [slices](https://pkg.go.dev/slices) ·
  [time](https://pkg.go.dev/time) · [strconv](https://pkg.go.dev/strconv) ·
  [regexp](https://pkg.go.dev/regexp)
- [Go blog: JSON and Go](https://go.dev/blog/json)
- [Go blog: Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
  — why RE2 has no backtracking
- [Go 1.24 release notes: omitzero](https://go.dev/doc/go1.24#encoding-json)

**Next: [maps_package](../maps_package/) →** — the same generic-helper treatment
for maps, and the copy that is not a copy.

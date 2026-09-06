## stdlib6 — regexp, compiled once

```go
var emailRe = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

func isEmail(s string) bool { return emailRe.MatchString(s) }
```

**Why it works**

- Compiling turns the pattern into a matcher; `MatchString` then runs it. The
  anchors `^` and `$` are what make `"a@b"` fail — without them the pattern would
  match anywhere inside the string.

**Under the hood**

- Go's `regexp` implements **RE2**: matching is linear in the input length and
  never backtracks, so a hostile pattern or input cannot cause the exponential
  blowup that powers ReDoS attacks in other languages. The trade is no
  backreferences and no lookahead.

**Common mistake**

- Calling `regexp.MustCompile` inside the function. Compilation is the expensive
  part; putting it at package level does it once at init. `MustCompile` panics on
  a bad pattern, which is correct for a literal — a pattern that does not compile
  is a bug you want at startup, not at request time.

**Key detail:** use backticks for patterns so `\s` and `\.` do not need double
escaping. And for fixed substrings, `strings.Contains`/`HasPrefix` are far faster
than any regexp — real email validation, meanwhile, is `net/mail.ParseAddress`.

**See also:** strings1 (the cheaper tools) · stdlib5 (parsing input) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — regexp: https://pkg.go.dev/regexp ·
  syntax: https://pkg.go.dev/regexp/syntax
- RE2: why Go's regexp has no backtracking: https://swtch.com/~rsc/regexp/regexp1.html

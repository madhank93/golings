## pointers1 — write through a pointer

```go
func double(p *int) {
    *p = *p * 2
}

n := 21
double(&n) // n is now 42
```

**Why it works**

- `&n` passes `n`'s **address**. Inside, `*p` on the right reads the value at
  that address and `*p` on the left writes to it — so the caller's variable
  changes.

**Under the hood**

- Go passes everything by value, pointers included: `double` gets a copy of the
  *address*, which is enough, because both copies name the same memory. That is
  the whole mechanism — there is no pass-by-reference in the language.

**Common mistake**

- Assigning to `p` instead of `*p`. `p = &other` rebinds the local copy of the
  pointer and the caller sees nothing. The dereference is what reaches through.

**Key detail:** the zero value of a pointer is `nil`, and dereferencing it panics
with `invalid memory address or nil pointer dereference`. There is no pointer
arithmetic — `p+1` does not compile, which removes a whole class of memory bugs.

**See also:** pointers2 (structs) · pointers3 (value vs pointer parameters) ·
unsafe1 (where arithmetic does exist) · the [chapter](../README.md)

**References**

- Go spec — Address operators: https://go.dev/ref/spec#Address_operators
- Go FAQ — pass by value: https://go.dev/doc/faq#pass_by_value

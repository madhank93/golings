# The unsafe package

`unsafe` is Go's escape hatch from the type system. It has no run-time
representation and generates no code — it is a compile-time instruction to stop
checking. With it you can read a struct's memory layout, reinterpret one type as
another, and convert between `[]byte` and `string` without copying.

The name is the documentation. Code using `unsafe` is exempt from the Go 1
compatibility promise, invisible to the race detector's usual guarantees, and
capable of corrupting memory in ways the rest of Go cannot. These two exercises
stick to the two well-established, spec-blessed idioms.

## 1. Layout introspection: `Sizeof`, `Offsetof`, `Alignof`

```go
type Record struct {
    A byte
    B int64
    C int32
}

unsafe.Sizeof(Record{})     // total size including padding
unsafe.Offsetof(r.C)        // byte offset of C within Record
unsafe.Alignof(r.B)         // alignment requirement of B
```

`unsafe1`. All three are **constant expressions** evaluated at compile time —
zero run-time cost, and unlike `reflect` they need no value at all. They are the
tool for answering "why is this struct 24 bytes when its fields add up to 13?":

```ascii
type Record struct {        offset  bytes
    A byte                    0     [A][pad × 7]
    B int64                   8     [........ B ........]
    C int32                  16     [.. C ..][pad × 4]
}                                   = 24 bytes total

reordered (B, C, A):        offset
    B int64                   0     [........ B ........]
    C int32                   8     [.. C ..][A][pad × 3]
    A byte                   12
                                    = 16 bytes total
```

Same fields, 33% less memory, because the compiler **does not reorder fields**.
This only matters when you have millions of instances — and when it does,
`unsafe.Sizeof` plus `go vet`'s `fieldalignment` check are how you see it.

## 2. Zero-copy `[]byte` ⇄ `string`

```go
func bytesToString(b []byte) string {
    return unsafe.String(unsafe.SliceData(b), len(b))
}

func stringToBytes(s string) []byte {
    return unsafe.Slice(unsafe.StringData(s), len(s))
}
```

`unsafe2`. The ordinary conversions `string(b)` and `[]byte(s)` **allocate and
copy**, because a string is immutable and a slice is not — sharing memory would
let one break the other's contract. These builders skip the copy by handing the
string the slice's backing array.

The obligation transfers to you: **the bytes must never be mutated while the
string exists.** Do that and you get a string whose contents change, which
violates an invariant the runtime and every map implementation rely on.

`unsafe.String`, `StringData`, `Slice`, and `SliceData` (Go 1.17/1.20) replaced
the old `*reflect.StringHeader` cast, which was never actually safe — the header
types were not guaranteed to match the runtime's layout. Use the builders.

Where this is worth it: decoding hot paths where a `[]byte` from a buffer is used
as a map key or compared, and the bytes are provably read-only afterwards.
Everywhere else, `string(b)` is one allocation and infinitely more explainable.

## 3. The rules that still apply

`unsafe.Pointer` is the universal pointer type, and the spec lists exactly six
legal conversion patterns. The two that matter here:

- **`*T1` → `unsafe.Pointer` → `*T2`** is legal when the layouts are compatible.
- **Pointer arithmetic must be done in one expression**:
  `unsafe.Pointer(uintptr(p) + offset)`. Storing the `uintptr` in a variable
  first is a bug, because a `uintptr` is *not* a reference — the garbage
  collector can move or free the object between the two statements.

Run anything using `unsafe` under `go vet` (which has an `unsafeptr` check) and
`go test -race`, and consider `GOEXPERIMENT=cgocheck2` for deeper validation.

## 4. When not to

Almost always. The honest checklist before using `unsafe`:

1. Have you measured, and is this actually the bottleneck?
2. Does a safe API exist? (`strings.Builder`, `slices`, generics, `binary`)
3. Can the assumption it relies on be broken by a future Go release?
4. Is it isolated behind a small, tested, documented function?

The standard library itself uses `unsafe` in exactly these conditions — inside
`strings`, `sync`, and `reflect`, wrapped in safe APIs you use without knowing.
That is the model: contained, justified, and never spread through a codebase.

## Gotchas

- **`uintptr` is not a pointer.** The GC does not track it; a stored `uintptr`
  can outlive the object it named.
- **Mutating bytes behind a zero-copy string** breaks string immutability with
  no error and no crash — until something hashes it.
- **`unsafe.Sizeof` is compile-time and excludes referenced data** — the size of
  a slice header, not of its elements.
- **Struct layout is not part of the language spec.** A future compiler may pad
  differently.
- **`unsafe` disables the compatibility promise** for that code.
- **The old `reflect.StringHeader`/`SliceHeader` casts are deprecated** — use
  `unsafe.String`/`Slice`.

## The exercises

- **unsafe1** — report a field's byte offset with `unsafe.Offsetof`.
- **unsafe2** — build a string that shares a `[]byte`'s backing array.

## Source references

- [Go spec: Package unsafe](https://go.dev/ref/spec#Package_unsafe) — the six
  legal `unsafe.Pointer` patterns
- [pkg.go.dev: unsafe](https://pkg.go.dev/unsafe) — `String`, `StringData`,
  `Slice`, `SliceData`
- [Go 1.20 release notes](https://go.dev/doc/go1.20#unsafe) ·
  [`go vet` unsafeptr](https://pkg.go.dev/cmd/vet)

**Next: [files](../files/) →** — back to safe ground, and the everyday I/O the
`io` interfaces were built for.

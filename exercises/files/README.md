# Files

File I/O in Go is two decisions: read the whole thing into memory, or stream it.
The standard library makes both a one-liner, and picking wrong is how a service
that worked on a 2 KB fixture falls over on a 2 GB log.

Everything here sits on the `io.Reader`/`io.Writer` interfaces from the
`stdlib_essentials` chapter, which is why the same code works against a file, a
network connection, or a `strings.Reader` in a test.

## 1. Whole-file: `os.ReadFile` / `os.WriteFile`

```go
err := os.WriteFile(path, []byte(content), 0o644)
data, err := os.ReadFile(path)
```

`files1`. No `Open`, no `defer Close`, no error-prone loop — the file is opened,
handled, and closed inside the call. This is the right choice for config files,
fixtures, and anything with a known small size.

The `0o644` is a Unix permission mode, applied only when the file is created:
owner read/write, everyone else read. `0o600` for secrets. On Windows only the
read-only bit is honoured.

Note the API generation: these are `os.ReadFile`/`WriteFile` (Go 1.16+), not the
older `ioutil.ReadFile` — the `io/ioutil` package is deprecated and its contents
moved into `os` and `io`.

## 2. Streaming: `bufio.Scanner`

```go
scanner := bufio.NewScanner(r)
for scanner.Scan() {
    line := scanner.Text()      // the line, without its newline
    …
}
if err := scanner.Err(); err != nil {   // NEVER skip this
    return err
}
```

`files2`. `Scan()` advances and returns false at EOF **or on error** — which is
why `scanner.Err()` is not optional: without it, a read failure silently looks
like a short file.

Memory stays flat regardless of input size, because only one line is held at a
time. That property is what makes this the default for logs, CSVs, and stdin.

Two knobs worth knowing:

- **`scanner.Buffer(buf, max)`** — the default maximum token is 64 KB, and a
  longer line stops the scan with `bufio.ErrTooLong`. Raise it deliberately for
  formats with long lines (JSONL, say).
- **`scanner.Split(bufio.ScanWords)`** — or `ScanRunes`, or a custom
  `SplitFunc`. Lines are just the default.

`scanner.Text()` allocates a string per line; `scanner.Bytes()` returns the
internal buffer, which is only valid until the next `Scan`.

## 3. The full form, and why `defer f.Close()` is not enough

```go
f, err := os.Open(path)      // read-only; os.Create truncates, OpenFile is explicit
if err != nil {
    return err
}
defer f.Close()
```

Open, **check the error, then** defer the close — deferring first would call
`Close` on a nil file. This is the `defer` chapter's rule with a real
consequence.

For **writes**, a deferred `Close` is not sufficient: `Close` can fail (a full
disk, a flush error), and a deferred call discards its error. When the write must
be durable, capture it:

```go
func save(path string, data []byte) (err error) {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer func() { err = errors.Join(err, f.Close()) }()
    _, err = f.Write(data)
    return err
}
```

Buffered writers need one more step: `bufio.Writer` holds data in memory until
`Flush()`, so an unflushed buffer means a silently truncated file.

## 4. Paths, and testing file code

```go
filepath.Join(dir, "note.txt")     // OS-correct separators — never "dir" + "/" + name
filepath.Ext(name)                  // ".txt"
os.MkdirAll(dir, 0o755)
```

Use `path/filepath` (OS-aware) rather than `path` (always forward slashes; for
URLs and embedded FS).

Tests get `t.TempDir()`, which creates a directory and removes it — and
everything in it — when the test finishes, with no cleanup code and no cross-test
interference. That is what `files1`'s test uses.

For code that only *reads*, taking an `fs.FS` instead of a path makes it
testable with `fstest.MapFS` and usable with `embed`:

```go
func Load(fsys fs.FS, name string) ([]byte, error) { return fs.ReadFile(fsys, name) }
```

## Gotchas

- **`os.ReadFile` on an unbounded file is a memory bomb.** Stream it.
- **`scanner.Err()` is mandatory** — otherwise errors look like EOF.
- **The default scanner token cap is 64 KB.**
- **`scanner.Bytes()` is only valid until the next `Scan`.** Copy to keep it.
- **A deferred `Close` on a writer swallows its error**, and an unflushed
  `bufio.Writer` loses data.
- **Defer the close *after* the error check.**
- **Check-then-open is a race** — `os.Stat` followed by `os.Open` can be wrong by
  the time you open. Just open and handle the error, with `errors.Is(err,
  fs.ErrNotExist)`.
- **Build paths with `filepath.Join`**, not string concatenation.

## The exercises

- **files1** — write a file and read it back with the whole-file calls.
- **files2** — count lines with `bufio.Scanner`, one line at a time.

## Source references

- [pkg.go.dev: os](https://pkg.go.dev/os) ·
  [bufio.Scanner](https://pkg.go.dev/bufio#Scanner) ·
  [path/filepath](https://pkg.go.dev/path/filepath) ·
  [io/fs](https://pkg.go.dev/io/fs)
- [Go 1.16 release notes: io/ioutil deprecation](https://go.dev/doc/go1.16#ioutil)
- [pkg.go.dev: testing.T.TempDir](https://pkg.go.dev/testing#T.TempDir) ·
  [testing/fstest](https://pkg.go.dev/testing/fstest)

**Next: [http_client](../http_client/) →** — the same reader and closer
discipline, over a network.

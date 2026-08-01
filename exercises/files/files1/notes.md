## files1 — whole-file read and write

```go
os.WriteFile(path, []byte(content), 0o644)
data, err := os.ReadFile(path)
```

**Why it works**

- `os.WriteFile` writes a whole `[]byte` to a file in one call (creating or
  truncating it); `os.ReadFile` slurps the whole file back. The `0o644` is the
  Unix permission mode (owner read/write, others read).

**Key detail:** these are the simple, no-`defer`-`Close` path for **small** files —
both open, do the I/O, and close internally. For large files or streaming, open
with `os.Open`/`os.Create` and use a reader/writer instead. The test uses
`t.TempDir()`, which gives a fresh directory that's auto-cleaned.

**References**

- Go by Example — Writing Files: https://gobyexample.com/writing-files
- pkg.go.dev — os: https://pkg.go.dev/os

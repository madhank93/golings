## files1 — whole-file read and write

```go
func saveAndLoad(path, content string) (string, error) {
    if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
        return "", err
    }
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

**Why it works**

- `os.WriteFile` and `os.ReadFile` open, do the I/O, and close in one call —
  no `defer f.Close()` to forget, and no partial-write loop.

**Under the hood**

- `0o644` is a Unix permission mode applied **only when the file is created**:
  owner read/write, others read. Use `0o600` for anything sensitive. On Windows
  only the read-only bit is honoured.

**Common mistake**

- Using these for files of unknown size. `os.ReadFile` loads the whole thing into
  memory, so a 2 GB log is a 2 GB allocation. Stream those with `bufio.Scanner`
  (files2) or `io.Copy`.

**Key detail:** these are the `os` versions, not `ioutil` — `io/ioutil` was
deprecated in Go 1.16 and its contents moved into `os` and `io`. The test uses
`t.TempDir()`, which creates a directory and removes it with everything in it
when the test ends.

**See also:** files2 (streaming) · stdlib2 (`io.Copy`) · defer2 (the
`Open`/`Close` form) · the [chapter](../README.md)

**References**

- pkg.go.dev — os.WriteFile: https://pkg.go.dev/os#WriteFile ·
  os.ReadFile: https://pkg.go.dev/os#ReadFile
- pkg.go.dev — testing.T.TempDir: https://pkg.go.dev/testing#T.TempDir

## binary4 — random access with io.ReaderAt

```go
offset := int64(n-1) * int64(pageSize)
size := pageSize
if n == 1 {
    offset, size = fileHeaderSize, pageSize-fileHeaderSize
}
page := make([]byte, size)
_, err := io.ReadFull(io.NewSectionReader(ra, offset, int64(size)), page)
```

**Why it works**

- `io.ReaderAt` reads at an **absolute offset** and keeps no cursor. That is the
  whole difference from `io.Reader`: no seek, no shared position, so several
  goroutines can read different pages of the same file at once with no lock.
  `*os.File` and `*bytes.Reader` both implement it.

**Under the hood**

- Paged layouts are how on-disk B-trees work: fixed-size pages let the engine
  compute an address instead of scanning, and let the OS page cache do the
  caching. SQLite's page 1 really does carry a 100-byte file header before its
  content, and pages really are numbered from 1 — the off-by-one is in the
  format, not in your code.

**Common mistake**

- Trusting the returned byte count. `ReadAt` may return fewer bytes than
  requested along with a nil error; wrapping it in `io.ReadFull` (directly, or
  via `io.NewSectionReader`) turns a short page into an error instead of a
  half-decoded record.

**Key detail:** `io.NewSectionReader(ra, off, n)` gives you a reader bounded to
one page — an `io.Reader` and an `io.ReaderAt` — so a page decoder written
against `io.Reader` physically cannot read past its page.

**See also:** binary1 (decoding what a page holds) · files1 (`os.ReadFile`) ·
the [chapter](../README.md)

**References**

- pkg.go.dev — io.ReaderAt: https://pkg.go.dev/io#ReaderAt
- SQLite file format §1.2 (pages, 1-based): https://www.sqlite.org/fileformat.html

## httpsrv3 — JSON in, JSON out

```go
func echoHandler(w http.ResponseWriter, r *http.Request) {
    var req echoReq
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(echoResp{Greeting: "Hello, " + req.Name + "!"}); err != nil {
        return
    }
}
```

**Why it works**

- `r.Body` is an `io.Reader` and `w` is an `io.Writer`, so the decoder and
  encoder stream straight through them — no intermediate buffer, and no
  `io.ReadAll`.

**Under the hood**

- `Encode` writes immediately, which is why `Content-Type` must be set **before**
  it. Once the first byte is written the headers are sent and any later change is
  ignored — and the status is already 200, so an encoding failure mid-response
  cannot be turned into a 500.

**Common mistake**

- Forgetting the `return` after `http.Error`. It writes a response but does not
  stop the handler, so execution continues and writes a second one — logging
  `superfluous response.WriteHeader call`.

**Key detail:** harden the decode side: `http.MaxBytesReader(w, r.Body, 1<<20)`
caps the body so a client cannot exhaust memory, and
`dec.DisallowUnknownFields()` turns a misspelled field into an error instead of a
silent zero value.

**See also:** httpsrv1 · stdlib1 (tags) · errors1 · logingest3 ·
the [chapter](../README.md)

**References**

- pkg.go.dev — json.Decoder: https://pkg.go.dev/encoding/json#Decoder
- pkg.go.dev — http.MaxBytesReader: https://pkg.go.dev/net/http#MaxBytesReader

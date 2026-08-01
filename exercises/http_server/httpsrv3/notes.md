## httpsrv3 — JSON request and response

```go
var req echoReq
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(echoResp{Greeting: "Hello, " + req.Name + "!"})
```

**Why it works**

- `json.NewDecoder(r.Body).Decode` streams and parses the request body into a
  struct; `json.NewEncoder(w).Encode` streams the response struct straight to the
  client. On a bad body, `http.Error` writes a 400 and you `return`.

**Key detail:** prefer the streaming `Decoder`/`Encoder` over `Unmarshal`/`Marshal` in
handlers — they read/write directly from the body without buffering the whole
payload. Set `Content-Type` **before** writing the body (the first `Write` locks
the header). Always validate/handle the decode error.

**References**

- pkg.go.dev — encoding/json: https://pkg.go.dev/encoding/json

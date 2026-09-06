## logingest3 — routing with Go 1.22 method patterns

```go
func NewServer(store *Store) *Server {
	s := &Server{store: store, mux: http.NewServeMux()}
	s.mux.HandleFunc("POST /events", s.handleIngest)
	s.mux.HandleFunc("GET /events/{source}", s.handleBySource)
	return s
}

func (s *Server) handleIngest(w http.ResponseWriter, r *http.Request) {
	var e Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "malformed JSON body", http.StatusBadRequest)
		return
	}
	if err := s.store.Add(e); err != nil {
		if errors.Is(err, ErrInvalidEvent) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "could not store event", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (s *Server) handleBySource(w http.ResponseWriter, r *http.Request) {
	events := s.store.BySource(r.PathValue("source"))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "could not encode events", http.StatusInternalServerError)
	}
}
```

**Why it works**

- Before Go 1.22 a `ServeMux` matched on path only, so every handler began with
  a `switch r.Method` and a manual 405. The pattern `"POST /events"` moves that
  into the routing table, where it can be read at a glance.
- `"GET /events/{source}"` captures a path segment. `r.PathValue("source")`
  reads it back — no `strings.Split(r.URL.Path, "/")` and no off-by-one.
- The mux is built once in `NewServer` and stored. Registering on every request
  would re-register the same pattern repeatedly, which `ServeMux` panics on.
- `Server` has a `ServeHTTP` method, so it *is* an `http.Handler`. The mux stays
  unexported, and stage five can hand the whole `Server` to `http.Server` and
  stage six can wrap it in middleware, without either knowing there is a mux
  inside.

**Key detail:** the `405 Method Not Allowed` is free. When a path matches a
registered pattern but the method does not, `ServeMux` answers 405 itself and
sets `Allow` — your handler never runs. That only happens because the method is
part of the pattern; a bare `"/events"` registration would send `DELETE` to
`handleIngest`.

**Key detail:** `errors.Is(err, ErrInvalidEvent)` is doing real work here, not
ceremony. `Store.Add` can fail for a reason that is the *client's* fault (a bad
event → 400) or, in a bigger system, one that is the *server's* (→ 500). The
sentinel from stage one is what lets one `if` tell those apart. Matching on
`err.Error()` text would work until someone rewords the message.

**Key detail:** the empty-array test asserts on the raw body, not the decoded
value — `json.Unmarshal` happily decodes both `null` and `[]` into a nil slice,
so a decoded assertion would pass even when the bug is present. The stage-two
decision to return `[]Event{}` instead of `nil` is what makes the body `[]`.

**References**

- Go 1.22 release notes — enhanced ServeMux patterns: https://go.dev/blog/routing-enhancements
- pkg.go.dev — http.ServeMux: https://pkg.go.dev/net/http#ServeMux
- pkg.go.dev — http.Request.PathValue: https://pkg.go.dev/net/http#Request.PathValue
- pkg.go.dev — httptest.NewRecorder: https://pkg.go.dev/net/http/httptest#NewRecorder

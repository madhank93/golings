// logingest3
// Stage three puts the store behind HTTP. Go 1.22 gave ServeMux method and
// wildcard patterns, so the routing table is the mux itself rather than a
// hand-rolled switch on r.Method and a strings.Split of the path.

// I AM NOT DONE
package logingest

import "net/http"

// Server exposes a Store over HTTP. The mux is built once, in NewServer, and
// kept — registering routes per request would be both slow and racy.
type Server struct {
	store *Store
	mux   *http.ServeMux
}

// NewServer returns a Server that serves store.
func NewServer(store *Store) *Server {
	s := &Server{store: store, mux: http.NewServeMux()}
	// FIXME: register the two routes on s.mux with HandleFunc, using Go 1.22
	// method-and-wildcard patterns:
	//
	//	s.mux.HandleFunc("POST /events", s.handleIngest)
	//	s.mux.HandleFunc("GET /events/{source}", s.handleBySource)
	//
	// Right now nothing is registered, so the mux answers every request with
	// 404 and none of the tests below can pass.
	return s
}

// ServeHTTP delegates to the mux, which is what makes Server an http.Handler
// without exposing the mux itself to callers.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// handleIngest accepts one JSON-encoded Event.
//
// A body that is not valid JSON is 400 Bad Request. An event the store rejects
// is also 400 — the client sent something wrong either way. A stored event is
// 201 Created with an empty body.
func (s *Server) handleIngest(w http.ResponseWriter, r *http.Request) {
	// FIXME: decode r.Body into an Event with json.NewDecoder, hand it to
	// s.store.Add, and write the status codes described above. Use
	// errors.Is(err, ErrInvalidEvent) to recognise a rejected event rather
	// than comparing error strings.
	//
	// You will need to import "encoding/json" and "errors".
	//
	// Right now this writes nothing, so net/http sends 200 and the event is
	// never stored.
}

// handleBySource returns the stored events for one source as a JSON array.
func (s *Server) handleBySource(w http.ResponseWriter, r *http.Request) {
	// FIXME: read the source segment with r.PathValue("source") — that is how
	// you get at the {source} wildcard in the route pattern — ask the store
	// for its events, set Content-Type to application/json, and encode them.
	//
	// Note that Store.BySource returns an empty slice rather than nil for an
	// unknown source. That is deliberate: json encodes a nil slice as null and
	// an empty slice as [], and a client doing data.forEach() breaks on null.
	//
	// Right now this writes nothing, so the body is empty rather than [].
}

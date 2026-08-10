// logingest3
// Stage three puts the store behind HTTP. Go 1.22 gave ServeMux method and
// wildcard patterns, so the routing table is the mux itself rather than a
// hand-rolled switch on r.Method and a strings.Split of the path.

package logingest

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Server exposes a Store over HTTP. The mux is built once, in NewServer, and
// kept — registering routes per request would be both slow and racy.
type Server struct {
	store *Store
	mux   *http.ServeMux
}

// NewServer returns a Server that serves store.
func NewServer(store *Store) *Server {
	s := &Server{store: store, mux: http.NewServeMux()}
	s.mux.HandleFunc("POST /events", s.handleIngest)
	s.mux.HandleFunc("GET /events/{source}", s.handleBySource)
	return s
}

// ServeHTTP delegates to the mux, which is what makes Server an http.Handler
// without exposing the mux itself to callers.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// handleIngest accepts one JSON-encoded Event.
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

// handleBySource returns the stored events for one source as a JSON array.
func (s *Server) handleBySource(w http.ResponseWriter, r *http.Request) {
	events := s.store.BySource(r.PathValue("source"))
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "could not encode events", http.StatusInternalServerError)
	}
}

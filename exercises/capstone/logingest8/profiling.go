// logingest8
// Stage eight is the last one: expose the runtime profiles, then assemble the
// whole service. net/http/pprof is written for the default mux — its init()
// registers the handlers on http.DefaultServeMux and nothing else — so putting
// it on a mux you own takes one deliberate step.

// I AM NOT DONE
package logingest

import (
	"log/slog"
	"net/http"
)

// PprofPrefix is the path all the runtime profile endpoints live under.
const PprofPrefix = "/debug/pprof/"

// WithPprof registers the net/http/pprof handlers on mux and returns it.
//
// The handlers are registered explicitly rather than by forwarding the whole of
// http.DefaultServeMux, so mux exposes the profiles and nothing else.
func WithPprof(mux *http.ServeMux) *http.ServeMux {
	// FIXME: register the five pprof handlers on mux:
	//
	//	mux.HandleFunc(PprofPrefix, pprof.Index)
	//	mux.HandleFunc(PprofPrefix+"cmdline", pprof.Cmdline)
	//	mux.HandleFunc(PprofPrefix+"profile", pprof.Profile)
	//	mux.HandleFunc(PprofPrefix+"symbol", pprof.Symbol)
	//	mux.HandleFunc(PprofPrefix+"trace", pprof.Trace)
	//
	// pprof.Index also serves the named profiles — /debug/pprof/heap,
	// /goroutine, /allocs and the rest — so those five registrations cover
	// everything.
	//
	// Import it as `"net/http/pprof"`, not with a blank name: you are calling
	// its functions here, not relying on its init().
	//
	// Right now mux comes back untouched, so every profile path 404s.
	return mux
}

// NewService assembles the capstone into one http.Handler: the event routes
// from stage three, the runtime profiles above, and request logging wrapped
// around both.
func NewService(store *Store, log *slog.Logger) http.Handler {
	// FIXME: build the service.
	//
	//  1. mux := http.NewServeMux()
	//  2. mount the stage-three Server at the root: mux.Handle("/", NewServer(store)).
	//     ServeMux picks the longest matching pattern, so the pprof paths
	//     registered below still win over this catch-all.
	//  3. WithPprof(mux)
	//  4. return WithLogging(mux, log) — the logger goes outermost so that
	//     profile requests are logged too.
	//
	// Right now only the event routes are wired up: no profiles, no logging.
	return NewServer(store)
}

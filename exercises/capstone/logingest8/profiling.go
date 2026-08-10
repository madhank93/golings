// logingest8
// Stage eight is the last one: expose the runtime profiles, then assemble the
// whole service. net/http/pprof is written for the default mux — its init()
// registers the handlers on http.DefaultServeMux and nothing else — so putting
// it on a mux you own takes one deliberate step.

package logingest

import (
	"log/slog"
	"net/http"
	"net/http/pprof"
)

// PprofPrefix is the path all the runtime profile endpoints live under.
const PprofPrefix = "/debug/pprof/"

// WithPprof registers the net/http/pprof handlers on mux and returns it.
//
// The handlers are registered explicitly rather than by forwarding the whole of
// http.DefaultServeMux, so mux exposes the profiles and nothing else.
func WithPprof(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(PprofPrefix, pprof.Index)
	mux.HandleFunc(PprofPrefix+"cmdline", pprof.Cmdline)
	mux.HandleFunc(PprofPrefix+"profile", pprof.Profile)
	mux.HandleFunc(PprofPrefix+"symbol", pprof.Symbol)
	mux.HandleFunc(PprofPrefix+"trace", pprof.Trace)
	return mux
}

// NewService assembles the capstone into one http.Handler: the event routes
// from stage three, the runtime profiles above, and request logging wrapped
// around both.
//
// The logger goes outermost so profile requests are logged too, and ServeMux
// prefers the longest matching pattern, so the pprof paths win over the "/"
// catch-all regardless of registration order.
func NewService(store *Store, log *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", NewServer(store))
	WithPprof(mux)
	return WithLogging(mux, log)
}

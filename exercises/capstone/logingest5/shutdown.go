// logingest5
// Stage five is shutdown. A service that dies the instant it is signalled drops
// the requests it is in the middle of answering; a service that never dies has
// to be killed. Graceful shutdown is the middle: stop accepting new work, let
// the work in flight finish, then give up after a deadline.

// I AM NOT DONE
package logingest

import (
	"context"
	"net"
	"net/http"
	"time"
)

// Serve serves h on ln until ctx is cancelled, then shuts down gracefully,
// giving in-flight requests up to grace to finish.
//
// It returns nil for a clean shutdown. http.ErrServerClosed is what Serve
// reports after a Shutdown, and it means everything went to plan, so it is not
// propagated to the caller. A failure to accept connections is a real error and
// is returned.
func Serve(ctx context.Context, ln net.Listener, h http.Handler, grace time.Duration) error {
	srv := &http.Server{
		Handler: h,
		// Without a read timeout a client that opens a connection and never
		// sends anything holds a slot indefinitely.
		ReadHeaderTimeout: 5 * time.Second,
	}

	// FIXME: three things are missing.
	//
	//  1. srv.Serve(ln) has to run in its own goroutine, reporting whatever it
	//     returns on a buffered channel. Buffered by one, so the goroutine can
	//     always send and exit even when nobody is listening any more —
	//     otherwise it leaks.
	//
	//  2. select on that channel and on ctx.Done(). A value on the channel is
	//     the server stopping by itself; treat http.ErrServerClosed as success
	//     (errors.Is) and anything else as the error to return.
	//
	//  3. when ctx is cancelled, build a *separate* deadline context for the
	//     shutdown and call srv.Shutdown with it:
	//
	//		shutCtx, cancel := context.WithTimeout(context.Background(), grace)
	//		defer cancel()
	//		return srv.Shutdown(shutCtx)
	//
	//     It must derive from Background, not from ctx. ctx is already
	//     cancelled by this point, so a context derived from it is born
	//     expired and Shutdown would kill in-flight requests immediately —
	//     which is the exact thing you are trying to avoid.
	//
	// You will need to import "errors".
	//
	// Right now this returns before serving anything, so every request fails.
	_ = srv
	return nil
}

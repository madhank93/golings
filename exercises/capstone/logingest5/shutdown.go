// logingest5
// Stage five is shutdown. A service that dies the instant it is signalled drops
// the requests it is in the middle of answering; a service that never dies has
// to be killed. Graceful shutdown is the middle: stop accepting new work, let
// the work in flight finish, then give up after a deadline.

package logingest

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

// Serve serves h on ln until ctx is cancelled, then shuts down gracefully,
// giving in-flight requests up to grace to finish.
func Serve(ctx context.Context, ln net.Listener, h http.Handler, grace time.Duration) error {
	srv := &http.Server{
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Buffered by one so this goroutine can always send and exit, even when
	// the select below has already moved on.
	serveErr := make(chan error, 1)
	go func() { serveErr <- srv.Serve(ln) }()

	select {
	case err := <-serveErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		// Derived from Background, not ctx: ctx is already cancelled, and a
		// context born expired would make Shutdown kill in-flight requests
		// immediately.
		shutCtx, cancel := context.WithTimeout(context.Background(), grace)
		defer cancel()
		return srv.Shutdown(shutCtx)
	}
}

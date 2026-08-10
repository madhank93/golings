package logingest

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

// listen returns a listener on a free loopback port plus the base URL for it.
// Port 0 lets the kernel pick, so parallel test runs never collide.
func listen(t *testing.T) (net.Listener, string) {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}
	return ln, "http://" + ln.Addr().String()
}

// start runs Serve in the background and hands back a channel carrying its
// return value, so a test can assert on both the timing and the error.
func start(ctx context.Context, ln net.Listener, h http.Handler, grace time.Duration) <-chan error {
	errc := make(chan error, 1)
	go func() { errc <- Serve(ctx, ln, h, grace) }()
	return errc
}

// Every request in this file goes through a client with a timeout. The
// listener is already open before Serve runs, so a Serve that never accepts
// still completes the TCP handshake — a bare http.Get would block forever
// rather than fail, and the test would hang instead of reporting.
func client(timeout time.Duration) *http.Client {
	return &http.Client{Timeout: timeout}
}

// waitReady polls until the server answers, so a slow start is not a flake.
func waitReady(t *testing.T, url string) {
	t.Helper()
	c := client(200 * time.Millisecond)
	deadline := time.Now().Add(3 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := c.Get(url)
		if err == nil {
			resp.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatal("server never became ready: Serve is not serving on the listener")
}

func TestServeAnswersRequestsWhileContextIsLive(t *testing.T) {
	ln, url := listen(t)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "ok")
	})
	errc := start(ctx, ln, h, time.Second)
	waitReady(t, url)

	resp, err := client(2 * time.Second).Get(url)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "ok" {
		t.Errorf("body = %q, want %q", body, "ok")
	}

	cancel()
	if err := <-errc; err != nil {
		t.Errorf("Serve returned %v, want nil", err)
	}
}

// Cancelling the context must actually stop the server, and reasonably fast.
func TestServeReturnsAfterCancel(t *testing.T) {
	ln, url := listen(t)
	ctx, cancel := context.WithCancel(context.Background())

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {})
	errc := start(ctx, ln, h, time.Second)
	waitReady(t, url)

	cancel()
	select {
	case err := <-errc:
		if err != nil {
			t.Errorf("Serve returned %v, want nil for a clean shutdown", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("Serve ignored the cancelled context and kept running")
	}
}

// The point of graceful shutdown: a request already being handled when the
// signal arrives still gets its response.
func TestServeLetsInFlightRequestFinish(t *testing.T) {
	ln, url := listen(t)
	ctx, cancel := context.WithCancel(context.Background())

	// The readiness probe and the slow request need separate routes: probing
	// the slow handler would both trip the `started` signal early and time the
	// probe out.
	started := make(chan struct{})
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ready", func(w http.ResponseWriter, _ *http.Request) {})
	mux.HandleFunc("GET /slow", func(w http.ResponseWriter, _ *http.Request) {
		close(started)
		// Long enough that the cancel below lands mid-request, well short of
		// the grace period so a correct implementation always finishes.
		time.Sleep(300 * time.Millisecond)
		fmt.Fprint(w, "finished")
	})
	errc := start(ctx, ln, mux, 5*time.Second)
	waitReady(t, url+"/ready")

	type result struct {
		body string
		err  error
	}
	resc := make(chan result, 1)
	go func() {
		resp, err := client(10 * time.Second).Get(url + "/slow")
		if err != nil {
			resc <- result{err: err}
			return
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		resc <- result{body: string(b), err: err}
	}()

	<-started
	cancel() // shutdown begins while the handler is still running

	select {
	case res := <-resc:
		if res.err != nil {
			t.Fatalf("in-flight request failed during shutdown: %v", res.err)
		}
		if res.body != "finished" {
			t.Errorf("body = %q, want %q", res.body, "finished")
		}
	case <-time.After(10 * time.Second):
		t.Fatal("in-flight request never completed")
	}

	if err := <-errc; err != nil {
		t.Errorf("Serve returned %v, want nil", err)
	}
}

// Once Serve has returned, the port is released.
func TestServeStopsAcceptingAfterShutdown(t *testing.T) {
	ln, url := listen(t)
	ctx, cancel := context.WithCancel(context.Background())

	h := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {})
	errc := start(ctx, ln, h, time.Second)
	waitReady(t, url)

	cancel()
	if err := <-errc; err != nil {
		t.Fatalf("Serve returned %v, want nil", err)
	}

	if resp, err := client(time.Second).Get(url); err == nil {
		resp.Body.Close()
		t.Error("server still accepted a request after Serve returned")
	}
}

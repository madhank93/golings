// tcp3
// Run an accept loop that stops cleanly when the listener closes.

package main_test

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

// Serve accepts connections until ln is closed, answering each one in turn.
// Every connection gets one line echoed back uppercased.
//
// Closing a listener is the ordinary way to stop a server, so the Accept error
// that follows is not a failure: net.ErrClosed means "we are shutting down"
// and Serve returns nil. Any other Accept error is real and is returned.
func Serve(ln net.Listener) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		line, err := bufio.NewReader(conn).ReadString('\n')
		if err == nil {
			fmt.Fprintf(conn, "%s\n", strings.ToUpper(strings.TrimSuffix(line, "\n"))) //nolint:errcheck
		}
		conn.Close()
	}
}

func ask(t *testing.T, addr, cmd string) string {
	t.Helper()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	if _, err := fmt.Fprintf(conn, "%s\n", cmd); err != nil {
		t.Fatal(err)
	}
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	return reply
}

func TestServeHandlesManyConnections(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- Serve(ln) }()

	for _, cmd := range []string{"one", "two", "three"} {
		want := strings.ToUpper(cmd) + "\n"
		if got := ask(t, ln.Addr().String(), cmd); got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}

	// Closing the listener must stop Serve, and stop it without an error.
	ln.Close()
	select {
	case err := <-done:
		if err != nil {
			t.Errorf("Serve after close: %v, want nil", err)
		}
	case <-time.After(5 * time.Second):
		t.Error("Serve did not return after the listener was closed")
	}
}

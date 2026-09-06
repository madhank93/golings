// tcp1
// Accept one TCP connection and answer it.

package main_test

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

// ServeOnce accepts a single connection on ln, reads one newline-terminated
// line from it, writes "PONG\n" back if the line was "PING", and closes the
// connection. It returns after that one exchange.
func ServeOnce(ln net.Listener) error {
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	if line != "PING\n" {
		return fmt.Errorf("unexpected command %q", line)
	}
	_, err = conn.Write([]byte("PONG\n"))
	return err
}

func TestServeOnce(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	done := make(chan error, 1)
	go func() { done <- ServeOnce(ln) }()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}

	if _, err := fmt.Fprint(conn, "PING\n"); err != nil {
		t.Fatal(err)
	}
	got, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatalf("reading the reply: %v", err)
	}
	if got != "PONG\n" {
		t.Errorf("got %q, want %q", got, "PONG\n")
	}
	if err := <-done; err != nil {
		t.Errorf("ServeOnce: %v", err)
	}
}

func TestServeOnceClosesTheConnection(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	go ServeOnce(ln) //nolint:errcheck

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	fmt.Fprint(conn, "PING\n") //nolint:errcheck

	// The server closing its side is what ends this read. A handler that
	// returns without closing leaves the client waiting for the deadline.
	r := bufio.NewReader(conn)
	if _, err := r.ReadString('\n'); err != nil {
		t.Fatal(err)
	}
	if _, err := r.ReadString('\n'); err == nil {
		t.Error("the server never closed the connection")
	}
}

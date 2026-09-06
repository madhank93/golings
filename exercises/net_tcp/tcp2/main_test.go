// tcp2
// Frame a byte stream into messages.

// I AM NOT DONE
package main_test

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

// Handle serves one connection, reading newline-terminated commands until the
// client closes its side, and answering each one. Every command is echoed back
// uppercased, one line per line, and "QUIT" ends the conversation.
//
// TCP delivers a byte stream, not messages: one Read can return half a command
// or three of them, so the newline is the only thing that says where a command
// ends.
func Handle(conn net.Conn) error {
	defer conn.Close()

	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		return nil
	}
	// FIXME: one command is not enough. Loop until the read fails (a clean
	// client close is not an error to report), answer every command, and
	// return when the command is "QUIT".
	_, err = fmt.Fprintf(conn, "%s\n", strings.ToUpper(strings.TrimSuffix(line, "\n")))
	return err
}

func dialHandled(t *testing.T) net.Conn {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ln.Close() })

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		Handle(conn) //nolint:errcheck
	}()

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Fatal(err)
	}
	return conn
}

func TestHandleManyCommandsOnOneConnection(t *testing.T) {
	conn := dialHandled(t)
	r := bufio.NewReader(conn)

	for _, want := range []string{"ONE", "TWO", "THREE"} {
		if _, err := fmt.Fprintf(conn, "%s\n", strings.ToLower(want)); err != nil {
			t.Fatal(err)
		}
		got, err := r.ReadString('\n')
		if err != nil {
			t.Fatal(err)
		}
		if got != want+"\n" {
			t.Fatalf("got %q, want %q", got, want+"\n")
		}
	}
}

// A single Write carrying three commands, and a command split across two
// Writes: both are legal TCP and both must produce three replies.
func TestHandleReframesWhatArrives(t *testing.T) {
	conn := dialHandled(t)
	r := bufio.NewReader(conn)

	if _, err := fmt.Fprint(conn, "one\ntwo\nthr"); err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"ONE\n", "TWO\n"} {
		got, err := r.ReadString('\n')
		if err != nil {
			t.Fatal(err)
		}
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	}

	if _, err := fmt.Fprint(conn, "ee\n"); err != nil {
		t.Fatal(err)
	}
	got, err := r.ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}
	if got != "THREE\n" {
		t.Errorf("a command split across two writes came back as %q", got)
	}
}

func TestHandleQuit(t *testing.T) {
	conn := dialHandled(t)
	if _, err := fmt.Fprint(conn, "QUIT\n"); err != nil {
		t.Fatal(err)
	}
	if _, err := bufio.NewReader(conn).ReadString('\n'); err == nil {
		t.Error("QUIT should end the conversation, not be echoed")
	}
}

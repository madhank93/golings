// tcp4
// Serve connections concurrently, one goroutine each.

package main_test

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

// Counter is shared by every connection, so its methods have to be safe to
// call from several goroutines at once.
type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Next() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
	return c.n
}

// ServeConcurrent accepts connections until ln is closed and handles each one
// in its own goroutine, so a slow client cannot block a fast one. Every
// connection is answered with its sequence number from c.
//
// The accept loop itself stays single-threaded: one goroutine accepts, and the
// work goes elsewhere. That is the shape of every Go server, http.Server
// included.
func ServeConcurrent(ln net.Listener, c *Counter) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer conn.Close()

			line, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				return
			}
			// A slow client holds only its own goroutine.
			if strings.TrimSuffix(line, "\n") == "SLOW" {
				time.Sleep(300 * time.Millisecond)
			}
			fmt.Fprintf(conn, "%d\n", c.Next()) //nolint:errcheck
		}()
	}
}

func request(t *testing.T, addr, cmd string) string {
	t.Helper()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
		return ""
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		t.Error(err)
		return ""
	}
	if _, err := fmt.Fprintf(conn, "%s\n", cmd); err != nil {
		t.Error(err)
		return ""
	}
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Error(err)
		return ""
	}
	return reply
}

// Six slow clients at once. Handled one at a time their waits add up; handled
// concurrently they overlap.
func TestServeConcurrentDoesNotHeadOfLineBlock(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	c := &Counter{}
	go ServeConcurrent(ln, c) //nolint:errcheck

	start := time.Now()
	var wg sync.WaitGroup
	for range 6 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if reply := request(t, ln.Addr().String(), "SLOW"); reply == "" {
				t.Error("no reply")
			}
		}()
	}
	wg.Wait()

	// Six clients, 300ms of work each. Served concurrently that is about
	// 300ms; served one at a time it is about 1.8s.
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("six connections took %s — are they being served one at a time?", elapsed)
	}
}

// Six connections must produce six distinct numbers. Run with -race: an
// unsynchronised counter is reported here even when the numbers happen to
// come out right.
func TestServeConcurrentCountsEveryConnection(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	c := &Counter{}
	go ServeConcurrent(ln, c) //nolint:errcheck

	var mu sync.Mutex
	seen := map[string]bool{}

	var wg sync.WaitGroup
	for range 6 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			reply := request(t, ln.Addr().String(), "fast")
			mu.Lock()
			defer mu.Unlock()
			seen[reply] = true
		}()
	}
	wg.Wait()

	if len(seen) != 6 {
		t.Errorf("got %d distinct replies from 6 connections: %v", len(seen), seen)
	}
}

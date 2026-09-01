package main

import "testing"

// The playground answering at all — even with a compile error or a failing
// test — means the snippet is runnable. Only the sandbox giving up counts.
func TestTimedOut(t *testing.T) {
	ev := func(msgs ...string) result {
		r := result{}
		for _, m := range msgs {
			r.Events = append(r.Events, struct{ Message string }{m})
		}
		return r
	}
	cases := []struct {
		name string
		in   result
		want bool
	}{
		{"stub compile error is runnable", result{Errors: "./prog.go:12:6: syntax error: unexpected =, expected name\n"}, false},
		{"failing test is runnable", ev("=== RUN   TestUserRoute\n    prog_test.go:34: want 200, got 404\n--- FAIL: TestUserRoute"), false},
		{"passing test is runnable", ev("PASS\n"), false},
		{"build timeout is blocked", result{Errors: "timeout running go build"}, true},
		{"slow process is blocked", ev("hello\n", "process took too long"), true},
	}
	for _, c := range cases {
		if got := timedOut(c.in); got != c.want {
			t.Errorf("%s: timedOut = %v, want %v", c.name, got, c.want)
		}
	}
}

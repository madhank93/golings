// mock1
// A spy is a test double that records how it was called, so you can assert on it.

// I AM NOT DONE
package main_test

import "testing"

// Notifier sends a message somewhere (email, Slack, …).
type Notifier interface {
	Notify(msg string)
}

// Alert notifies when temp exceeds threshold, and stays silent otherwise.
func Alert(n Notifier, temp, threshold int) {
	// FIXME: only notify when the temperature is above the threshold.
}

// spyNotifier records every message it is asked to send.
type spyNotifier struct{ messages []string }

func (s *spyNotifier) Notify(msg string) { s.messages = append(s.messages, msg) }

func TestAlertFiresWhenTooHot(t *testing.T) {
	spy := &spyNotifier{}
	Alert(spy, 30, 25)
	if len(spy.messages) != 1 || spy.messages[0] != "too hot" {
		t.Errorf("want one \"too hot\" notification, got %v", spy.messages)
	}
}

func TestAlertSilentWhenCool(t *testing.T) {
	spy := &spyNotifier{}
	Alert(spy, 20, 25)
	if len(spy.messages) != 0 {
		t.Errorf("want no notifications, got %v", spy.messages)
	}
}

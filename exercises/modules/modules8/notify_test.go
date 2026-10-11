package notify_test

import (
	"testing"

	"example.com/notify"
)

func TestSendEmail(t *testing.T) {
	got, err := notify.Send("email", "ana@example.com", "hi")
	if err != nil {
		t.Fatal(err)
	}
	if want := "mailed ana@example.com: hi"; got != want {
		t.Errorf("Send = %q, want %q", got, want)
	}
}

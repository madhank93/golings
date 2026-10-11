package bank_test

import (
	"testing"

	"example.com/bank"
)

func TestStatement(t *testing.T) {
	tests := []struct {
		cents int
		want  string
	}{
		{1250, "Balance: $12.50"},
		{-150, "Balance: -$1.50"},
	}
	for _, tt := range tests {
		if got := bank.Statement(tt.cents); got != tt.want {
			t.Errorf("Statement(%d) = %q, want %q", tt.cents, got, tt.want)
		}
	}
}

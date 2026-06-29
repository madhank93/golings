// files1 — reading and writing files
// os.WriteFile and os.ReadFile handle whole-file I/O in a single call.

package main_test

import (
	"os"
	"path/filepath"
	"testing"
)

func saveAndLoad(path, content string) (string, error) {
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TestSaveAndLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.txt")
	got, err := saveAndLoad(path, "hello files")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "hello files" {
		t.Errorf("want %q, got %q", "hello files", got)
	}
}

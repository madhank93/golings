package tui

import (
	"path/filepath"
	"testing"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/madhank93/golings/golings/exercises"
)

func wheelMsg(x int, btn tea.MouseButton) tea.MouseMsg {
	return tea.MouseMsg{X: x, Y: 10, Button: btn, Action: tea.MouseActionPress}
}

// notesModel points at real exercises, so Notes() and Chapter() return content.
func notesModel(t *testing.T) Model {
	t.Helper()
	tr, _ := exercises.LoadState(filepath.Join(t.TempDir(), "s.json"))
	exs := []exercises.Exercise{
		{Name: "sync2", Path: "../../exercises/sync/sync2/main_test.go", Mode: "test", Hint: "once"},
		{Name: "sync3", Path: "../../exercises/sync/sync3/main_test.go", Mode: "test", Hint: "atomic"},
	}
	m := Model{
		tracker: tr, phase: phaseMain, keys: defaultKeys(), help: help.New(),
		progress: progress.New(), spinner: spinner.New(), output: viewport.New(0, 0),
		total: len(exs),
	}
	m.items = buildItems(exs, tr)
	m.cursor = m.firstSelectable()
	return m
}

// The Learn section is long (note plus the whole topic chapter), so a re-verify
// on every save must not throw a reader back to where the section starts.
func TestReverifyKeepsScrollPosition(t *testing.T) {
	m := notesModel(t)
	nm, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
	m = nm.(Model)

	if m.current().Notes() == "" || m.current().Chapter() == "" {
		t.Fatalf("test needs an exercise with notes and a chapter")
	}

	// Open the Learn section: this transition does scroll to it.
	nm, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'x'}})
	m = nm.(Model)
	if m.output.YOffset != m.notesTop {
		t.Errorf("opening notes should scroll to the Learn section: want %d, got %d", m.notesTop, m.output.YOffset)
	}

	for range 5 {
		nm, _ = m.Update(wheelMsg(m.leftPaneW+5, tea.MouseButtonWheelDown))
		m = nm.(Model)
	}
	scrolled := m.output.YOffset
	if scrolled <= m.notesTop {
		t.Fatalf("wheel over the right pane did not scroll: %d", scrolled)
	}

	nm, _ = m.Update(verifiedMsg{name: "sync2", status: exercises.StatusDone, result: exercises.Result{Out: "ok"}})
	m = nm.(Model)

	if m.output.YOffset < scrolled {
		t.Errorf("re-verify yanked the viewport back: was %d, now %d", scrolled, m.output.YOffset)
	}
}

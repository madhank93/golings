package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func press(t *testing.T, m Model, r rune) Model {
	t.Helper()
	nm, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
	return nm.(Model)
}

// The chapter is one document per topic. Showing it with every exercise's
// walk-through repeated the same pages under all six variables exercises, so it
// lives behind its own key.
func TestChapterIsNotPartOfTheWalkthrough(t *testing.T) {
	m := notesModel(t)
	nm, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 30})
	m = nm.(Model)

	m = press(t, m, 'x')
	if !m.showNotes || m.showChapter {
		t.Fatalf("x should open the note alone: notes=%v chapter=%v", m.showNotes, m.showChapter)
	}
	withNote := m.output.TotalLineCount()

	m = press(t, m, 'c')
	if !m.showChapter {
		t.Fatal("c should open the chapter")
	}
	if m.output.TotalLineCount() <= withNote {
		t.Error("opening the chapter added nothing to the pane")
	}
	if m.output.YOffset != m.chapterTop {
		t.Errorf("opening the chapter should scroll to it: want %d, got %d", m.chapterTop, m.output.YOffset)
	}

	m = press(t, m, 'c')
	if m.showChapter {
		t.Error("c should toggle the chapter back off")
	}

	// Moving to another exercise starts from a clean slate.
	m = press(t, m, 'c')
	nm, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = nm.(Model)
	if m.showChapter || m.showNotes {
		t.Errorf("selection change should reset the panes: notes=%v chapter=%v", m.showNotes, m.showChapter)
	}
}

// The key line in the header is hand-written, so it drifts from the bindings.
func TestHeaderKeyLineMentionsBothKeys(t *testing.T) {
	m := notesModel(t)
	nm, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 30})
	m = nm.(Model)
	v := m.View()
	for _, want := range []string{"x explain", "c chapter"} {
		if !strings.Contains(v, want) {
			t.Errorf("header key line is missing %q", want)
		}
	}
}

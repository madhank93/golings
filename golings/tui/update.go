package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mauricioabreu/golings/golings/exercises"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.layout()
		m.output.SetContent(m.detail())
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)

	case spinner.TickMsg:
		if m.verifying {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
		return m, nil

	case fileChangedMsg:
		// re-verify the current exercise on any save; keep listening
		m.verifying = true
		m.notice = ""
		return m, tea.Batch(waitForChange(m.watchCh), verifyCmd(m.current()), m.spinner.Tick)

	case verifiedMsg:
		return m.handleVerified(msg)
	}

	return m, nil
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, m.keys.Up):
		m.moveCursor(-1)
		m.onSelectionChange()
		m.output.SetContent(m.detail())
		return m, nil

	case key.Matches(msg, m.keys.Down):
		m.moveCursor(1)
		m.onSelectionChange()
		m.output.SetContent(m.detail())
		return m, nil

	case key.Matches(msg, m.keys.Run):
		m.verifying = true
		m.notice = ""
		return m, tea.Batch(verifyCmd(m.current()), m.spinner.Tick)

	case key.Matches(msg, m.keys.Hint):
		m.showHint = !m.showHint
		m.output.SetContent(m.detail())
		return m, nil

	case key.Matches(msg, m.keys.Reset):
		cur := m.current()
		if err := exercises.Reset(cur); err != nil {
			m.notice = "reset failed: " + err.Error()
		} else {
			m.tracker.Unmark(cur.Name)
			_ = m.tracker.Save()
			m.refreshHeaderCounts()
			m.notice = "reset " + cur.Name + " to original"
		}
		m.hasResult = false
		m.status = exercises.StatusPending
		m.verifying = true
		return m, tea.Batch(verifyCmd(cur), m.spinner.Tick)

	case key.Matches(msg, m.keys.Next):
		m.advance()
		m.onSelectionChange()
		m.verifying = true
		return m, tea.Batch(verifyCmd(m.current()), m.spinner.Tick)
	}

	// let the viewport scroll (pgup/pgdn etc.)
	var cmd tea.Cmd
	m.output, cmd = m.output.Update(msg)
	return m, cmd
}

func (m Model) handleVerified(msg verifiedMsg) (tea.Model, tea.Cmd) {
	// ignore stale results for an exercise we've navigated away from
	if msg.name != m.current().Name {
		return m, nil
	}
	m.verifying = false
	m.status = msg.status
	m.result = msg.result
	m.hasResult = true

	if msg.status == exercises.StatusDone && !m.tracker.IsDone(msg.name) {
		m.tracker.MarkDone(msg.name, now())
		_ = m.tracker.Save()
		m.refreshHeaderCounts()
		m.advance()
		m.onSelectionChange()
		m.output.SetContent(m.detail())
		// verify the newly-selected exercise too
		m.verifying = true
		return m, tea.Batch(verifyCmd(m.current()), m.spinner.Tick)
	}

	m.output.SetContent(m.detail())
	return m, nil
}

// onSelectionChange resets per-exercise detail state when the cursor moves.
func (m *Model) onSelectionChange() {
	m.showHint = false
	m.hasResult = false
	m.result = exercises.Result{}
	m.output.GotoTop()
}

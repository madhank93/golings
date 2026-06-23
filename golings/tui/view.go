package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mauricioabreu/golings/golings/exercises"
)

// Adaptive colors so the UI reads well on both light and dark terminals.
var (
	cBlue  = lipgloss.AdaptiveColor{Light: "#0550ae", Dark: "#79c0ff"}
	cLink  = lipgloss.AdaptiveColor{Light: "#0969da", Dark: "#58a6ff"}
	cGreen = lipgloss.AdaptiveColor{Light: "#1a7f37", Dark: "#3fb950"}
	cRed   = lipgloss.AdaptiveColor{Light: "#cf222e", Dark: "#f85149"}
	cAmber = lipgloss.AdaptiveColor{Light: "#9a6700", Dark: "#d29922"}
	cTeal  = lipgloss.AdaptiveColor{Light: "#0e7490", Dark: "#39c5cf"}
	cDim   = lipgloss.AdaptiveColor{Light: "#6e7781", Dark: "#8b949e"}
)

var (
	topicStyle    = lipgloss.NewStyle().Bold(true).Foreground(cBlue)
	selectedStyle = lipgloss.NewStyle().Bold(true).Reverse(true) // inverts fg/bg — readable on any theme
	doneStyle     = lipgloss.NewStyle().Foreground(cGreen)
	failStyle     = lipgloss.NewStyle().Foreground(cRed)
	lintStyle     = lipgloss.NewStyle().Foreground(cAmber)
	passStyle     = lipgloss.NewStyle().Foreground(cTeal)
	dimStyle      = lipgloss.NewStyle().Foreground(cDim)
	titleStyle    = lipgloss.NewStyle().Bold(true)
	noticeStyle   = lipgloss.NewStyle().Foreground(cAmber)
	paneStyle     = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)
)

// topicOf extracts the topic directory from an exercise path.
func topicOf(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 2 {
		return parts[1]
	}
	return path
}

func (m *Model) layout() {
	headerH, footerH := 3, 2
	bodyH := m.height - headerH - footerH
	if bodyH < 3 {
		bodyH = 3
	}
	leftW := m.width * 4 / 10
	if leftW < 28 {
		leftW = 28
	}
	rightW := m.width - leftW - 6 // borders + padding
	if rightW < 20 {
		rightW = 20
	}
	m.output.Width = rightW
	m.output.Height = bodyH - 3 // leave a title line inside the right pane
	// leave room on the header line for the title + "n/N (p%)  🔥 streak N"
	m.progress.Width = max(10, m.width-46)
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading…"
	}
	if m.phase == phaseWelcome {
		return m.welcome()
	}
	return strings.Join([]string{
		m.header(),
		m.body(),
		m.footer(),
	}, "\n")
}

var (
	linkStyle  = lipgloss.NewStyle().Foreground(cLink).Underline(true)
	labelStyle = lipgloss.NewStyle().Foreground(cDim).Width(12)
	markStyle  = lipgloss.NewStyle().Foreground(cAmber)
	boxStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
			BorderForeground(cTeal).Padding(1, 3)
)

// welcome renders the splash screen shown before the main view.
func (m Model) welcome() string {
	title := titleStyle.Foreground(cTeal).Render("🐹  golings")
	tagline := dimStyle.Render("Learn Go the rustlings way — 97 exercises, basics → advanced")

	meta := lipgloss.JoinVertical(lipgloss.Left,
		labelStyle.Render("Repo")+linkStyle.Render("https://github.com/madhank93/golings"),
		labelStyle.Render("Site")+linkStyle.Render("https://golings.madhan.app"),
		labelStyle.Render("Maintainer")+"Madhan Kumaravelu  "+dimStyle.Render("(@madhank93)"),
	)

	how := lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("How it works"),
		"  • Open the highlighted exercise's file and make it compile/pass",
		"  • Remove the  "+markStyle.Render("// I AM NOT DONE")+"  marker when you think it's done",
		"  • Save — golings auto-runs it; "+titleStyle.Render("tests AND golangci-lint")+" must pass",
		"  • It then auto-advances to the next exercise",
	)

	keys := dimStyle.Render("Keys   ↑↓/jk move · ⏎ run · e edit · h hint · r reset · n next · q quit")
	cta := markStyle.Render("press any key to start →")

	content := lipgloss.JoinVertical(lipgloss.Left,
		title, tagline, "", meta, "", how, "", keys, "", cta,
	)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, boxStyle.Render(content))
}

func (m Model) header() string {
	done := m.tracker.DoneCount()
	ratio := 0.0
	if m.total > 0 {
		ratio = float64(done) / float64(m.total)
	}
	bar := m.progress.ViewAs(ratio)
	stats := fmt.Sprintf(" %d/%d (%.0f%%)   🔥 streak %d", done, m.total, ratio*100, m.tracker.Streak.Count)
	return titleStyle.Render("golings") + "  " + bar + stats
}

func (m Model) body() string {
	headerH, footerH := 3, 2
	bodyH := m.height - headerH - footerH
	if bodyH < 3 {
		bodyH = 3
	}
	leftW := m.width * 4 / 10
	if leftW < 28 {
		leftW = 28
	}

	left := paneStyle.Width(leftW).Height(bodyH).Render(m.list(bodyH))
	right := paneStyle.Width(m.output.Width).Height(bodyH).Render(m.rightPane())
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

// list renders the windowed topic/exercise list around the cursor.
func (m Model) list(height int) string {
	var lines []string
	for i, it := range m.items {
		if it.isHeader {
			lines = append(lines, topicStyle.Render(fmt.Sprintf("%s (%d/%d)", it.topic, it.done, it.total)))
			continue
		}
		row := "  " + m.glyph(it.ex) + " " + it.ex.Name
		if i == m.cursor {
			row = selectedStyle.Render("→ " + m.glyph(it.ex) + " " + it.ex.Name)
		}
		lines = append(lines, row)
	}
	return windowLines(lines, m.cursorLine(), height)
}

// glyph returns the colored status marker for an exercise row (cheap, no run).
func (m Model) glyph(ex exercises.Exercise) string {
	switch {
	case m.tracker.IsDone(ex.Name):
		return doneStyle.Render("✓")
	case ex.State() == exercises.Done: // marker removed but not verified done
		return lintStyle.Render("●")
	default:
		return dimStyle.Render("○")
	}
}

func (m Model) rightPane() string {
	ex := m.current()
	title := titleStyle.Render(ex.Name) + dimStyle.Render(" ["+ex.Mode+"] ") + m.badge()
	desc := ""
	if d := ex.Description(); d != "" {
		desc = "\n" + dimStyle.Italic(true).Render(d)
	}
	notice := ""
	if m.notice != "" {
		notice = "\n" + noticeStyle.Render(m.notice)
	}
	return title + desc + notice + "\n" + m.output.View()
}

func (m Model) badge() string {
	if m.verifying {
		return m.spinner.View() + " running"
	}
	if !m.hasResult {
		return dimStyle.Render("not run")
	}
	switch m.status {
	case exercises.StatusDone:
		return doneStyle.Render("✓ done")
	case exercises.StatusLintFail:
		return lintStyle.Render("⚠ lint")
	case exercises.StatusFailing:
		return failStyle.Render("● failing")
	case exercises.StatusPending:
		return passStyle.Render("◐ remove marker")
	default:
		return dimStyle.Render("○ pending")
	}
}

// detail is the scrollable content of the right pane.
func (m Model) detail() string {
	if m.showHint {
		return "HINT\n\n" + m.current().Hint
	}
	if !m.hasResult {
		return dimStyle.Render("Press ⏎ to run this exercise (or just edit the file and save).")
	}
	switch m.status {
	case exercises.StatusFailing:
		return m.result.Err + "\n" + m.result.Out
	case exercises.StatusPending:
		return passStyle.Render("✓ Compiles and tests pass!") +
			"\n\nRemove the '// I AM NOT DONE' marker to complete this exercise.\n\n" +
			m.result.Out
	case exercises.StatusLintFail:
		return "Tests pass, but golangci-lint found issues:\n\n" + m.result.Out
	case exercises.StatusDone:
		return doneStyle.Render("✓ Passed and lint-clean!") + "\n\n" + m.result.Out
	default:
		return dimStyle.Render("Edit the exercise, then save to verify.")
	}
}

func (m Model) footer() string {
	return m.help.View(m.keys)
}

// cursorLine maps the cursor (index into items) to its rendered line number.
func (m Model) cursorLine() int { return m.cursor }

// windowLines returns at most height lines centered on the focus line.
func windowLines(lines []string, focus, height int) string {
	if len(lines) <= height {
		return strings.Join(lines, "\n")
	}
	start := focus - height/2
	if start < 0 {
		start = 0
	}
	if start+height > len(lines) {
		start = len(lines) - height
	}
	return strings.Join(lines[start:start+height], "\n")
}

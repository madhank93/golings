package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mauricioabreu/golings/golings/exercises"
)

// item is one row in the left pane: either a topic header or an exercise.
type item struct {
	isHeader    bool
	topic       string
	done, total int // header only
	ex          exercises.Exercise
}

// Model is the Bubble Tea model for the golings TUI.
type Model struct {
	infoFile string
	tracker  *exercises.Tracker
	watchCh  chan string

	items  []item
	cursor int // index into items; always points at an exercise row

	// detail of the currently selected exercise
	status    exercises.Status
	result    exercises.Result
	verifying bool
	hasResult bool
	showHint  bool

	keys     keyMap
	help     help.Model
	progress progress.Model
	spinner  spinner.Model
	output   viewport.Model

	total  int
	width  int
	height int

	notice string // transient message (e.g. after reset)
}

// New builds the model, loads exercises and saved progress, and starts the
// file watcher goroutine.
func New(infoFile string) (Model, error) {
	exs, err := exercises.List(infoFile)
	if err != nil {
		return Model{}, err
	}
	tracker, err := exercises.LoadState(exercises.StateFile)
	if err != nil {
		return Model{}, err
	}

	ch := make(chan string)
	go startWatcher(ch)

	sp := spinner.New()
	sp.Spinner = spinner.Dot

	m := Model{
		infoFile: infoFile,
		tracker:  tracker,
		watchCh:  ch,
		keys:     defaultKeys(),
		help:     help.New(),
		progress: progress.New(progress.WithDefaultGradient()),
		spinner:  sp,
		output:   viewport.New(0, 0),
		total:    len(exs),
	}
	m.items = buildItems(exs, tracker)
	m.cursor = m.firstSelectable()
	m.moveToFirstPending()
	return m, nil
}

// Init starts the watcher listener, the spinner, and verifies the current item.
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		waitForChange(m.watchCh),
		m.spinner.Tick,
		verifyCmd(m.current()),
	)
}

// buildItems flattens exercises into topic headers + exercise rows.
func buildItems(exs []exercises.Exercise, tracker *exercises.Tracker) []item {
	var items []item
	lastTopic := ""
	var headerIdx int
	for _, ex := range exs {
		topic := topicOf(ex.Path)
		if topic != lastTopic {
			items = append(items, item{isHeader: true, topic: topic})
			headerIdx = len(items) - 1
			lastTopic = topic
		}
		items = append(items, item{ex: ex})
		items[headerIdx].total++
		if tracker.IsDone(ex.Name) {
			items[headerIdx].done++
		}
	}
	return items
}

// current returns the exercise under the cursor.
func (m Model) current() exercises.Exercise {
	if m.cursor >= 0 && m.cursor < len(m.items) && !m.items[m.cursor].isHeader {
		return m.items[m.cursor].ex
	}
	return exercises.Exercise{}
}

func (m Model) firstSelectable() int {
	for i, it := range m.items {
		if !it.isHeader {
			return i
		}
	}
	return 0
}

// moveToFirstPending positions the cursor on the first not-yet-done exercise.
func (m *Model) moveToFirstPending() {
	for i, it := range m.items {
		if it.isHeader {
			continue
		}
		if !m.tracker.IsDone(it.ex.Name) {
			m.cursor = i
			return
		}
	}
}

// moveCursor steps to the next/previous exercise row, skipping headers.
func (m *Model) moveCursor(delta int) {
	i := m.cursor
	for {
		i += delta
		if i < 0 || i >= len(m.items) {
			return // out of range; keep current
		}
		if !m.items[i].isHeader {
			m.cursor = i
			return
		}
	}
}

// advance moves to the next not-done exercise after the cursor.
func (m *Model) advance() {
	for i := m.cursor + 1; i < len(m.items); i++ {
		if m.items[i].isHeader {
			continue
		}
		if !m.tracker.IsDone(m.items[i].ex.Name) {
			m.cursor = i
			return
		}
	}
}

// refreshHeaderCounts recomputes per-topic done counts from the tracker.
func (m *Model) refreshHeaderCounts() {
	var hdr int
	for i := range m.items {
		if m.items[i].isHeader {
			m.items[i].done = 0
			hdr = i
			continue
		}
		if m.tracker.IsDone(m.items[i].ex.Name) {
			m.items[hdr].done++
		}
	}
}

// verifiedMsg carries the result of verifying an exercise.
type verifiedMsg struct {
	name   string
	status exercises.Status
	result exercises.Result
}

// verifyCmd runs the gated verification off the UI thread.
func verifyCmd(e exercises.Exercise) tea.Cmd {
	if e.Name == "" {
		return nil
	}
	return func() tea.Msg {
		st, res := e.Verify()
		return verifiedMsg{name: e.Name, status: st, result: res}
	}
}

func now() time.Time { return time.Now() }

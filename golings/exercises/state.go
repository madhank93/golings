package exercises

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"time"
)

// StateFile is the default progress file name (kept beside info.toml).
const StateFile = ".golings-state.json"

const dayLayout = "2006-01-02"

// Streak tracks consecutive days on which at least one exercise was completed.
type Streak struct {
	Count   int    `json:"count"`
	LastDay string `json:"lastDay"` // YYYY-MM-DD, "" if none yet
}

// Tracker is persisted metadata about progress. The file markers ("I AM NOT
// DONE") remain the source of truth for done-ness; this only records when each
// exercise was first completed, the streak, and the last position.
type Tracker struct {
	// Completed maps exercise name -> RFC3339 timestamp of first completion.
	Completed map[string]string `json:"completed"`
	Streak    Streak            `json:"streak"`
	Current   string            `json:"current"`

	path string
}

// LoadState reads the state file. A missing file yields an empty, ready State.
func LoadState(path string) (*Tracker, error) {
	s := &Tracker{Completed: map[string]string{}, path: path}

	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return s, err
	}
	if err := json.Unmarshal(data, s); err != nil {
		return s, err
	}
	if s.Completed == nil {
		s.Completed = map[string]string{}
	}
	s.path = path
	return s, nil
}

// Save writes the state file (pretty-printed).
func (s *Tracker) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, data, 0o644)
}

// IsDone reports whether the exercise has a recorded completion.
func (s *Tracker) IsDone(name string) bool {
	_, ok := s.Completed[name]
	return ok
}

// DoneCount returns how many exercises have been completed.
func (s *Tracker) DoneCount() int { return len(s.Completed) }

// MarkDone records a first-time completion at now and updates the streak.
// Re-marking an already-completed exercise is a no-op.
func (s *Tracker) MarkDone(name string, now time.Time) {
	if _, ok := s.Completed[name]; ok {
		return
	}
	s.Completed[name] = now.Format(time.RFC3339)
	s.bumpStreak(now)
}

// bumpStreak advances the streak on the first completion of a given day.
func (s *Tracker) bumpStreak(now time.Time) {
	today := now.Format(dayLayout)
	switch s.Streak.LastDay {
	case today:
		// already counted a completion today
	case now.AddDate(0, 0, -1).Format(dayLayout):
		s.Streak.Count++
	default:
		s.Streak.Count = 1
	}
	s.Streak.LastDay = today
}

// Reconcile backfills completions for exercises that are already Done on disk
// but have no recorded timestamp (e.g. finished before the state file existed).
func (s *Tracker) Reconcile(doneNames []string, now time.Time) {
	for _, name := range doneNames {
		s.MarkDone(name, now)
	}
}

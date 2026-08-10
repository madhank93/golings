package exercises

import (
	"fmt"
	"os/exec"
)

// Reset restores an exercise to its committed (shipped, broken) state with git:
// the single exercise file, or the whole directory for a package-mode exercise.
// Requires the repo to be a git checkout and the target to be committed.
func Reset(e Exercise) error {
	target := Target(e)
	cmd := exec.Command("git", "checkout", "HEAD", "--", target)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("git checkout failed for %s: %w: %s", target, err, out)
	}
	return nil
}

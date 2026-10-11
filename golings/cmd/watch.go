package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/madhank93/golings/golings/tui"
	"github.com/spf13/cobra"
)

func WatchCmd(infoFile string) *cobra.Command {
	var start string
	cmd := &cobra.Command{
		Use:           "watch",
		Short:         "Interactive TUI: verify exercises as you edit",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := tui.New(infoFile)
			if err != nil {
				return err
			}
			defer m.Close()
			if start != "" {
				if err := m.MoveTo(start); err != nil {
					color.Red("No exercise found for '%s'", start)
					return err
				}
			}
			p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
			_, err = p.Run()
			return err
		},
	}
	cmd.Flags().StringVarP(&start, "exercise", "e", "", "open on this exercise instead of the first pending one")
	return cmd
}

package cli

import (
	"github.com/spf13/cobra"
)

func (c cli) PlayCommand() *cobra.Command {
	var playCommand = &cobra.Command{
		Use:   "play",
		Short: "For playing a tonal structure.",
		Args:  cobra.ExactArgs(0),
	}

	subCommands := []*cobra.Command{
		c.ToneCommand(),
		c.ScaleCommand(),
		c.ChordCommand(),
		c.ProgressionCommand(),
	}
	playCommand.AddCommand(subCommands...)

	return playCommand
}

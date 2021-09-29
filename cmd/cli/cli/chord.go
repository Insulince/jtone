package cli

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Insulince/jtone/pkg/chord"
)

func (c cli) ChordCommand() *cobra.Command {
	const (
		defaultDuration = 1000
	)

	var playChordCommand = &cobra.Command{
		Use:   "chord",
		Short: "Plays provided chord.",
		Args:  cobra.ExactArgs(1),
		RunE:  c.RunChordCommand,
	}

	playChordCommand.Flags().IntP("duration", "d", defaultDuration, fmt.Sprintf("the duration to play the chord for in milliseconds (defaults to %v)", defaultDuration))

	return playChordCommand
}

func (c cli) RunChordCommand(cmd *cobra.Command, args []string) error {
	_ = cmd.Context()

	blueprint := args[0]

	ms, err := cmd.Flags().GetInt("duration")
	if err != nil {
		return errors.Wrap(err, "getting duration flag")
	}
	duration := time.Duration(ms) * time.Millisecond

	ch, err := chord.From(blueprint)
	if err != nil {
		return errors.Wrap(err, "parsing unrecognized chord")
	}

	<-ch.Play(sr, duration)

	return nil
}

package cli

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Insulince/jtone/pkg/progression"
)

func (c cli) ProgressionCommand() *cobra.Command {
	const (
		defaultDuration = 1000
	)

	var playProgressionCommand = &cobra.Command{
		Use:   "progression",
		Short: "Plays each chord in provided progression.",
		Args:  cobra.ExactArgs(1),
		RunE:  c.RunProgressionCommand,
	}

	playProgressionCommand.Flags().IntP("duration", "d", defaultDuration, fmt.Sprintf("the duration to play each chord for in milliseconds (defaults to %v)", defaultDuration))

	return playProgressionCommand
}

func (c cli) RunProgressionCommand(cmd *cobra.Command, args []string) error {
	_ = cmd.Context()

	blueprint := args[0]

	ms, err := cmd.Flags().GetInt("duration")
	if err != nil {
		return errors.Wrap(err, "getting duration flag")
	}
	duration := time.Duration(ms) * time.Millisecond

	p, err := progression.From(blueprint)
	if err != nil {
		return errors.Wrap(err, "new progression")
	}

	<-p.Play(sr, duration)

	return nil
}

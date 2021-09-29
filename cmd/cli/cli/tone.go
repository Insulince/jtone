package cli

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Insulince/jtone/pkg/tone"
)

func (c cli) ToneCommand() *cobra.Command {
	const (
		defaultDuration = 1000
	)

	var playToneCommand = &cobra.Command{
		Use:   "tone",
		Short: "Plays provided tone.",
		Args:  cobra.ExactArgs(1),
		RunE:  c.RunToneCommand,
	}

	playToneCommand.Flags().IntP("duration", "d", defaultDuration, fmt.Sprintf("the duration to play the tone for in milliseconds (defaults to %v)", defaultDuration))

	return playToneCommand
}

func (c cli) RunToneCommand(cmd *cobra.Command, args []string) error {
	_ = cmd.Context()

	blueprint := args[0]

	ms, err := cmd.Flags().GetInt("duration")
	if err != nil {
		return errors.Wrap(err, "getting duration flag")
	}
	duration := time.Duration(ms) * time.Millisecond

	tn, err := tone.From(blueprint)
	if err != nil {
		return errors.Wrap(err, "tone from")
	}

	<-tn.Play(sr, duration)

	return nil
}

package cli

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/Insulince/jtone/pkg/scale"
)

func (c cli) ScaleCommand() *cobra.Command {
	const (
		defaultDuration = 250
	)

	var playScaleCommand = &cobra.Command{
		Use:   "scale",
		Short: "Plays each tone in provided scale.",
		Args:  cobra.ExactArgs(1),
		RunE:  c.RunScaleCommand,
	}

	playScaleCommand.Flags().IntP("duration", "d", defaultDuration, fmt.Sprintf("the duration to play each note of the scale in milliseconds (defaults to %v)", defaultDuration))
	playScaleCommand.Flags().BoolP("reverse", "r", false, "should play the scale in reverse")

	return playScaleCommand
}

func (c cli) RunScaleCommand(cmd *cobra.Command, args []string) error {
	_ = cmd.Context()

	blueprint := args[0]

	ms, err := cmd.Flags().GetInt("duration")
	if err != nil {
		return errors.Wrap(err, "getting duration flag")
	}
	duration := time.Duration(ms) * time.Millisecond

	reverse, err := cmd.Flags().GetBool("reverse")
	if err != nil {
		return errors.Wrap(err, "getting reverse flag")
	}

	s, err := scale.From(blueprint)
	if err != nil {
		return errors.Wrap(err, "parsing unrecognized scale")
	}

	if !reverse {
		<-s.Play(sr, duration)
	} else {
		<-s.PlayReverse(sr, duration)
	}

	return nil
}

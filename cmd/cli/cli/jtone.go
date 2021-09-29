package cli

import (
	"fmt"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	bufferDuration   = 100 * time.Millisecond
	samplesPerSecond = 44100 // 44.1kHz
)

var (
	sr = beep.SampleRate(samplesPerSecond)
)

type (
	cli struct {
	}
)

func New() *cobra.Command {
	var c cli

	coreCommand := cobra.Command{
		Use:     "jtone",
		Short:   "core cli command",
		Args:    cobra.ExactArgs(0),
		Version: fmt.Sprintf("Version: %s.%s.%s", "1", "0", "0"),
	}

	subCommands := []*cobra.Command{
		c.PlayCommand(),
	}
	coreCommand.AddCommand(subCommands...)

	if err := speaker.Init(sr, sr.N(bufferDuration)); err != nil {
		panic(errors.Wrap(err, "speaker init"))
	}

	return &coreCommand
}

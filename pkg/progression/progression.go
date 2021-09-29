package progression

import (
	"fmt"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"

	"github.com/Insulince/jtone/pkg/chord"
	"github.com/Insulince/jtone/pkg/util"
)

type (
	Progression struct {
		chords []chord.Chord
	}
)

var (
	_ fmt.Stringer = Progression{}
)

func New(chs []chord.Chord) Progression {
	var p Progression

	p.chords = chs

	return p
}

func From(blueprint string) (Progression, error) {
	blueprint = strings.ReplaceAll(blueprint, " ", "")
	if blueprint == "" {
		return Progression{}, errors.Errorf("cannot parse chord, blank")
	}

	var chs []chord.Chord
	pieces := strings.Split(blueprint, ";")
	for _, piece := range pieces {
		ch, err := chord.From(piece)
		if err != nil {
			return Progression{}, errors.Wrap(err, "chord from")
		}
		chs = append(chs, ch)
	}

	p := New(chs)
	return p, nil
}

func MustFrom(blueprint string) Progression {
	p, err := From(blueprint)
	if err != nil {
		panic(errors.Wrap(err, "must from"))
	}
	return p
}

func (p Progression) String() string {
	panic("implement me")
}

func (p Progression) Dump() string {
	var out string
	for _, ch := range p.chords {
		out += ch.String() + "\n"
	}
	return out
}

func (p Progression) Streamer(sr beep.SampleRate, duration time.Duration) beep.Streamer {
	var chordStreamers []beep.Streamer
	for _, ch := range p.chords {
		chordStreamers = append(chordStreamers, ch.Streamer(sr, duration))
	}
	streamer := beep.Seq(chordStreamers...)
	return streamer
}

func (p Progression) Play(sr beep.SampleRate, duration time.Duration) chan struct{} {
	fmt.Println(p.Dump())
	done := make(chan struct{})
	var streamers []beep.Streamer
	streamers = append(streamers, p.Streamer(sr, duration))
	streamers = append(streamers, beep.Silence(sr.N(util.EndingSilenceDuration)))
	streamers = append(streamers, beep.Callback(func() { close(done) }))
	speaker.Play(beep.Seq(streamers...))
	return done
}

func (p Progression) Chords() []chord.Chord {
	return p.chords
}

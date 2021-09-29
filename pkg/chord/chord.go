package chord

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"

	"github.com/Insulince/jtone/pkg/degree"
	"github.com/Insulince/jtone/pkg/tone"
	"github.com/Insulince/jtone/pkg/util"
)

type (
	Chord struct {
		root  tone.Tone
		tones []tone.Tone
	}
)

var (
	_ fmt.Stringer = Chord{}
)

func New(tones []tone.Tone) Chord {
	var ch Chord

	if len(tones) == 0 {
		return ch
	}

	ch.tones = tones

	root := tones[0]
	for _, tn := range tones {
		if tn.Cmp(root) == -1 {
			root = tn
		}
	}
	ch.root = root

	return ch
}

func From(blueprint string) (Chord, error) {
	blueprint = strings.ReplaceAll(blueprint, " ", "")
	if blueprint == "" {
		return Chord{}, errors.Errorf("cannot parse chord, blank")
	}

	pieces := strings.Split(blueprint, ":")
	if len(pieces) == 1 {
		identifier := pieces[0]

		if ch, found := NameToChord[identifier]; found {
			return ch, nil
		}
		if ch, found := NameLongToChord[identifier]; found {
			return ch, nil
		}
		if ch, found := NameShortToChord[identifier]; found {
			return ch, nil
		}

		var tns []tone.Tone
		pieces := strings.Split(identifier, ",")
		for _, piece := range pieces {
			tn, err := tone.From(piece)
			if err != nil {
				return Chord{}, errors.Wrap(err, "tone from")
			}
			tns = append(tns, tn)
		}

		ch := New(tns)
		return ch, nil
	} else if len(pieces) == 2 {
		key := pieces[0]
		identifier := pieces[1]

		tn, err := tone.From(key)
		if err != nil {
			return Chord{}, errors.Wrap(err, "tone from")
		}

		ch, err := From(identifier)
		if err != nil {
			return Chord{}, errors.Wrap(err, "chord from")
		}

		ch = ch.Transpose(tn)

		return ch, nil
	} else {
		return Chord{}, errors.Errorf("unrecognized chord blueprint, too many pieces")
	}
}

func MustFrom(blueprint string) Chord {
	c, err := From(blueprint)
	if err != nil {
		panic(errors.Wrap(err, "must from"))
	}
	return c
}

func (ch Chord) String() string {
	degreeList := degree.Stringify(ch.Degrees())
	nameShort, found := DegreesToNameShort[degreeList]
	if !found {
		return fmt.Sprintf("%s:%s", ch.root, degreeList)
	}
	return fmt.Sprintf("%s:%s", ch.root, nameShort)
}

func (ch Chord) Dump() string {
	degreeList := degree.Stringify(ch.Degrees())
	toneList := tone.Stringify(ch.tones)

	return fmt.Sprintf("%s\nROOT: %s\nDEGREES: %s\nTONES: %s", ch.String(), ch.root.Dump(), degreeList, toneList)
}

func (ch Chord) Streamer(sr beep.SampleRate, duration time.Duration) beep.Streamer {
	var toneStreamers []beep.Streamer
	for _, t := range ch.tones {
		toneStreamers = append(toneStreamers, t.Streamer(sr, duration))
	}
	streamer := beep.Mix(toneStreamers...)
	return streamer
}

func (ch Chord) Play(sr beep.SampleRate, duration time.Duration) chan struct{} {
	fmt.Println(ch.Dump())
	done := make(chan struct{})
	var streamers []beep.Streamer
	streamers = append(streamers, ch.Streamer(sr, duration))
	streamers = append(streamers, beep.Silence(sr.N(util.EndingSilenceDuration)))
	streamers = append(streamers, beep.Callback(func() { close(done) }))
	speaker.Play(beep.Seq(streamers...))
	return done
}

func (ch Chord) Degrees() []degree.Degree {
	if len(ch.tones) == 0 {
		return nil
	}

	root := ch.tones[0]
	minOctave := math.MaxInt64
	for _, tn := range ch.tones {
		if tn.Cmp(root) == -1 {
			root = tn
		}

		if tn.Octave() < minOctave {
			minOctave = tn.Octave()
		}
	}

	rootNote := root.Note()
	var ds []degree.Degree
	for _, tn := range ch.tones {
		noteValue := int(tn.Note() - rootNote)
		relativeOctave := tn.Octave() - minOctave
		octaveValue := relativeOctave * degree.NumDegrees
		n := noteValue + octaveValue

		d := degree.Degree(n)
		ds = append(ds, d)
	}

	return ds
}

func (ch Chord) Root() tone.Tone {
	return ch.root
}

func (ch Chord) Tones() []tone.Tone {
	return ch.tones
}

func (ch Chord) Transpose(t tone.Tone) Chord {
	semitones := ch.root.Dist(t)
	ch.root = t

	var tns []tone.Tone
	for _, tn := range ch.tones {
		nt := tn.Shift(semitones)
		tns = append(tns, nt)
	}
	ch.tones = tns

	return ch
}

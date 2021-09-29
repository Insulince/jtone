package scale

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
	Scale struct {
		root  tone.Tone
		tones []tone.Tone
	}
)

var (
	_ fmt.Stringer = Scale{}
)

func New(tones []tone.Tone) Scale {
	var s Scale

	if len(tones) == 0 {
		return s
	}

	s.tones = tones

	root := tones[0]
	for _, tn := range tones {
		if tn.Cmp(root) == -1 {
			root = tn
		}
	}
	s.root = root

	return s
}

func From(blueprint string) (Scale, error) {
	blueprint = strings.ReplaceAll(blueprint, " ", "")
	if blueprint == "" {
		return Scale{}, errors.Errorf("cannot parse scale, blank")
	}

	pieces := strings.Split(blueprint, ":")
	if len(pieces) == 1 {
		identifier := pieces[0]

		if s, found := NameToScale[identifier]; found {
			return s, nil
		}

		var tns []tone.Tone
		pieces := strings.Split(identifier, ",")
		for _, piece := range pieces {
			tn, err := tone.From(piece)
			if err != nil {
				return Scale{}, errors.Wrap(err, "tone from")
			}
			tns = append(tns, tn)
		}

		s := New(tns)
		return s, nil
	} else if len(pieces) == 2 {
		key := pieces[0]
		identifier := pieces[1]

		tn, err := tone.From(key)
		if err != nil {
			return Scale{}, errors.Wrap(err, "tone from")
		}

		s, err := From(identifier)
		if err != nil {
			return Scale{}, errors.Wrap(err, "scale from")
		}

		s = s.Transpose(tn)

		return s, nil
	} else {
		return Scale{}, errors.Errorf("unrecognized scale blueprint, too many pieces")
	}
}

func MustFrom(blueprint string) Scale {
	s, err := From(blueprint)
	if err != nil {
		panic(errors.Wrap(err, "must from"))
	}
	return s
}

func (s Scale) String() string {
	degrees := degree.Stringify(s.Degrees())
	nameShort, found := DegreesToName[degrees]
	if !found {
		return fmt.Sprintf("%s:%s", s.root, degrees)
	}
	return fmt.Sprintf("%s:%s", s.root, nameShort)
}

func (s Scale) Dump() string {
	degreeList := degree.Stringify(s.Degrees())
	toneList := tone.Stringify(s.tones)

	return fmt.Sprintf("%s\nROOT: %s\nDEGREES: %s\nTONES: %s", s.String(), s.root.Dump(), degreeList, toneList)
}

func (s Scale) Streamer(sr beep.SampleRate, duration time.Duration) beep.Streamer {
	var toneStreamers []beep.Streamer
	for _, t := range s.tones {
		toneStreamers = append(toneStreamers, t.Streamer(sr, duration))
	}
	streamer := beep.Seq(toneStreamers...)
	return streamer
}

func (s Scale) ReverseStreamer(sr beep.SampleRate, duration time.Duration) beep.Streamer {
	var toneStreamers []beep.Streamer
	for i := len(s.tones) - 1; i >= 0; i-- {
		t := s.tones[i]
		toneStreamers = append(toneStreamers, t.Streamer(sr, duration))
	}
	streamer := beep.Seq(toneStreamers...)
	return streamer
}

func (s Scale) Play(sr beep.SampleRate, duration time.Duration) chan struct{} {
	fmt.Println(s.Dump())
	done := make(chan struct{})
	var streamers []beep.Streamer
	streamers = append(streamers, s.Streamer(sr, duration))
	streamers = append(streamers, beep.Silence(sr.N(util.EndingSilenceDuration)))
	streamers = append(streamers, beep.Callback(func() { close(done) }))
	speaker.Play(beep.Seq(streamers...))
	return done
}

func (s Scale) PlayReverse(sr beep.SampleRate, duration time.Duration) chan struct{} {
	fmt.Println(s.Dump())
	done := make(chan struct{})
	var streamers []beep.Streamer
	streamers = append(streamers, s.ReverseStreamer(sr, duration))
	streamers = append(streamers, beep.Silence(sr.N(util.EndingSilenceDuration)))
	streamers = append(streamers, beep.Callback(func() { close(done) }))
	speaker.Play(beep.Seq(streamers...))
	return done
}

func (s Scale) Degrees() []degree.Degree {
	if len(s.tones) == 0 {
		return nil
	}

	root := s.tones[0]
	minOctave := math.MaxInt64
	for _, tn := range s.tones {
		if tn.Cmp(root) == -1 {
			root = tn
		}

		if tn.Octave() < minOctave {
			minOctave = tn.Octave()
		}
	}

	rootNote := root.Note()
	var ds []degree.Degree
	for _, tn := range s.tones {
		noteValue := int(tn.Note() - rootNote)
		relativeOctave := tn.Octave() - minOctave
		octaveValue := relativeOctave * degree.NumDegrees
		n := noteValue + octaveValue

		d := degree.Degree(n)
		ds = append(ds, d)
	}

	return ds
}

func (s Scale) Root() tone.Tone {
	return s.root
}

func (s Scale) Tones() []tone.Tone {
	return s.tones
}

func (s Scale) Transpose(t tone.Tone) Scale {
	semitones := s.root.Dist(t)
	s.root = t

	var tns []tone.Tone
	for _, tn := range s.tones {
		nt := tn.Shift(semitones)
		tns = append(tns, nt)
	}
	s.tones = tns

	return s
}

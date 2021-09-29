package tone

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/pkg/errors"

	"github.com/Insulince/jtone/pkg/degree"
	"github.com/Insulince/jtone/pkg/note"
	"github.com/Insulince/jtone/pkg/util"
)

const (
	// HzA4 is the hz value for the A4 note. A4 is the anchor note for this
	// code, as it is a very common note which has a terminating decimal value.
	// All other notes are calculated relative to this value, hence "anchor"
	// note.
	HzA4        = 440.0
	A4Semitones = 49 // A4 is th 49th key on a standard piano.
	AnchorNote  = note.A
)

type (
	Tone struct {
		note   note.Note
		octave int
		degree degree.Degree
	}
)

var (
	_ fmt.Stringer = Tone{}
)

var (
	toneRegex = regexp.MustCompile("")
)

func New(n note.Note, octave int, d degree.Degree) Tone {
	var t Tone

	t.note = n
	t.octave = octave
	t.degree = d

	return t
}

func From(blueprint string) (Tone, error) {
	blueprint = strings.ReplaceAll(blueprint, " ", "")
	if blueprint == "" {
		return C5, nil
	}

	pieces := strings.Split(blueprint, "#")
	if len(pieces) == 1 {
		piece := pieces[0]

		if degree.IsDegreeNotation(piece) {
			d, err := degree.Parse(piece)
			if err != nil {
				return Tone{}, errors.Wrap(err, "parsing degree")
			}

			t := New(note.C, 5, degree.D1)
			t = t.Shift(int(d))

			return t, nil
		} else if IsToneNotation(piece) {
			n, octave, err := parseName(piece)
			if err != nil {
				return Tone{}, errors.Wrap(err, "parsing tone name")
			}

			t := New(n, octave, degree.D1)

			return t, nil
		} else {
			return Tone{}, errors.Errorf("unrecognized tone format")
		}
	} else if len(pieces) == 2 {
		name := pieces[0]
		deg := pieces[1]

		t, err := From(name)
		if err != nil {
			return Tone{}, errors.Wrap(err, "from root")
		}

		d, err := degree.Parse(deg)
		if err != nil {
			return Tone{}, errors.Wrap(err, "parsing degree")
		}

		t.degree = d

		return t, nil
	} else {
		return Tone{}, errors.Errorf("unrecognized tone blueprint, too many pieces")
	}
}

func MustFrom(blueprint string) Tone {
	t, err := From(blueprint)
	if err != nil {
		panic("must from")
	}
	return t
}

func FromList(blueprintList string) ([]Tone, error) {
	blueprintList = strings.ReplaceAll(blueprintList, " ", "")
	if blueprintList == "" {
		return nil, nil
	}

	var tns []Tone

	pieces := strings.Split(blueprintList, ",")
	for _, piece := range pieces {
		tn, err := From(piece)
		if err != nil {
			return nil, errors.Wrap(err, "tone from")
		}
		tns = append(tns, tn)
	}

	return tns, nil
}

func MustFromList(blueprintList string) []Tone {
	tns, err := FromList(blueprintList)
	if err != nil {
		panic(errors.Wrap(err, "must from list"))
	}
	return tns
}

func (t Tone) Transpose(d degree.Degree) {
	t.degree = d
}

func (t Tone) String() string {
	return fmt.Sprintf("%v%v#%s", t.note, t.octave, t.degree)
}

func (t Tone) Dump() string {
	return fmt.Sprintf("%s\nHZ: %.2fhz", t, t.Hz())
}

func Stringify(ts []Tone) string {
	var tones []string
	for _, d := range ts {
		tones = append(tones, fmt.Sprintf("%v", d))
	}
	toneList := strings.Join(tones, ",")
	return toneList
}

func (t Tone) Note() note.Note {
	st := t.Shift(int(t.degree)) // Shift the tone according to its degree.
	return st.note
}

func (t Tone) Octave() int {
	st := t.Shift(int(t.degree)) // Shift the tone according to its degree.
	return st.octave
}

func (t Tone) Degree() degree.Degree {
	return t.degree
}

// TODO(justin): Consider moving into its own package? This is a very important function that isn't solely related to tones.
func (t Tone) Sine(sr beep.SampleRate) beep.StreamerFunc {
	var n float64                  // Counts how many samples have been processed. Is modular with respect to wave period.
	period := float64(sr) / t.Hz() // The period of the wave to complete hz waves in 1 second.
	var flag bool
	return func(samples [][2]float64) (int, bool) {
		for i := range samples {
			n = math.Mod(n, period)                   // Take the mod of n with respect to period to keep values of n small.
			progress := n / period                    // How far into this wave we have progressed.
			value := math.Sin(progress * 2 * math.Pi) // The value of the speaker compression for this sample. [-1.0, 1.0]
			samples[i][0] = value
			samples[i][1] = value
			n++ // A new sample has been processed, increment n.
		}
		if !flag {
			flag = true
			fmt.Println(t.Dump())
		}
		return len(samples), true
	}
}

func (t Tone) Streamer(sr beep.SampleRate, duration time.Duration) beep.Streamer {
	streamer := beep.Take(sr.N(duration), t.Sine(sr))
	return streamer
}

func (t Tone) Play(sr beep.SampleRate, duration time.Duration) chan struct{} {
	fmt.Println(t.Dump())
	done := make(chan struct{})
	var streamers []beep.Streamer
	streamers = append(streamers, t.Streamer(sr, duration))
	streamers = append(streamers, beep.Silence(sr.N(util.EndingSilenceDuration)))
	streamers = append(streamers, beep.Callback(func() { close(done) }))
	speaker.Play(beep.Seq(streamers...))
	return done
}

func (t Tone) Hz() float64 {
	// Formula recreated from: https://en.wikipedia.org/wiki/Piano_key_frequencies
	noteDist := int(t.Note() - AnchorNote) // The number of semitones the requested note is from the anchor note (A).
	octaveDist := t.Octave() * 12          // The number of semitones this octave is from 0.
	dist := noteDist + octaveDist + 1      // The number of semitones this note is from key number 1 (A0).
	semitones := dist - A4Semitones        // The total number of semitones the requested tone is from the anchor tone (A4).
	hz := math.Pow(2, float64(semitones)/12.0) * HzA4

	return hz
}

func (t Tone) Shift(semitones int) Tone {
	var newTone = t
	if semitones > 0 {
		for i := 0; i < semitones; i++ {
			newTone = newTone.Up()
		}
	} else {
		for i := 0; i > semitones; i-- {
			newTone = newTone.Down()
		}
	}
	return newTone
}

func (t Tone) Up() Tone {
	n := t.note.Up()

	octave := t.octave
	if t.note == note.B {
		octave = t.octave + 1
	}

	newTone := New(n, octave, t.degree)
	return newTone
}

func (t Tone) Down() Tone {
	n := t.note.Down()

	octave := t.octave
	if t.note == note.C {
		octave = t.octave - 1
	}

	newTone := New(n, octave, t.degree)
	return newTone
}

func (t Tone) Cmp(t2 Tone) int {
	tHz := t.Hz()
	t2Hz := t2.Hz()

	if tHz > t2Hz {
		return 1
	} else if tHz < t2Hz {
		return -1
	} else {
		return 0
	}
}

func (t Tone) Dist(t2 Tone) int {
	noteDist := int(t2.Note() - t.Note())
	octaveDist := t2.Octave() - t.Octave()
	octaveNoteDist := octaveDist * 12

	return noteDist + octaveNoteDist
}

func IsToneNotation(blueprint string) bool {
	return toneRegex.Match([]byte(blueprint))
}

// parseName... A, As, Bb, and B, are assumed to come from the 4th octave if
// not specified. All other notes are assumed to come from the 5th octave if
// not specified. This is simply a convention that reflects commonly played
// notes. I.e. A in the 5th octave is higher pitched than one would usually pick
// when describing a plain "A" note, and C in the 4th octave is lower pitched
// than one would usually pick when describing a plain "C" note.
func parseName(name string) (n note.Note, octave int, _ error) {
	name = strings.Trim(name, " \t\n")

	if name == "" { // ^$
		return 0, 0, errors.New("bad name: blank")
	}

	switch len(name) {
	case 1: // ^.$
		firstChar := string(name[0])
		switch firstChar {
		case note.NameA: // ^A$
			return note.A, 4, nil
		case note.NameB: // ^B$
			return note.B, 4, nil
		case note.NameC: // ^C$
			return note.C, 5, nil
		case note.NameD: // ^D$
			return note.D, 5, nil
		case note.NameE: // ^E$
			return note.E, 5, nil
		case note.NameF: // ^F$
			return note.F, 5, nil
		case note.NameG: // ^G$
			return note.G, 5, nil
		default: // ^?$
			return 0, 0, errors.New("bad note: not valid natural")
		}
	case 2: // ^..$
		firstChar := string(name[0])
		secondChar := string(name[1])
		switch secondChar {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8": // ^.\d$
			octave, _ = strconv.Atoi(secondChar)
			switch firstChar {
			case note.NameA: // ^A\d$
				return note.A, octave, nil
			case note.NameB: // ^B\d$
				return note.B, octave, nil
			case note.NameC: // ^C\d$
				return note.C, octave, nil
			case note.NameD: // ^D\d$
				return note.D, octave, nil
			case note.NameE: // ^E\d$
				return note.E, octave, nil
			case note.NameF: // ^F\d$
				return note.F, octave, nil
			case note.NameG: // ^G\d$
				return note.G, octave, nil
			default: // ^?\d$
				return 0, 0, errors.New("bad note: not valid natural")
			}
		case note.NameSharp: // ^.s$
			switch firstChar {
			case note.NameA: // ^As$
				return note.As, 4, nil
			case note.NameC: // ^Cs$
				return note.Cs, 5, nil
			case note.NameD: // ^Ds$
				return note.Ds, 5, nil
			case note.NameF: // ^Fs$
				return note.Fs, 5, nil
			case note.NameG: // ^Gs$
				return note.Gs, 5, nil
			case note.NameB, note.NameE: // ^(B|E)s$
				return 0, 0, errors.New("no such note: Bs or Es")
			default: // ^?s$
				return 0, 0, errors.New("bad note: not valid sharp")
			}
		case note.NameFlat: // ^.b$
			switch firstChar {
			case note.NameA: // ^Ab$
				return note.Ab, 5, nil
			case note.NameB: // ^Bb$
				return note.Bb, 4, nil
			case note.NameD: // ^Db$
				return note.Db, 5, nil
			case note.NameE: // ^Eb$
				return note.Eb, 5, nil
			case note.NameG: // ^Gb$
				return note.Gb, 5, nil
			case note.NameC, note.NameF: // ^(C|F)b$
				return 0, 0, errors.New("no such note: Cb or Fb")
			default: // ^?b$
				return 0, 0, errors.New("bad letter: not valid flat")
			}
		default: // ^.?$
			return 0, 0, errors.New("bad accidental: not sharp or flat")
		}
	case 3: // ^...$
		firstChar := string(name[0])
		secondChar := string(name[1])
		thirdChar := string(name[2])
		switch thirdChar {
		case "0", "1", "2", "3", "4", "5", "6", "7", "8": // ^..\d$
			octave, _ = strconv.Atoi(thirdChar)
		default: // ^..?$
			return 0, 0, errors.New("bad octave: not 0-8")
		}
		switch secondChar {
		case note.NameSharp: // ^.s\d$
			switch firstChar {
			case note.NameA: // ^As\d$
				return note.As, octave, nil
			case note.NameC: // ^Cs\d$
				return note.Cs, octave, nil
			case note.NameD: // ^Ds\d$
				return note.Ds, octave, nil
			case note.NameF: // ^Fs\d$
				return note.Fs, octave, nil
			case note.NameG: // ^Gs\d$
				return note.Gs, octave, nil
			case note.NameB, note.NameE: // ^(B|E)s\d$
				return 0, 0, errors.New("no such note: Bs or Es")
			default: // ^?s\d$
				return 0, 0, errors.New("bad note: not valid sharp")
			}
		case note.NameFlat: // ^.b\d$
			switch firstChar {
			case note.NameA: // ^Ab\d$
				return note.Ab, octave, nil
			case note.NameB: // ^Bb\d$
				return note.Bb, octave, nil
			case note.NameD: // ^Db\d$
				return note.Db, octave, nil
			case note.NameE: // ^Eb\d$
				return note.Eb, octave, nil
			case note.NameG: // ^Gb\d$
				return note.Gb, octave, nil
			case note.NameC, note.NameF: // ^(C|F)b\d$
				return 0, 0, errors.New("no such note: Cb or Fb")
			default: // ^?b\d$
				return 0, 0, errors.New("bad letter: not valid flat")
			}
		default: // ^.?\d$
			return 0, 0, errors.New("bad accidental: not sharp or flat")
		}
	default: // ^.{4,}$
		return 0, 0, errors.New("bad name: too long")
	}
}

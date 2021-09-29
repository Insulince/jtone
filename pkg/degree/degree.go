package degree

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/Insulince/jtone/pkg/note"
)

const (
	NumDegrees = 12
	NumNotes   = 7
)

type (
	Degree int
)

var (
	degreeRegex = regexp.MustCompile("^(\\d+)([sb])?$")
)

func Parse(degree string) (Degree, error) {
	degree = strings.ReplaceAll(degree, " ", "")
	if degree == "" {
		return 0, errors.Errorf("bad degree, blank")
	}

	matches := degreeRegex.FindAllStringSubmatch(degree, -1)
	if len(matches) != 1 {
		return 0, errors.Errorf("bad degree, unrecognized format: %s", degree)
	}
	match := matches[0]
	if len(match) != 3 {
		// NOTE(justin): This should not be possible, but is included for
		// completeness.
		return 0, errors.Errorf("bad degree, unknown reason: %s", degree)
	}
	n := match[1]
	accidental := match[2]

	v, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "cannot parse degree: %s", degree)
	}
	if v < 1 {
		return 0, errors.Errorf("degrees start at 1: %s", degree)
	}

	// NOTE(justin): Degrees start at one, but our arithmetic does not,
	// hence decrement.
	v--

	var noteAccidentalShift int
	switch accidental {
	case note.NameNatural:
		noteAccidentalShift = 0
	case note.NameFlat:
		// NOTE(justin): To simplify, if we encounter a flat, we convert it
		// to be the equivalent of a sharp.
		v--
		accidental = note.NameSharp
		fallthrough
	case note.NameSharp:
		noteAccidentalShift = 1
	default:
		return 0, errors.Errorf("unrecognized accidental: %s", accidental)
	}

	notes := int(v % NumNotes)
	octaves := int(v / NumNotes)
	octaveNotes := octaves * NumDegrees

	// Catch invalid degrees.
	switch {
	case (notes == 2 || notes == 6) && accidental == note.NameSharp, // Modular with 3s, 7s
		(notes == 0 || notes == 3) && accidental == note.NameFlat, // Modular with 1b, 4b
		degree == "1b": // Special exception for the exact degree value "1b", because v does not behave well with this value.
		return 0, errors.Errorf("no such degree: %s", degree)
	}

	var cumulativeAccidentalShift int
	switch notes {
	case 0:
		cumulativeAccidentalShift = 0
	case 1:
		cumulativeAccidentalShift = 1
	case 2, 3:
		cumulativeAccidentalShift = 2
	case 4:
		cumulativeAccidentalShift = 3
	case 5:
		cumulativeAccidentalShift = 4
	case 6:
		cumulativeAccidentalShift = 5
	default:
		return 0, errors.Errorf("degree value modulo 6 is somehow not within [0,6]: %v", notes)
	}

	value := octaveNotes + notes + noteAccidentalShift + cumulativeAccidentalShift

	d := Degree(value)
	return d, nil
}

func ParseList(degreeList string) ([]Degree, error) {
	degreeList = strings.ReplaceAll(degreeList, " ", "")
	if degreeList == "" {
		return nil, nil
	}

	degrees := strings.Split(degreeList, ",")

	var ds []Degree
	for i, degree := range degrees {
		d, err := Parse(degree)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing degree %v", i)
		}
		ds = append(ds, d)
	}

	return ds, nil
}

// Modularize gets the modular value of d. This is done by taking the remainder
// of d and NumDegrees, and returning the value as a Degree, which will fall
// between [0, NumDegrees).
//
// Modularize ensures that the resulting degree can always be compared with one
// of the modular degrees in this package.
func (d Degree) Modularize() Degree {
	return d % NumDegrees
}

func (d Degree) Name() string {
	name, found := DegreeToName[d]
	if found {
		return name
	}

	value := int(d)

	modularDegree := d.Modularize()

	var noteValue int
	switch modularDegree {
	case D1, D1s:
		noteValue = 0
	case D2, D2s:
		noteValue = 1
	case D3:
		noteValue = 2
	case D4, D4s:
		noteValue = 3
	case D5, D5s:
		noteValue = 4
	case D6, D6s:
		noteValue = 5
	case D7:
		noteValue = 6
	}

	var accidental string
	switch modularDegree {
	case D1, D2, D3, D4, D5, D6, D7:
		accidental = note.NameNatural
	case D1s, D2s, D4s, D5s, D6s:
		accidental = note.NameSharp
	}

	octave := value / NumDegrees
	octaveValue := octave * NumNotes
	n := octaveValue + noteValue

	// NOTE(justin): Arithmetic starts at 0, but degrees start at 1, so
	// increment.
	n++

	name = fmt.Sprintf("%v%s", n, accidental)
	return name
}

func (d Degree) String() string {
	return d.Name()
}

func (d Degree) Dump() string {
	return d.String()
}

func Stringify(ds []Degree) string {
	var degrees []string
	for _, d := range ds {
		degrees = append(degrees, d.String())
	}
	degreeList := strings.Join(degrees, ",")
	return degreeList
}

func IsDegreeNotation(blueprint string) bool {
	return degreeRegex.Match([]byte(blueprint))
}

package note

import (
	"fmt"
)

const (
	Natural = iota
	Sharp
	Flat

	NameNatural = ""
	NameSharp   = "s"
	NameFlat    = "b"
)

type (
	Accidental int
)

var (
	_ fmt.Stringer = Accidental(0)
)

func AccidentalFrom(accidental string) Accidental {
	switch accidental {
	case NameNatural:
		return Natural
	case NameSharp:
		return Sharp
	case NameFlat:
		return Flat
	default:
		panic("unknown accidental")
	}
}

func (a Accidental) String() string {
	switch a {
	case Natural:
		return NameNatural
	case Sharp:
		return NameSharp
	case Flat:
		return NameFlat
	default:
		panic("unknown accidental")
	}
}

func (a Accidental) Up() Accidental {
	switch a {
	case Natural:
		return Sharp
	case Sharp, Flat:
		return Natural
	default:
		panic("unknown accidental")
	}
}

func (a Accidental) Down() Accidental {
	switch a {
	case Natural:
		return Flat
	case Sharp, Flat:
		return Natural
	default:
		panic("unknown accidental")
	}
}

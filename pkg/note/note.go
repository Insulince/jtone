package note

import (
	"fmt"
)

const (
	C  Note = iota
	Cs      // Db
	D
	Ds // Eb
	E
	F
	Fs // Gb
	G
	Gs // Ab
	A
	As // Bb
	B

	Db = Cs
	Eb = Ds
	Gb = Fs
	Ab = Gs
	Bb = As

	NameC  = "C"
	NameCs = NameC + NameSharp
	NameD  = "D"
	NameDs = NameD + NameSharp
	NameE  = "E"
	NameF  = "F"
	NameFs = NameF + NameSharp
	NameG  = "G"
	NameGs = NameG + NameSharp
	NameA  = "A"
	NameAs = NameA + NameSharp
	NameB  = "B"

	NameDb = NameD + NameFlat
	NameEb = NameE + NameFlat
	NameGb = NameG + NameFlat
	NameAb = NameA + NameFlat
	NameBb = NameB + NameFlat
)

type (
	Note int
)

var (
	_ fmt.Stringer = Note(0)
)

func From(name string) Note {
	switch name {
	case NameC:
		return C
	case NameCs:
		return Cs
	case NameDb:
		return Db
	case NameD:
		return D
	case NameDs:
		return Ds
	case NameEb:
		return Eb
	case NameE:
		return E
	case NameF:
		return F
	case NameFs:
		return Fs
	case NameGb:
		return Gb
	case NameG:
		return G
	case NameGs:
		return Gs
	case NameAb:
		return Ab
	case NameA:
		return A
	case NameAs:
		return As
	case NameBb:
		return Bb
	case NameB:
		return B
	default:
		panic("unknown note")
	}
}

func (n Note) String() string {
	switch n {
	case C:
		return NameC
	case Cs:
		return NameCs
	case D:
		return NameD
	case Ds:
		return NameDs
	case E:
		return NameE
	case F:
		return NameF
	case Fs:
		return NameFs
	case G:
		return NameG
	case Gs:
		return NameGs
	case A:
		return NameA
	case As:
		return NameAs
	case B:
		return NameB
	default:
		panic("no such note")
	}
}

func (n Note) Up() Note {
	switch n {
	case C:
		return Cs
	case Cs:
		return D
	case D:
		return Ds
	case Ds:
		return E
	case E:
		return F
	case F:
		return Fs
	case Fs:
		return G
	case G:
		return Gs
	case Gs:
		return A
	case A:
		return As
	case As:
		return B
	case B:
		return C
	default:
		panic("unknown note")
	}
}

func (n Note) Down() Note {
	switch n {
	case C:
		return B
	case Cs:
		return C
	case D:
		return Cs
	case Ds:
		return D
	case E:
		return Ds
	case F:
		return E
	case Fs:
		return F
	case G:
		return Fs
	case Gs:
		return G
	case A:
		return Gs
	case As:
		return A
	case B:
		return As
	default:
		panic("unknown note")
	}
}

func (n Note) Accidental() Accidental {
	switch n {
	case C, D, E, F, G, A, B:
		return Natural
	case Cs, Ds, Fs, Gs, As:
		return Sharp
	default:
		panic("unknown note")
	}
}

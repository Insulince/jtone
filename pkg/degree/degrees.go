package degree

// ModularDegrees
const (
	D1 Degree = iota
	D1s
	D2
	D2s
	D3
	D4
	D4s
	D5
	D5s
	D6
	D6s
	D7

	D2b = D1s
	D3b = D2s
	D5b = D4s
	D6b = D5s
	D7b = D6s
)

// Names
const (
	NameD1  = "1"
	NameD1s = "1s"
	NameD2  = "2"
	NameD2s = "2s"
	NameD3  = "3"
	NameD4  = "4"
	NameD4s = "4s"
	NameD5  = "5"
	NameD5s = "5s"
	NameD6  = "6"
	NameD6s = "6s"
	NameD7  = "7"

	NameD2b = "2b"
	NameD3b = "3b"
	NameD5b = "5b"
	NameD6b = "6b"
	NameD7b = "7b"
)

var (
	NameToDegree = map[string]Degree{
		NameD1:  D1,
		NameD1s: D1s,
		NameD2:  D2,
		NameD2s: D2s,
		NameD3:  D3,
		NameD4:  D4,
		NameD4s: D4s,
		NameD5:  D5,
		NameD5s: D5s,
		NameD6:  D6,
		NameD6s: D6s,
		NameD7:  D7,

		NameD2b: D2b,
		NameD3b: D3b,
		NameD5b: D5b,
		NameD6b: D6b,
		NameD7b: D7b,
	}

	DegreeToName = map[Degree]string{
		D1:  NameD1,
		D1s: NameD1s,
		D2:  NameD2,
		D2s: NameD2s,
		D3:  NameD3,
		D4:  NameD4,
		D4s: NameD4s,
		D5:  NameD5,
		D5s: NameD5s,
		D6:  NameD6,
		D6s: NameD6s,
		D7:  NameD7,
	}
)

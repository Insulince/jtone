package scale

import (
	"testing"

	"github.com/Insulince/jtone/pkg/tone"
)

func Test_t(t *testing.T) {
	toneList := "A,C,Fs6"

	s, err := From(toneList)
	if err != nil {
		panic(err)
	}
	s.Dump()
}

func Test_From(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		toneList := "C1,C0"

		s, err := From(toneList)
		if err != nil {
			panic(err)
		}
		t.Log(s.Dump())
	})
	t.Run("2", func(t *testing.T) {
		toneList := "C4,D4,E4,F4,G3"

		s, err := From(toneList)
		if err != nil {
			panic(err)
		}
		t.Log(s.Dump())
	})
	t.Run("3", func(t *testing.T) {
		toneList := "C5,D5,E5,F5,G3"

		s, err := From(toneList)
		if err != nil {
			panic(err)
		}
		t.Log(s.Dump())
	})
}

func Test_Scale_String(t *testing.T) {
	s := MustFrom("1,3,5")
	s = s.Transpose(tone.MustFrom("C"))
	str := s.String()
	t.Log(str)
}

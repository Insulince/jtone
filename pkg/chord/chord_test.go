package chord

import (
	"testing"

	"github.com/Insulince/jtone/pkg/tone"
)

func Test_Chord_String(t *testing.T) {
	ch := MustFrom("1,3,5")
	ch = ch.Transpose(tone.MustFrom("C"))
	str := ch.String()
	t.Log(str)
}

package tone

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Insulince/jtone/pkg/degree"
	"github.com/Insulince/jtone/pkg/note"
)

func Test_NewTone(t *testing.T) {
	tone, err := From("A4")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, HzA4, tone.Hz())
}

func Test_FromList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		// Tests relative degrees.
		//
		// CHORD:          E4,Fs4,G4
		// TONES:          E4 Fs4 G4
		// NOTES:          E  Fs  G
		// DEGREES:        1  2   2s
		// DEGREE INDEXES: 0  2   3

		toneList := "E,Fs,G"

		tns, err := FromList(toneList)

		require.NoError(t, err)
		require.Len(t, tns, 3)
		assert.Equal(t, tns[0].Note(), note.E)
		assert.Equal(t, tns[0].Octave(), E.Octave())
		assert.Equal(t, tns[1].Note(), note.Fs)
		assert.Equal(t, tns[1].Octave(), Fs.Octave())
		assert.Equal(t, tns[2].Note(), note.G)
		assert.Equal(t, tns[2].Octave(), Gs.Octave())
	})

	t.Run("2", func(t *testing.T) {
		// Tests that the root is not simply the first note found.
		//
		// CHORD:          C1,C0
		// TONES:          C1 C0
		// NOTES:          C  C
		// DEGREES:        8  1
		// DEGREE INDEXES: 12 0

		toneList := "C1,C0"

		tns, err := FromList(toneList)

		require.NoError(t, err)
		require.Len(t, tns, 2)
		assert.Equal(t, tns[0].Note(), note.C)
		assert.Equal(t, tns[0].Octave(), C1.Octave())
		assert.Equal(t, tns[1].Note(), note.C)
		assert.Equal(t, tns[1].Octave(), C0.Octave())
	})
}

func Test_Tone_Dist(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		tn := New(note.C, 5, degree.D1)
		tn2 := New(note.C, 5, degree.D1)
		dist := tn.Dist(tn2)
		t.Log(dist)
	})
	t.Run("1 note", func(t *testing.T) {
		tn := New(note.C, 5, degree.D1)
		tn2 := New(note.Cs, 5, degree.D1)
		dist := tn.Dist(tn2)
		t.Log(dist)
	})
	t.Run("2 notes", func(t *testing.T) {
		tn := New(note.C, 5, degree.D1)
		tn2 := New(note.D, 5, degree.D1)
		dist := tn.Dist(tn2)
		t.Log(dist)
	})
	t.Run("1 octave", func(t *testing.T) {
		tn := New(note.C, 5, degree.D1)
		tn2 := New(note.C, 6, degree.D1)
		dist := tn.Dist(tn2)
		t.Log(dist)
	})
	t.Run("2 octaves", func(t *testing.T) {
		tn := New(note.C, 5, degree.D1)
		tn2 := New(note.C, 7, degree.D1)
		dist := tn.Dist(tn2)
		t.Log(dist)
	})
}

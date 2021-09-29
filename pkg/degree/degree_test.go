package degree

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	t.Run("from a list", func(t *testing.T) {
		ds, err := ParseList("1,1s,2b,2,2s,3b,3,4,4s,5b,5,5s,6b,6,6s,7b,7")

		require.NoError(t, err)
		assert.Len(t, ds, 17)
	})

	t.Run("bad degrees", func(t *testing.T) {
		t.Run("octave 1", func(t *testing.T) {
			t.Run("3s", func(t *testing.T) {
				_, err := ParseList("3s")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("7s", func(t *testing.T) {
				_, err := ParseList("7s")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("1b", func(t *testing.T) {
				_, err := ParseList("1b")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("4b", func(t *testing.T) {
				_, err := ParseList("4b")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
		})
		t.Run("octave 2", func(t *testing.T) {
			t.Run("10s", func(t *testing.T) {
				_, err := ParseList("10s")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("14s", func(t *testing.T) {
				_, err := ParseList("14s")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("8b", func(t *testing.T) {
				_, err := ParseList("8b")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("11b", func(t *testing.T) {
				_, err := ParseList("11b")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
		})
		t.Run("octave 3", func(t *testing.T) {
			t.Run("17s", func(t *testing.T) {
				_, err := ParseList("17s")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("21s", func(t *testing.T) {
				_, err := ParseList("21s")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("15b", func(t *testing.T) {
				_, err := ParseList("15b")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
			t.Run("18b", func(t *testing.T) {
				_, err := ParseList("18b")

				require.Error(t, err)
				assert.Contains(t, err.Error(), "no such degree")
			})
		})
	})

	t.Run("exhaustive", func(t *testing.T) {

		type test struct {
			name string
			d    Degree
		}

		t.Run("automated", func(t *testing.T) {
			// NOTE(justin): Looks confusing? I agree. But this is a way to automate
			// testing every degree name for the first N octaves. Every name as well
			// as expected degree value is calculated. This is bordering on the
			// point where the test code itself needs to be tested, so in the
			// interest of keeping a sanity check, I will also include manual tests
			// for the first 3 octaves, though it truly is redundant if this is
			// working as expected.
			for octave := 0; octave < 10; octave++ {
				t.Run(fmt.Sprintf("octave %v", octave+1), func(t *testing.T) {
					degreeOffset := Degree(NumDegrees * octave)
					nameOffset := octave*7 + 1
					tests := []test{
						{name: fmt.Sprintf("%v", nameOffset+0), d: D1 + degreeOffset},
						{name: fmt.Sprintf("%vs", nameOffset+0), d: D1s + degreeOffset},
						{name: fmt.Sprintf("%vb", nameOffset+1), d: D2b + degreeOffset},
						{name: fmt.Sprintf("%v", nameOffset+1), d: D2 + degreeOffset},
						{name: fmt.Sprintf("%vs", nameOffset+1), d: D2s + degreeOffset},
						{name: fmt.Sprintf("%vb", nameOffset+2), d: D3b + degreeOffset},
						{name: fmt.Sprintf("%v", nameOffset+2), d: D3 + degreeOffset},
						{name: fmt.Sprintf("%v", nameOffset+3), d: D4 + degreeOffset},
						{name: fmt.Sprintf("%vs", nameOffset+3), d: D4s + degreeOffset},
						{name: fmt.Sprintf("%vb", nameOffset+4), d: D5b + degreeOffset},
						{name: fmt.Sprintf("%v", nameOffset+4), d: D5 + degreeOffset},
						{name: fmt.Sprintf("%vs", nameOffset+4), d: D5s + degreeOffset},
						{name: fmt.Sprintf("%vb", nameOffset+5), d: D6b + degreeOffset},
						{name: fmt.Sprintf("%v", nameOffset+5), d: D6 + degreeOffset},
						{name: fmt.Sprintf("%vs", nameOffset+5), d: D6s + degreeOffset},
						{name: fmt.Sprintf("%vb", nameOffset+6), d: D7b + degreeOffset},
						{name: fmt.Sprintf("%v", nameOffset+6), d: D7 + degreeOffset},
					}

					for _, tst := range tests {
						t.Run(tst.name, func(t *testing.T) {
							ds, err := ParseList(tst.name)

							require.NoError(t, err)
							require.Len(t, ds, 1)
							assert.Equal(t, tst.d, ds[0])
						})
					}
				})
			}
		})
		t.Run("manual", func(t *testing.T) {
			t.Run("octave 1", func(t *testing.T) {
				tests := []test{
					{name: "1", d: D1},
					{name: "1s", d: D1s},
					{name: "2b", d: D2b},
					{name: "2", d: D2},
					{name: "2s", d: D2s},
					{name: "3b", d: D3b},
					{name: "3", d: D3},
					{name: "4", d: D4},
					{name: "4s", d: D4s},
					{name: "5b", d: D5b},
					{name: "5", d: D5},
					{name: "5s", d: D5s},
					{name: "6b", d: D6b},
					{name: "6", d: D6},
					{name: "6s", d: D6s},
					{name: "7b", d: D7b},
					{name: "7", d: D7},
				}

				for _, tst := range tests {
					t.Run(tst.name, func(t *testing.T) {
						ds, err := ParseList(tst.name)

						require.NoError(t, err)
						require.Len(t, ds, 1)
						assert.Equal(t, tst.d, ds[0])
					})
				}
			})
			t.Run("octave 2", func(t *testing.T) {
				degreeOffset := Degree(NumDegrees * 1)
				tests := []test{
					{name: "8", d: D1 + degreeOffset},
					{name: "8s", d: D1s + degreeOffset},
					{name: "9b", d: D2b + degreeOffset},
					{name: "9", d: D2 + degreeOffset},
					{name: "9s", d: D2s + degreeOffset},
					{name: "10b", d: D3b + degreeOffset},
					{name: "10", d: D3 + degreeOffset},
					{name: "11", d: D4 + degreeOffset},
					{name: "11s", d: D4s + degreeOffset},
					{name: "12b", d: D5b + degreeOffset},
					{name: "12", d: D5 + degreeOffset},
					{name: "12s", d: D5s + degreeOffset},
					{name: "13b", d: D6b + degreeOffset},
					{name: "13", d: D6 + degreeOffset},
					{name: "13s", d: D6s + degreeOffset},
					{name: "14b", d: D7b + degreeOffset},
					{name: "14", d: D7 + degreeOffset},
				}

				for _, tst := range tests {
					t.Run(tst.name, func(t *testing.T) {
						ds, err := ParseList(tst.name)

						require.NoError(t, err)
						require.Len(t, ds, 1)
						assert.Equal(t, tst.d, ds[0])
					})
				}
			})
			t.Run("octave 3", func(t *testing.T) {
				degreeOffset := Degree(NumDegrees * 2)
				tests := []test{
					{name: "15", d: D1 + degreeOffset},
					{name: "15s", d: D1s + degreeOffset},
					{name: "16b", d: D2b + degreeOffset},
					{name: "16", d: D2 + degreeOffset},
					{name: "16s", d: D2s + degreeOffset},
					{name: "17b", d: D3b + degreeOffset},
					{name: "17", d: D3 + degreeOffset},
					{name: "18", d: D4 + degreeOffset},
					{name: "18s", d: D4s + degreeOffset},
					{name: "19b", d: D5b + degreeOffset},
					{name: "19", d: D5 + degreeOffset},
					{name: "19s", d: D5s + degreeOffset},
					{name: "20b", d: D6b + degreeOffset},
					{name: "20", d: D6 + degreeOffset},
					{name: "20s", d: D6s + degreeOffset},
					{name: "21b", d: D7b + degreeOffset},
					{name: "21", d: D7 + degreeOffset},
				}

				for _, tst := range tests {
					t.Run(tst.name, func(t *testing.T) {
						ds, err := ParseList(tst.name)

						require.NoError(t, err)
						require.Len(t, ds, 1)
						assert.Equal(t, tst.d, ds[0])
					})
				}
			})
		})
	})
}

func Test_Degree_Name(t *testing.T) {
	t.Run("the modular names", func(t *testing.T) {
		t.Run("D1", func(t *testing.T) {
			d := D1
			name := d.Name()
			assert.Equal(t, name, NameD1)
		})
		t.Run("D1s", func(t *testing.T) {
			d := D1s
			name := d.Name()
			assert.Equal(t, name, NameD1s)
		})
		t.Run("D2", func(t *testing.T) {
			d := D2
			name := d.Name()
			assert.Equal(t, name, NameD2)
		})
		t.Run("D2s", func(t *testing.T) {
			d := D2s
			name := d.Name()
			assert.Equal(t, name, NameD2s)
		})
		t.Run("D3", func(t *testing.T) {
			d := D3
			name := d.Name()
			assert.Equal(t, name, NameD3)
		})
		t.Run("D4", func(t *testing.T) {
			d := D4
			name := d.Name()
			assert.Equal(t, name, NameD4)
		})
		t.Run("D4s", func(t *testing.T) {
			d := D4s
			name := d.Name()
			assert.Equal(t, name, NameD4s)
		})
		t.Run("D5", func(t *testing.T) {
			d := D5
			name := d.Name()
			assert.Equal(t, name, NameD5)
		})
		t.Run("D5s", func(t *testing.T) {
			d := D5s
			name := d.Name()
			assert.Equal(t, name, NameD5s)
		})
		t.Run("D6", func(t *testing.T) {
			d := D6
			name := d.Name()
			assert.Equal(t, name, NameD6)
		})
		t.Run("D6s", func(t *testing.T) {
			d := D6s
			name := d.Name()
			assert.Equal(t, name, NameD6s)
		})
		t.Run("D7", func(t *testing.T) {
			d := D7
			name := d.Name()
			assert.Equal(t, name, NameD7)
		})
		t.Run("D2b -> D1s", func(t *testing.T) {
			d := D2b
			name := d.Name()
			assert.Equal(t, name, NameD1s)
		})
		t.Run("D3b -> D2s", func(t *testing.T) {
			d := D3b
			name := d.Name()
			assert.Equal(t, name, NameD2s)
		})
		t.Run("D5b -> D4s", func(t *testing.T) {
			d := D5b
			name := d.Name()
			assert.Equal(t, name, NameD4s)
		})
		t.Run("D6b -> D5s", func(t *testing.T) {
			d := D6b
			name := d.Name()
			assert.Equal(t, name, NameD5s)
		})
		t.Run("D7b -> D6s", func(t *testing.T) {
			d := D7b
			name := d.Name()
			assert.Equal(t, name, NameD6s)
		})
	})
	t.Run("extended names", func(t *testing.T) {
		// Tests that the next 3 octaves above the base can all be mapped to degrees and back.
		degrees := []string{"8", "8s", "9", "9s", "10", "11", "11s", "12", "12s", "13", "13s", "14", "15", "15s", "16", "16s", "17", "18", "18s", "19", "19s", "20", "20s", "21", "22", "22s", "23", "23s", "24", "25", "25s", "26", "26s", "27", "27s", "28"}
		degreeList := strings.Join(degrees, ",")
		ds, err := ParseList(degreeList)
		require.NoError(t, err)

		require.Len(t, ds, len(degrees))
		for i, d := range ds {
			assert.Equal(t, degrees[i], d.Name())
		}
	})
}

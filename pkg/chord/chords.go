package chord

import (
	"github.com/Insulince/jtone/pkg/tone"
)

// TODO(justin): Standardize

const (
	SymbolMajor      = "maj"
	SymbolMinor      = "min"
	SymbolDominant   = "dom"
	SymbolDiminished = "dim"
	SymbolAugmented  = "aug"
	SymbolSuspended  = "sus"
	SymbolAdd        = "add"
)

const (
	// Sourced from https://en.wikipedia.org/wiki/Chord_names_and_symbols_(popular_music)
	// Note that all instances of flat tones were replaces with their sharp
	// counterparts. This is done to allow easier string matching of names. It
	// would be difficult to know that a user played a half diminished ninth
	// chord when the degrees can be represented in multiple ways. Using all
	// sharps fixes this problem, and the preference of sharps over flats is a
	// consistent preference throughout this codebase. Note that names which
	// include a flat are kept as is, since those are simply identifiers and
	// don't inform the tone of its nature in any way.

	MajorFifthDegrees               = "1,3,5"
	MinorFifthDegrees               = "1,2s,5"
	AugmentedFifthDegrees           = "1,3,5s"
	DiminishedFifthDegrees          = "1,2s,4s"
	DominantSeventhDegrees          = "1,3,5,6s"
	MajorSeventhDegrees             = "1,3,5,7"
	MinorMajorSeventhDegrees        = "1,2s,5,7"
	MinorSeventhDegrees             = "1,2s,5,6s"
	AugmentedMajorSeventhDegrees    = "1,3,5s,7"
	AugmentedSeventhDegrees         = "1,3,5s,6s"
	HalfDiminishedSeventhDegrees    = "1,2s,4s,6s"
	DiminishedSeventhDegrees        = "1,2s,4s,6"
	SeventhFlatFiveDegrees          = "1,3,4s,6s"
	MajorNinthDegrees               = "1,3,5,7,9"
	DominantNinthDegrees            = "1,3,5,6s,9"
	DominantMinorNinthDegrees       = "1,3,5,6s,8s"
	MinorMajorNinthDegrees          = "1,2s,5,7,9"
	MinorNinthDegrees               = "1,2s,5,6s,9"
	AugmentedMajorNinthDegrees      = "1,3,5s,7,9"
	AugmentedDominantNinthDegrees   = "1,3,5s,6s,9"
	HalfDiminishedNinthDegrees      = "1,2s,4s,6s,9"
	HalfDiminishedMinorNinthDegrees = "1,2s,4s,6s,8s"
	DiminishedNinthDegrees          = "1,2s,4s,6,9"
	DiminishedMinorNinthDegrees     = "1,2s,4s,6,8s"
	DominantEleventhDegrees         = "1,3,5,6s,9,11"
	MajorEleventhDegrees            = "1,3,5,7,9,11"
	MinorMajorEleventhDegrees       = "1,2s,5,7,9,11"
	MinorEleventhDegrees            = "1,2s,5,6s,9,11"
	AugmentedMajorEleventhDegrees   = "1,3,5s,7,9,11"
	AugmentedEleventhDegrees        = "1,3,5s,6s,9,11"
	HalfDiminishedEleventhDegrees   = "1,2s,4s,6s,9,11"
	DiminishedEleventhDegrees       = "1,2s,4s,6,9,11"
	MajorThirteenthDegrees          = "1,3,5,7,9,11,13"
	ThirteenthDegrees               = "1,3,5,6s,9,11,13"
	MinorMajorThirteenthDegrees     = "1,2s,5,7,9,11,13"
	MinorThirteenthDegrees          = "1,2s,5,6s,9,11,13"
	AugmentedMajorThirteenthDegrees = "1,3,5s,7,9,11,13"
	AugmentedThirteenthDegrees      = "1,3,5s,6s,9,11,13"
	HalfDiminishedThirteenthDegrees = "1,2s,4s,6s,9,11,13"
	MajorTriadDegrees               = MajorFifthDegrees
	MinorTriadDegrees               = MinorFifthDegrees
	AugmentedTriadDegrees           = AugmentedFifthDegrees
	DiminishedTriadDegrees          = DiminishedFifthDegrees

	NameMajorFifth               = "major-fifth"
	NameMinorFifth               = "minor-fifth"
	NameAugmentedFifth           = "augmented-fifth"
	NameDiminishedFifth          = "diminished-fifth"
	NameDominantSeventh          = "dominant-seventh"
	NameMajorSeventh             = "major-seventh"
	NameMinorMajorSeventh        = "minor-major-seventh"
	NameMinorSeventh             = "minor-seventh"
	NameAugmentedMajorSeventh    = "augmented-major-seventh"
	NameAugmentedSeventh         = "augmented-seventh"
	NameHalfDiminishedSeventh    = "half-diminished-seventh"
	NameDiminishedSeventh        = "diminished-seventh"
	NameSeventhFlatFive          = "seventh-flat-five"
	NameMajorNinth               = "major-ninth"
	NameDominantNinth            = "dominant-ninth"
	NameDominantMinorNinth       = "dominant-minor-ninth"
	NameMinorMajorNinth          = "minor-major-ninth"
	NameMinorNinth               = "minor-ninth"
	NameAugmentedMajorNinth      = "augmented-major-ninth"
	NameAugmentedDominantNinth   = "augmented-dominant-ninth"
	NameHalfDiminishedNinth      = "half-diminished-ninth"
	NameHalfDiminishedMinorNinth = "half-diminished-minor-ninth"
	NameDiminishedNinth          = "diminished-ninth"
	NameDiminishedMinorNinth     = "diminished-minor-ninth"
	NameDominantEleventh         = "eleventh"
	NameMajorEleventh            = "major-eleventh"
	NameMinorMajorEleventh       = "minor-major-eleventh"
	NameMinorEleventh            = "minor-eleventh"
	NameAugmentedMajorEleventh   = "augmented-major-eleventh"
	NameAugmentedEleventh        = "augmented-eleventh"
	NameHalfDiminishedEleventh   = "half-diminished-eleventh"
	NameDiminishedEleventh       = "diminished-eleventh"
	NameMajorThirteenth          = "major-thirteenth"
	NameThirteenth               = "thirteenth"
	NameMinorMajorThirteenth     = "minor-major-thirteenth"
	NameMinorThirteenth          = "minor-thirteenth"
	NameAugmentedMajorThirteenth = "augmented-major-thirteenth"
	NameAugmentedThirteenth      = "augmented-thirteenth"
	NameHalfDiminishedThirteenth = "half-diminished-thirteenth"
	NameMajorTriad               = "major"
	NameMinorTriad               = "minor"
	NameAugmentedTriad           = "augmented"
	NameDiminishedTriad          = "diminished"

	NameLongMajorFifth      = "maj5"
	NameLongMinorFifth      = "min5"
	NameLongAugmentedFifth  = "aug5"
	NameLongDiminishedFifth = "dim5"

	NameLongMajorSeventh          = "maj7"
	NameLongMinorSeventh          = "min7"
	NameLongDominantSeventh       = "dom7"
	NameLongDiminishedSeventh     = "dim7"
	NameLongAugmentedSeventh      = "aug7"
	NameLongMinorMajorSeventh     = "minmaj7"
	NameLongAugmentedMajorSeventh = "augmaj7"
	NameLongHalfDiminishedSeventh = "hd7"
	NameLongSeventhFlatFive       = "7dim5" // *

	NameLongMajorNinth               = "maj8"
	NameLongMinorNinth               = "min9"
	NameLongDominantNinth            = "dom9"
	NameLongDiminishedNinth          = "dim9"
	NameLongAugmentedNinth           = "aug9" // * Implement
	NameLongMinorMajorNinth          = "minmaj9"
	NameLongAugmentedMajorNinth      = "augmaj9"
	NameLongHalfDiminishedNinth      = "hd9"
	NameLongDominantMinorNinth       = "79b"     // *
	NameLongAugmentedDominantNinth   = "augdom9" // *
	NameLongHalfDiminishedMinorNinth = "hd9b"    // *
	NameLongDiminishedMinorNinth     = "dim9b"   // *

	NameLongMajorEleventh          = "maj11"
	NameLongMinorEleventh          = "min11"
	NameLongDominantEleventh       = "dom11"
	NameLongDiminishedEleventh     = "dim11"
	NameLongAugmentedEleventh      = "aug11"
	NameLongMinorMajorEleventh     = "minmaj11"
	NameLongAugmentedMajorEleventh = "augmaj11"
	NameLongHalfDiminishedEleventh = "hd11"

	NameLongMajorThirteenth          = "maj13"
	NameLongMinorThirteenth          = "min13"
	NameLongDominantThirteenth       = "dom13"
	NameLongDiminishedThirteenth     = "dim13" // * Implement
	NameLongAugmentedThirteenth      = "aug13"
	NameLongMinorMajorThirteenth     = "minmaj13"
	NameLongAugmentedMajorThirteenth = "augmaj13"
	NameLongHalfDiminishedThirteenth = "hd13"

	NameLongMajorTriad      = "maj"
	NameLongMinorTriad      = "min"
	NameLongAugmentedTriad  = "aug"
	NameLongDiminishedTriad = "dim"

	NameShortMajorFifth               = "M5"
	NameShortMinorFifth               = "m5"
	NameShortAugmentedFifth           = "+5"
	NameShortDiminishedFifth          = "o5"
	NameShortDominantSeventh          = "7"
	NameShortMajorSeventh             = "M7"
	NameShortMinorMajorSeventh        = "mM7"
	NameShortMinorSeventh             = "m7"
	NameShortAugmentedMajorSeventh    = "+M7"
	NameShortAugmentedSeventh         = "+7"
	NameShortHalfDiminishedSeventh    = "07"
	NameShortDiminishedSeventh        = "o7"
	NameShortSeventhFlatFive          = "75b"
	NameShortMajorNinth               = "M9"
	NameShortDominantNinth            = "9"
	NameShortDominantMinorNinth       = "79b"
	NameShortMinorMajorNinth          = "mM9"
	NameShortMinorNinth               = "m9"
	NameShortAugmentedMajorNinth      = "+M9"
	NameShortAugmentedDominantNinth   = "+9"
	NameShortHalfDiminishedNinth      = "09"
	NameShortHalfDiminishedMinorNinth = "09b"
	NameShortDiminishedNinth          = "o9"
	NameShortDiminishedMinorNinth     = "o9b"
	NameShortDominantEleventh         = "11"
	NameShortMajorEleventh            = "M11"
	NameShortMinorMajorEleventh       = "mM11"
	NameShortMinorEleventh            = "m11"
	NameShortAugmentedMajorEleventh   = "+M11"
	NameShortAugmentedEleventh        = "+11"
	NameShortHalfDiminishedEleventh   = "011"
	NameShortDiminishedEleventh       = "o11"
	NameShortMajorThirteenth          = "M13"
	NameShortThirteenth               = "13"
	NameShortMinorMajorThirteenth     = "mM13"
	NameShortMinorThirteenth          = "m13"
	NameShortAugmentedMajorThirteenth = "+M13"
	NameShortAugmentedThirteenth      = "+13"
	NameShortHalfDiminishedThirteenth = "013"
	NameShortMajorTriad               = "M"
	NameShortMinorTriad               = "m"
	NameShortAugmentedTriad           = "+"
	NameShortDiminishedTriad          = "o"
)

var (
	MajorFifth               = New(tone.MustFromList(MajorFifthDegrees))
	MinorFifth               = New(tone.MustFromList(MinorFifthDegrees))
	AugmentedFifth           = New(tone.MustFromList(AugmentedFifthDegrees))
	DiminishedFifth          = New(tone.MustFromList(DiminishedFifthDegrees))
	DominantSeventh          = New(tone.MustFromList(DominantSeventhDegrees))
	MajorSeventh             = New(tone.MustFromList(MajorSeventhDegrees))
	MinorMajorSeventh        = New(tone.MustFromList(MinorMajorSeventhDegrees))
	MinorSeventh             = New(tone.MustFromList(MinorSeventhDegrees))
	AugmentedMajorSeventh    = New(tone.MustFromList(AugmentedMajorSeventhDegrees))
	AugmentedSeventh         = New(tone.MustFromList(AugmentedSeventhDegrees))
	HalfDiminishedSeventh    = New(tone.MustFromList(HalfDiminishedSeventhDegrees))
	DiminishedSeventh        = New(tone.MustFromList(DiminishedSeventhDegrees))
	SeventhFlatFive          = New(tone.MustFromList(SeventhFlatFiveDegrees))
	MajorNinth               = New(tone.MustFromList(MajorNinthDegrees))
	DominantNinth            = New(tone.MustFromList(DominantNinthDegrees))
	DominantMinorNinth       = New(tone.MustFromList(DominantMinorNinthDegrees))
	MinorMajorNinth          = New(tone.MustFromList(MinorMajorNinthDegrees))
	MinorNinth               = New(tone.MustFromList(MinorNinthDegrees))
	AugmentedMajorNinth      = New(tone.MustFromList(AugmentedMajorNinthDegrees))
	AugmentedDominantNinth   = New(tone.MustFromList(AugmentedDominantNinthDegrees))
	HalfDiminishedNinth      = New(tone.MustFromList(HalfDiminishedNinthDegrees))
	HalfDiminishedMinorNinth = New(tone.MustFromList(HalfDiminishedMinorNinthDegrees))
	DiminishedNinth          = New(tone.MustFromList(DiminishedNinthDegrees))
	DiminishedMinorNinth     = New(tone.MustFromList(DiminishedMinorNinthDegrees))
	DominantEleventh         = New(tone.MustFromList(DominantEleventhDegrees))
	MajorEleventh            = New(tone.MustFromList(MajorEleventhDegrees))
	MinorMajorEleventh       = New(tone.MustFromList(MinorMajorEleventhDegrees))
	MinorEleventh            = New(tone.MustFromList(MinorEleventhDegrees))
	AugmentedMajorEleventh   = New(tone.MustFromList(AugmentedMajorEleventhDegrees))
	AugmentedEleventh        = New(tone.MustFromList(AugmentedEleventhDegrees))
	HalfDiminishedEleventh   = New(tone.MustFromList(HalfDiminishedEleventhDegrees))
	DiminishedEleventh       = New(tone.MustFromList(DiminishedEleventhDegrees))
	MajorThirteenth          = New(tone.MustFromList(MajorThirteenthDegrees))
	Thirteenth               = New(tone.MustFromList(ThirteenthDegrees))
	MinorMajorThirteenth     = New(tone.MustFromList(MinorMajorThirteenthDegrees))
	MinorThirteenth          = New(tone.MustFromList(MinorThirteenthDegrees))
	AugmentedMajorThirteenth = New(tone.MustFromList(AugmentedMajorThirteenthDegrees))
	AugmentedThirteenth      = New(tone.MustFromList(AugmentedThirteenthDegrees))
	HalfDiminishedThirteenth = New(tone.MustFromList(HalfDiminishedThirteenthDegrees))
	MajorTriad               = MajorFifth
	MinorTriad               = MinorFifth
	AugmentedTriad           = AugmentedFifth
	DiminishedTriad          = DiminishedFifth

	NameToChord = map[string]Chord{
		NameMajorFifth:               MajorFifth,
		NameMinorFifth:               MinorFifth,
		NameAugmentedFifth:           AugmentedFifth,
		NameDiminishedFifth:          DiminishedFifth,
		NameDominantSeventh:          DominantSeventh,
		NameMajorSeventh:             MajorSeventh,
		NameMinorMajorSeventh:        MinorMajorSeventh,
		NameMinorSeventh:             MinorSeventh,
		NameAugmentedMajorSeventh:    AugmentedMajorSeventh,
		NameAugmentedSeventh:         AugmentedSeventh,
		NameHalfDiminishedSeventh:    HalfDiminishedSeventh,
		NameDiminishedSeventh:        DiminishedSeventh,
		NameSeventhFlatFive:          SeventhFlatFive,
		NameMajorNinth:               MajorNinth,
		NameDominantNinth:            DominantNinth,
		NameDominantMinorNinth:       DominantMinorNinth,
		NameMinorMajorNinth:          MinorMajorNinth,
		NameMinorNinth:               MinorNinth,
		NameAugmentedMajorNinth:      AugmentedMajorNinth,
		NameAugmentedDominantNinth:   AugmentedDominantNinth,
		NameHalfDiminishedNinth:      HalfDiminishedNinth,
		NameHalfDiminishedMinorNinth: HalfDiminishedMinorNinth,
		NameDiminishedNinth:          DiminishedNinth,
		NameDiminishedMinorNinth:     DiminishedMinorNinth,
		NameDominantEleventh:         DominantEleventh,
		NameMajorEleventh:            MajorEleventh,
		NameMinorMajorEleventh:       MinorMajorEleventh,
		NameMinorEleventh:            MinorEleventh,
		NameAugmentedMajorEleventh:   AugmentedMajorEleventh,
		NameAugmentedEleventh:        AugmentedEleventh,
		NameHalfDiminishedEleventh:   HalfDiminishedEleventh,
		NameDiminishedEleventh:       DiminishedEleventh,
		NameMajorThirteenth:          MajorThirteenth,
		NameThirteenth:               Thirteenth,
		NameMinorMajorThirteenth:     MinorMajorThirteenth,
		NameMinorThirteenth:          MinorThirteenth,
		NameAugmentedMajorThirteenth: AugmentedMajorThirteenth,
		NameAugmentedThirteenth:      AugmentedThirteenth,
		NameHalfDiminishedThirteenth: HalfDiminishedThirteenth,
		NameMajorTriad:               MajorFifth,
		NameMinorTriad:               MinorFifth,
		NameAugmentedTriad:           AugmentedFifth,
		NameDiminishedTriad:          DiminishedFifth,
	}

	NameLongToChord = map[string]Chord{
		NameLongMajorFifth:               MajorFifth,
		NameLongMinorFifth:               MinorFifth,
		NameLongAugmentedFifth:           AugmentedFifth,
		NameLongDiminishedFifth:          DiminishedFifth,
		NameLongDominantSeventh:          DominantSeventh,
		NameLongMajorSeventh:             MajorSeventh,
		NameLongMinorMajorSeventh:        MinorMajorSeventh,
		NameLongMinorSeventh:             MinorSeventh,
		NameLongAugmentedMajorSeventh:    AugmentedMajorSeventh,
		NameLongAugmentedSeventh:         AugmentedSeventh,
		NameLongHalfDiminishedSeventh:    HalfDiminishedSeventh,
		NameLongDiminishedSeventh:        DiminishedSeventh,
		NameLongSeventhFlatFive:          SeventhFlatFive,
		NameLongMajorNinth:               MajorNinth,
		NameLongDominantNinth:            DominantNinth,
		NameLongDominantMinorNinth:       DominantMinorNinth,
		NameLongMinorMajorNinth:          MinorMajorNinth,
		NameLongMinorNinth:               MinorNinth,
		NameLongAugmentedMajorNinth:      AugmentedMajorNinth,
		NameLongAugmentedDominantNinth:   AugmentedDominantNinth,
		NameLongHalfDiminishedNinth:      HalfDiminishedNinth,
		NameLongHalfDiminishedMinorNinth: HalfDiminishedMinorNinth,
		NameLongDiminishedNinth:          DiminishedNinth,
		NameLongDiminishedMinorNinth:     DiminishedMinorNinth,
		NameLongDominantEleventh:         DominantEleventh,
		NameLongMajorEleventh:            MajorEleventh,
		NameLongMinorMajorEleventh:       MinorMajorEleventh,
		NameLongMinorEleventh:            MinorEleventh,
		NameLongAugmentedMajorEleventh:   AugmentedMajorEleventh,
		NameLongAugmentedEleventh:        AugmentedEleventh,
		NameLongHalfDiminishedEleventh:   HalfDiminishedEleventh,
		NameLongDiminishedEleventh:       DiminishedEleventh,
		NameLongMajorThirteenth:          MajorThirteenth,
		NameLongDominantThirteenth:       Thirteenth,
		NameLongMinorMajorThirteenth:     MinorMajorThirteenth,
		NameLongMinorThirteenth:          MinorThirteenth,
		NameLongAugmentedMajorThirteenth: AugmentedMajorThirteenth,
		NameLongAugmentedThirteenth:      AugmentedThirteenth,
		NameLongHalfDiminishedThirteenth: HalfDiminishedThirteenth,
		NameLongMajorTriad:               MajorFifth,
		NameLongMinorTriad:               MinorFifth,
		NameLongAugmentedTriad:           AugmentedFifth,
		NameLongDiminishedTriad:          DiminishedFifth,
	}

	NameShortToChord = map[string]Chord{
		NameShortMajorFifth:               MajorFifth,
		NameShortMinorFifth:               MinorFifth,
		NameShortAugmentedFifth:           AugmentedFifth,
		NameShortDiminishedFifth:          DiminishedFifth,
		NameShortDominantSeventh:          DominantSeventh,
		NameShortMajorSeventh:             MajorSeventh,
		NameShortMinorMajorSeventh:        MinorMajorSeventh,
		NameShortMinorSeventh:             MinorSeventh,
		NameShortAugmentedMajorSeventh:    AugmentedMajorSeventh,
		NameShortAugmentedSeventh:         AugmentedSeventh,
		NameShortHalfDiminishedSeventh:    HalfDiminishedSeventh,
		NameShortDiminishedSeventh:        DiminishedSeventh,
		NameShortSeventhFlatFive:          SeventhFlatFive,
		NameShortMajorNinth:               MajorNinth,
		NameShortDominantNinth:            DominantNinth,
		NameShortDominantMinorNinth:       DominantMinorNinth,
		NameShortMinorMajorNinth:          MinorMajorNinth,
		NameShortMinorNinth:               MinorNinth,
		NameShortAugmentedMajorNinth:      AugmentedMajorNinth,
		NameShortAugmentedDominantNinth:   AugmentedDominantNinth,
		NameShortHalfDiminishedNinth:      HalfDiminishedNinth,
		NameShortHalfDiminishedMinorNinth: HalfDiminishedMinorNinth,
		NameShortDiminishedNinth:          DiminishedNinth,
		NameShortDiminishedMinorNinth:     DiminishedMinorNinth,
		NameShortDominantEleventh:         DominantEleventh,
		NameShortMajorEleventh:            MajorEleventh,
		NameShortMinorMajorEleventh:       MinorMajorEleventh,
		NameShortMinorEleventh:            MinorEleventh,
		NameShortAugmentedMajorEleventh:   AugmentedMajorEleventh,
		NameShortAugmentedEleventh:        AugmentedEleventh,
		NameShortHalfDiminishedEleventh:   HalfDiminishedEleventh,
		NameShortDiminishedEleventh:       DiminishedEleventh,
		NameShortMajorThirteenth:          MajorThirteenth,
		NameShortThirteenth:               Thirteenth,
		NameShortMinorMajorThirteenth:     MinorMajorThirteenth,
		NameShortMinorThirteenth:          MinorThirteenth,
		NameShortAugmentedMajorThirteenth: AugmentedMajorThirteenth,
		NameShortAugmentedThirteenth:      AugmentedThirteenth,
		NameShortHalfDiminishedThirteenth: HalfDiminishedThirteenth,
		NameShortMajorTriad:               MajorFifth,
		NameShortMinorTriad:               MinorFifth,
		NameShortAugmentedTriad:           AugmentedFifth,
		NameShortDiminishedTriad:          DiminishedFifth,
	}

	DegreesToName = map[string]string{
		MajorFifthDegrees:               NameMajorFifth,
		MinorFifthDegrees:               NameMinorFifth,
		AugmentedFifthDegrees:           NameAugmentedFifth,
		DiminishedFifthDegrees:          NameDiminishedFifth,
		DominantSeventhDegrees:          NameDominantSeventh,
		MajorSeventhDegrees:             NameMajorSeventh,
		MinorMajorSeventhDegrees:        NameMinorMajorSeventh,
		MinorSeventhDegrees:             NameMinorSeventh,
		AugmentedMajorSeventhDegrees:    NameAugmentedMajorSeventh,
		AugmentedSeventhDegrees:         NameAugmentedSeventh,
		HalfDiminishedSeventhDegrees:    NameHalfDiminishedSeventh,
		DiminishedSeventhDegrees:        NameDiminishedSeventh,
		SeventhFlatFiveDegrees:          NameSeventhFlatFive,
		MajorNinthDegrees:               NameMajorNinth,
		DominantNinthDegrees:            NameDominantNinth,
		DominantMinorNinthDegrees:       NameDominantMinorNinth,
		MinorMajorNinthDegrees:          NameMinorMajorNinth,
		MinorNinthDegrees:               NameMinorNinth,
		AugmentedMajorNinthDegrees:      NameAugmentedMajorNinth,
		AugmentedDominantNinthDegrees:   NameAugmentedDominantNinth,
		HalfDiminishedNinthDegrees:      NameHalfDiminishedNinth,
		HalfDiminishedMinorNinthDegrees: NameHalfDiminishedMinorNinth,
		DiminishedNinthDegrees:          NameDiminishedNinth,
		DiminishedMinorNinthDegrees:     NameDiminishedMinorNinth,
		DominantEleventhDegrees:         NameDominantEleventh,
		MajorEleventhDegrees:            NameMajorEleventh,
		MinorMajorEleventhDegrees:       NameMinorMajorEleventh,
		MinorEleventhDegrees:            NameMinorEleventh,
		AugmentedMajorEleventhDegrees:   NameAugmentedMajorEleventh,
		AugmentedEleventhDegrees:        NameAugmentedEleventh,
		HalfDiminishedEleventhDegrees:   NameHalfDiminishedEleventh,
		DiminishedEleventhDegrees:       NameDiminishedEleventh,
		MajorThirteenthDegrees:          NameMajorThirteenth,
		ThirteenthDegrees:               NameThirteenth,
		MinorMajorThirteenthDegrees:     NameMinorMajorThirteenth,
		MinorThirteenthDegrees:          NameMinorThirteenth,
		AugmentedMajorThirteenthDegrees: NameAugmentedMajorThirteenth,
		AugmentedThirteenthDegrees:      NameAugmentedThirteenth,
		HalfDiminishedThirteenthDegrees: NameHalfDiminishedThirteenth,
	}

	DegreesToNameLong = map[string]string{
		MajorFifthDegrees:               NameLongMajorFifth,
		MinorFifthDegrees:               NameLongMinorFifth,
		AugmentedFifthDegrees:           NameLongAugmentedFifth,
		DiminishedFifthDegrees:          NameLongDiminishedFifth,
		DominantSeventhDegrees:          NameLongDominantSeventh,
		MajorSeventhDegrees:             NameLongMajorSeventh,
		MinorMajorSeventhDegrees:        NameLongMinorMajorSeventh,
		MinorSeventhDegrees:             NameLongMinorSeventh,
		AugmentedMajorSeventhDegrees:    NameLongAugmentedMajorSeventh,
		AugmentedSeventhDegrees:         NameLongAugmentedSeventh,
		HalfDiminishedSeventhDegrees:    NameLongHalfDiminishedSeventh,
		DiminishedSeventhDegrees:        NameLongDiminishedSeventh,
		SeventhFlatFiveDegrees:          NameLongSeventhFlatFive,
		MajorNinthDegrees:               NameLongMajorNinth,
		DominantNinthDegrees:            NameLongDominantNinth,
		DominantMinorNinthDegrees:       NameLongDominantMinorNinth,
		MinorMajorNinthDegrees:          NameLongMinorMajorNinth,
		MinorNinthDegrees:               NameLongMinorNinth,
		AugmentedMajorNinthDegrees:      NameLongAugmentedMajorNinth,
		AugmentedDominantNinthDegrees:   NameLongAugmentedDominantNinth,
		HalfDiminishedNinthDegrees:      NameLongHalfDiminishedNinth,
		HalfDiminishedMinorNinthDegrees: NameLongHalfDiminishedMinorNinth,
		DiminishedNinthDegrees:          NameLongDiminishedNinth,
		DiminishedMinorNinthDegrees:     NameLongDiminishedMinorNinth,
		DominantEleventhDegrees:         NameLongDominantEleventh,
		MajorEleventhDegrees:            NameLongMajorEleventh,
		MinorMajorEleventhDegrees:       NameLongMinorMajorEleventh,
		MinorEleventhDegrees:            NameLongMinorEleventh,
		AugmentedMajorEleventhDegrees:   NameLongAugmentedMajorEleventh,
		AugmentedEleventhDegrees:        NameLongAugmentedEleventh,
		HalfDiminishedEleventhDegrees:   NameLongHalfDiminishedEleventh,
		DiminishedEleventhDegrees:       NameLongDiminishedEleventh,
		MajorThirteenthDegrees:          NameLongMajorThirteenth,
		ThirteenthDegrees:               NameLongDominantThirteenth,
		MinorMajorThirteenthDegrees:     NameLongMinorMajorThirteenth,
		MinorThirteenthDegrees:          NameLongMinorThirteenth,
		AugmentedMajorThirteenthDegrees: NameLongAugmentedMajorThirteenth,
		AugmentedThirteenthDegrees:      NameLongAugmentedThirteenth,
		HalfDiminishedThirteenthDegrees: NameLongHalfDiminishedThirteenth,
	}

	DegreesToNameShort = map[string]string{
		MajorFifthDegrees:               NameShortMajorFifth,
		MinorFifthDegrees:               NameShortMinorFifth,
		AugmentedFifthDegrees:           NameShortAugmentedFifth,
		DiminishedFifthDegrees:          NameShortDiminishedFifth,
		DominantSeventhDegrees:          NameShortDominantSeventh,
		MajorSeventhDegrees:             NameShortMajorSeventh,
		MinorMajorSeventhDegrees:        NameShortMinorMajorSeventh,
		MinorSeventhDegrees:             NameShortMinorSeventh,
		AugmentedMajorSeventhDegrees:    NameShortAugmentedMajorSeventh,
		AugmentedSeventhDegrees:         NameShortAugmentedSeventh,
		HalfDiminishedSeventhDegrees:    NameShortHalfDiminishedSeventh,
		DiminishedSeventhDegrees:        NameShortDiminishedSeventh,
		SeventhFlatFiveDegrees:          NameShortSeventhFlatFive,
		MajorNinthDegrees:               NameShortMajorNinth,
		DominantNinthDegrees:            NameShortDominantNinth,
		DominantMinorNinthDegrees:       NameShortDominantMinorNinth,
		MinorMajorNinthDegrees:          NameShortMinorMajorNinth,
		MinorNinthDegrees:               NameShortMinorNinth,
		AugmentedMajorNinthDegrees:      NameShortAugmentedMajorNinth,
		AugmentedDominantNinthDegrees:   NameShortAugmentedDominantNinth,
		HalfDiminishedNinthDegrees:      NameShortHalfDiminishedNinth,
		HalfDiminishedMinorNinthDegrees: NameShortHalfDiminishedMinorNinth,
		DiminishedNinthDegrees:          NameShortDiminishedNinth,
		DiminishedMinorNinthDegrees:     NameShortDiminishedMinorNinth,
		DominantEleventhDegrees:         NameShortDominantEleventh,
		MajorEleventhDegrees:            NameShortMajorEleventh,
		MinorMajorEleventhDegrees:       NameShortMinorMajorEleventh,
		MinorEleventhDegrees:            NameShortMinorEleventh,
		AugmentedMajorEleventhDegrees:   NameShortAugmentedMajorEleventh,
		AugmentedEleventhDegrees:        NameShortAugmentedEleventh,
		HalfDiminishedEleventhDegrees:   NameShortHalfDiminishedEleventh,
		DiminishedEleventhDegrees:       NameShortDiminishedEleventh,
		MajorThirteenthDegrees:          NameShortMajorThirteenth,
		ThirteenthDegrees:               NameShortThirteenth,
		MinorMajorThirteenthDegrees:     NameShortMinorMajorThirteenth,
		MinorThirteenthDegrees:          NameShortMinorThirteenth,
		AugmentedMajorThirteenthDegrees: NameShortAugmentedMajorThirteenth,
		AugmentedThirteenthDegrees:      NameShortAugmentedThirteenth,
		HalfDiminishedThirteenthDegrees: NameShortHalfDiminishedThirteenth,
	}
)

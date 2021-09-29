package scale

import (
	"github.com/Insulince/jtone/pkg/tone"
)

const (
	// Sourced from https://en.wikipedia.org/wiki/List_of_musical_scales_and_modes
	// Note that all instances of flat tones were replaces with their sharp
	// counterparts. This is done to allow easier string matching of names. It
	// would be difficult to know that a user played a natural minor scale when
	// the degrees can be represented in multiple ways. Using all sharps fixes
	// this problem, and the preference of sharps over flats is a consistent
	// preference throughout this codebase.

	AcousticDegrees           = "1,2,3,4s,5,6,6s"
	NaturalMinorDegrees       = "1,2,2s,4,5,5s,6s"
	AlgerianDegrees           = "1,2,2s,4s,5,5s,7"
	AlteredDegrees            = "1,1s,2s,3,4s,5s,6s" // Duplicates LocrianDegrees
	AugmentedDegrees          = "1,2s,3,5,5s,7"
	BepobDominantDegrees      = "1,2,3,4,5,6,6s,7"
	BluesDegrees              = "1,2s,4,4s,5,6s"
	ChromaticDegrees          = "1,1s,2,2s,3,4,4s,5,5s,6,6s,7"
	DorianDegrees             = "1,2,2s,4,5,6,6s"
	DoubleHarmonicDegrees     = "1,1s,3,4,5,5s,7"
	EnigmaticDegrees          = "1,1s,3,4s,5s,6s,7"
	FlamencoDegrees           = "1,1s,3,4,5,5s,7" // Duplicates DoubleHarmonicDegrees
	GypsyDegrees              = "1,2,2s,4s,5,5s,6s"
	HalfDiminishedDegrees     = "1,2,2s,4,4s,5s,7"
	HarmonicMinorDegrees      = "1,2,3,4,5,5s,7"
	HirajoshiDegrees          = "1,3,4s,5,7"
	HungarianMinorDegrees     = "1,2,2s,4s,5,5s,7" // Duplicates AlgerianDegrees
	HungarianMajorDegrees     = "1,2s,3,4s,5,6,6s"
	InDegrees                 = "1,1s,4,5,5s"
	InsenDegrees              = "1,1s,4,5,6s"
	MajorDegrees              = "1,2,3,4,5,6,7"
	IstrianDegrees            = "1,1s,2s,3,4s,5"
	IwatoDegrees              = "1,1s,4,4s,6s"
	LocrianDegrees            = "1,1s,2s,4,4s,5s,6s"
	LydianAugmentedDegrees    = "1,2,3,4s,5s,6,7"
	LydianDegrees             = "1,2,3,4s,5,6,7"
	MajorBebopDegrees         = "1,2,3,4,5,5s,6,7"
	MajorLocrianDegrees       = "1,2,3,4,4s,5s,6s"
	MajorPentatonicDegrees    = "1,2,3,5,6"
	MelodicMinorDegrees       = "1,2,2s,4,5,6,7"
	MinorPentatonicDegrees    = "1,2s,4,5,6s"
	MixolydianDegrees         = "1,2,3,4,5,6,6s"
	NeapolitanMajorDegrees    = "1,1s,2s,4,5,6,7"
	NepolitanMinorDegrees     = "1,1s,2s,4,5,5s,7"
	Octatonic1Degrees         = "1,2,2s,4,4s,5s,6,7"
	Octatonic2Degrees         = "1,1s,2s,3,4s,5,6,6s"
	PersianDegrees            = "1,1s,3,4,4s,5s,7"
	PhrygianDominantDegrees   = "1,1s,3,4,5,5s,6s"
	PhrygianDegrees           = "1,1s,2s,4,5,5s,6s"
	PrometheusDegrees         = "1,2,3,4s,6,6s"
	HarmonicsDegrees          = "1,2s,3,4,5,6"
	TritoneDegrees            = "1,1s,3,4s,5,6s"
	TwoSemitoneTritoneDegrees = "1,1s,2,4s,5,5s"
	UkrainianDorianDegrees    = "1,2,2s,4s,5,6,6s"
	WholeToneDegrees          = "1,2,3,4s,5s,6s"
	YoDegrees                 = "1,2s,4,5,6s" // Duplicates MinorPentatonicDegrees

	NameAcoustic           = "acoustic"
	NameNaturalMinor       = "natural-minor"
	NameAlgerian           = "algerian"
	NameAltered            = "altered"
	NameAugmented          = "augmented"
	NameBepobDominant      = "bepob-dominant"
	NameBlues              = "blues"
	NameChromatic          = "chromatic"
	NameDorian             = "dorian"
	NameDoubleHarmonic     = "double-harmonic"
	NameEnigmatic          = "enigmatic"
	NameFlamenco           = "flamenco"
	NameGypsy              = "gypsy"
	NameHalfDiminished     = "half-diminished"
	NameHarmonicMinor      = "harmonic-minor"
	NameHirajoshi          = "hirajoshi"
	NameHungarianMinor     = "hungarian-minor"
	NameHungarianMajor     = "hungarian-major"
	NameIn                 = "in"
	NameInsen              = "insen"
	NameMajor              = "major"
	NameIstrian            = "istrian"
	NameIwato              = "iwato"
	NameLocrian            = "locrian"
	NameLydianAugmented    = "lydian-augmented"
	NameLydian             = "lydian"
	NameMajorBebop         = "major-bebop"
	NameMajorLocrian       = "major-locrian"
	NameMajorPentatonic    = "major-pentatonic"
	NameMelodicMinor       = "melodic-minor"
	NameMinorPentatonic    = "minor-pentatonic"
	NameMixolydian         = "mixolydian"
	NameNeapolitanMajor    = "neapolitan-major"
	NameNepolitanMinor     = "nepolitan-minor"
	NameOctatonic1         = "octatonic-1"
	NameOctatonic2         = "octatonic-2"
	NamePersian            = "persian"
	NamePhrygianDominant   = "phrygian-dominant"
	NamePhrygian           = "phrygian"
	NamePrometheus         = "prometheus"
	NameHarmonics          = "harmonics"
	NameTritone            = "tritone"
	NameTwoSemitoneTritone = "two-semitone-tritone"
	NameUkrainianDorian    = "ukrainian-dorian"
	NameWholeTone          = "whole-tone"
	NameYo                 = "yo"
)

var (
	Acoustic           = New(tone.MustFromList(AcousticDegrees))
	NaturalMinor       = New(tone.MustFromList(NaturalMinorDegrees))
	Algerian           = New(tone.MustFromList(AlgerianDegrees))
	Altered            = New(tone.MustFromList(AlteredDegrees))
	Augmented          = New(tone.MustFromList(AugmentedDegrees))
	BepobDominant      = New(tone.MustFromList(BepobDominantDegrees))
	Blues              = New(tone.MustFromList(BluesDegrees))
	Chromatic          = New(tone.MustFromList(ChromaticDegrees))
	Dorian             = New(tone.MustFromList(DorianDegrees))
	DoubleHarmonic     = New(tone.MustFromList(DoubleHarmonicDegrees))
	Enigmatic          = New(tone.MustFromList(EnigmaticDegrees))
	Flamenco           = New(tone.MustFromList(FlamencoDegrees))
	Gypsy              = New(tone.MustFromList(GypsyDegrees))
	HalfDiminished     = New(tone.MustFromList(HalfDiminishedDegrees))
	HarmonicMinor      = New(tone.MustFromList(HarmonicMinorDegrees))
	Hirajoshi          = New(tone.MustFromList(HirajoshiDegrees))
	HungarianMinor     = New(tone.MustFromList(HungarianMinorDegrees))
	HungarianMajor     = New(tone.MustFromList(HungarianMajorDegrees))
	In                 = New(tone.MustFromList(InDegrees))
	Insen              = New(tone.MustFromList(InsenDegrees))
	Major              = New(tone.MustFromList(MajorDegrees))
	Istrian            = New(tone.MustFromList(IstrianDegrees))
	Iwato              = New(tone.MustFromList(IwatoDegrees))
	Locrian            = New(tone.MustFromList(LocrianDegrees))
	LydianAugmented    = New(tone.MustFromList(LydianAugmentedDegrees))
	Lydian             = New(tone.MustFromList(LydianDegrees))
	MajorBebop         = New(tone.MustFromList(MajorBebopDegrees))
	MajorLocrian       = New(tone.MustFromList(MajorLocrianDegrees))
	MajorPentatonic    = New(tone.MustFromList(MajorPentatonicDegrees))
	MelodicMinor       = New(tone.MustFromList(MelodicMinorDegrees))
	MinorPentatonic    = New(tone.MustFromList(MinorPentatonicDegrees))
	Mixolydian         = New(tone.MustFromList(MixolydianDegrees))
	NeapolitanMajor    = New(tone.MustFromList(NeapolitanMajorDegrees))
	NepolitanMinor     = New(tone.MustFromList(NepolitanMinorDegrees))
	Octatonic1         = New(tone.MustFromList(Octatonic1Degrees))
	Octatonic2         = New(tone.MustFromList(Octatonic2Degrees))
	Persian            = New(tone.MustFromList(PersianDegrees))
	PhrygianDominant   = New(tone.MustFromList(PhrygianDominantDegrees))
	Phrygian           = New(tone.MustFromList(PhrygianDegrees))
	Prometheus         = New(tone.MustFromList(PrometheusDegrees))
	Harmonics          = New(tone.MustFromList(HarmonicsDegrees))
	Tritone            = New(tone.MustFromList(TritoneDegrees))
	TwoSemitoneTritone = New(tone.MustFromList(TwoSemitoneTritoneDegrees))
	UkrainianDorian    = New(tone.MustFromList(UkrainianDorianDegrees))
	WholeTone          = New(tone.MustFromList(WholeToneDegrees))
	Yo                 = New(tone.MustFromList(YoDegrees))

	NameToScale = map[string]Scale{
		NameAcoustic:           Acoustic,
		NameNaturalMinor:       NaturalMinor,
		NameAlgerian:           Algerian,
		NameAltered:            Altered,
		NameAugmented:          Augmented,
		NameBepobDominant:      BepobDominant,
		NameBlues:              Blues,
		NameChromatic:          Chromatic,
		NameDorian:             Dorian,
		NameDoubleHarmonic:     DoubleHarmonic,
		NameEnigmatic:          Enigmatic,
		NameFlamenco:           Flamenco,
		NameGypsy:              Gypsy,
		NameHalfDiminished:     HalfDiminished,
		NameHarmonicMinor:      HarmonicMinor,
		NameHirajoshi:          Hirajoshi,
		NameHungarianMinor:     HungarianMinor,
		NameHungarianMajor:     HungarianMajor,
		NameIn:                 In,
		NameInsen:              Insen,
		NameMajor:              Major,
		NameIstrian:            Istrian,
		NameIwato:              Iwato,
		NameLocrian:            Locrian,
		NameLydianAugmented:    LydianAugmented,
		NameLydian:             Lydian,
		NameMajorBebop:         MajorBebop,
		NameMajorLocrian:       MajorLocrian,
		NameMajorPentatonic:    MajorPentatonic,
		NameMelodicMinor:       MelodicMinor,
		NameMinorPentatonic:    MinorPentatonic,
		NameMixolydian:         Mixolydian,
		NameNeapolitanMajor:    NeapolitanMajor,
		NameNepolitanMinor:     NepolitanMinor,
		NameOctatonic1:         Octatonic1,
		NameOctatonic2:         Octatonic2,
		NamePersian:            Persian,
		NamePhrygianDominant:   PhrygianDominant,
		NamePhrygian:           Phrygian,
		NamePrometheus:         Prometheus,
		NameHarmonics:          Harmonics,
		NameTritone:            Tritone,
		NameTwoSemitoneTritone: TwoSemitoneTritone,
		NameUkrainianDorian:    UkrainianDorian,
		NameWholeTone:          WholeTone,
		NameYo:                 Yo,
	}

	DegreesToName = map[string]string{
		AcousticDegrees:           NameAcoustic,
		NaturalMinorDegrees:       NameNaturalMinor,
		AlgerianDegrees:           NameAlgerian,
		AugmentedDegrees:          NameAugmented,
		BepobDominantDegrees:      NameBepobDominant,
		BluesDegrees:              NameBlues,
		ChromaticDegrees:          NameChromatic,
		DorianDegrees:             NameDorian,
		DoubleHarmonicDegrees:     NameDoubleHarmonic,
		EnigmaticDegrees:          NameEnigmatic,
		GypsyDegrees:              NameGypsy,
		HalfDiminishedDegrees:     NameHalfDiminished,
		HarmonicMinorDegrees:      NameHarmonicMinor,
		HirajoshiDegrees:          NameHirajoshi,
		HungarianMajorDegrees:     NameHungarianMajor,
		InDegrees:                 NameIn,
		InsenDegrees:              NameInsen,
		MajorDegrees:              NameMajor,
		IstrianDegrees:            NameIstrian,
		IwatoDegrees:              NameIwato,
		LocrianDegrees:            NameLocrian,
		LydianAugmentedDegrees:    NameLydianAugmented,
		LydianDegrees:             NameLydian,
		MajorBebopDegrees:         NameMajorBebop,
		MajorLocrianDegrees:       NameMajorLocrian,
		MajorPentatonicDegrees:    NameMajorPentatonic,
		MelodicMinorDegrees:       NameMelodicMinor,
		MinorPentatonicDegrees:    NameMinorPentatonic,
		MixolydianDegrees:         NameMixolydian,
		NeapolitanMajorDegrees:    NameNeapolitanMajor,
		NepolitanMinorDegrees:     NameNepolitanMinor,
		Octatonic1Degrees:         NameOctatonic1,
		Octatonic2Degrees:         NameOctatonic2,
		PersianDegrees:            NamePersian,
		PhrygianDominantDegrees:   NamePhrygianDominant,
		PhrygianDegrees:           NamePhrygian,
		PrometheusDegrees:         NamePrometheus,
		HarmonicsDegrees:          NameHarmonics,
		TritoneDegrees:            NameTritone,
		TwoSemitoneTritoneDegrees: NameTwoSemitoneTritone,
		UkrainianDorianDegrees:    NameUkrainianDorian,
		WholeToneDegrees:          NameWholeTone,
	}
)

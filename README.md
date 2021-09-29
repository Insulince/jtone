# jtone

A CLI tool for building and playing various sonic structures.

Note that this is an unfinished project and large portions of the functionality may be lacking or inconsistent.

## Installation

Note that I run in a Windows environment, adjust accordingly for your environment.

Install via `go get`:

```
go get -u github.com/Insulince/jtone
```

Remember to install dependencies:

```
go mod vendor
```

Build the `cmd/cli` package as you would any other executable:

```
go build -o ./cmd/cli/bin/jtone.exe ./cmd/cli
```

### IntelliJ Users

For those who use IntelliJ, a run configuration called `jtone:build` is included under the `.run` folder in the project root. This should be automatically picked up by the IDE, and running it will place the executable under `cmd/cli/bin/jtone.exe`.

## Usage

The CLI currently has one command, `play`. This command accepts the name of a structure you wish to play, as well as a notation "blueprint" of what that structure should sound like.

The 4 sonic structures you can create and play are:

- Tone - A single tone oscillating at a fixed hz value.
- Chord - A series of tones played _simultaneously_.
- Scale - A series of tones played _sequentially_.
- Progression - A series of chords played sequentially.

### Notation

(For a more technical explanation, see the [BNF grammar below](#bnf-grammar))

Notation starts at the tone level and is gradually built upon as we move up to more complex structures. At the core, a tone can be represented as a note (**uppercase** `A`-`G`), an accidental (sharp, flat, or natural), and an octave (an integer).

Accidentals are represented by a lowercase `s` for sharp, and a lowercase `b` for flat. A natural accidental is represented by not specifying flat or sharp.

Octaves are optional and if left off a "reasonable" octave will be selected (4th/5th octave depending on the note). For example, if you simply want to hear a "C" note and don't care what octave, you don't need to specify one, and the 5th octave is selected automatically.

An important distinction to make here: Notes are different than tones. Notes refer to the 12 tones in an octave (in _any_ octave). Thare are only 12 notes (17 if you include the fact that sharps can be rewritten as flats, but these are ultimately two different representations of the same note). Tones on the other hand refer to an actual frequency. They are a note placed in an octave. Tones have notes, but notes by themselves do not have tones.

This is a regular expression for every valid way of representing a single tone, barring impossible tones like B sharp:

```regexp
[A-G][sb]?\d+?
```

Here are some examples of valid ways to request a note:

- `C` - C natural, in this case octave 5 will be selected - `523.25hz`
- `Fs3` - F sharp in the third octave - `185.00hz`
- `D8` - D natural in the ninth octave - `4698.64hz`
- `Gb` - G flat in the auto-selected 5th octave - `739.99hz`

#### Degrees

There is an important detour to take here about [degree notation](https://en.wikipedia.org/wiki/Degree_(music)), which is also a valid way of representing a tone. Degrees let you specify the way notes relate to each other without specifying what the actual notes are. This comes in handy with larger sonic structures like scales and chords, but is also applicable for tone, as there is a specific way of using degrees in the notation.

The quick and dirty in how degrees affect notation is that a degree represents a number of semitones that a "root" note should be shifted. Note that degrees also use sharps and flats and have the same restrictions that notes do in that similar to how there is no E sharp note there is also no 3 sharp degree.

Everything mentioned so far about notation was void of degrees and worked directly with a specific tone, but you can also request these structures in degree format and then provide the tone of the root degree, the root being the degree with the lowest value. This is done by specifying the root note, then a `#` symbol, then the degree of your tone. So here are some examples in the world of tones:

- `C4#1` - A tone of degree `1` is created, and the root should be set to `C4`. The result is a tone set to `C4`. This note is unchanged because the degree is `1`, which means the note should not be shifted at all.
- `C4#3` - A tone of degree `3` is created, and the root should be set to `C4`. The result is a tone set to `E4`. The reason for this is simply because you take the provided tone, `C4`, and shift it by the degree specified, `3`. Degree `3` is `4` semitones, so shifting `C4` that much gives us `E4`.
- `As6#5s` - A tone of degree `5s` is created, and the root should be set to `As6`. The result is a tone set to `Fs7`. The note is shifted from `As6` by `5s` degrees, which is `8` semitones.

So this may seem like overkill, and I would agree with you for tones. It is included on tones for completeness with the other structures, but this degree notation becomes vitally important in structures which use more than one tone.

Basically the TL;DR here is, _don't worry about degrees when writing tones unless you are doing something that requires them_. If you don't specify a tone via the `#` notation, then degree `1` will be used which is the same as leaving the note untouched.

### Playing Tones

You can play a tone with the following command:

```
jtone play tone <tone>
```

Where `tone` is a string identifying the tone you wish to play based on the earlier mentioned [notation](#notation).

### Playing Chords

Chords can be played with the following command:

```
jtone play chord <chord>
```

Valid chords come in 3 forms:

1. A named chord: `major-fifth`
2. A series of tones: `C,E,G`
3. A series of degrees: `1,3,5`

Each of the above are identical ways of playing the major fifth chord in the key of C in the fifth octave. The key of C is selected because it is the default, and the default octave for an unspecified C note is 5, so the root note for each of these chords will be `C5`. You can change the key by providing tone at the front of the chord seperated by a colon (`:`). Here are the same examples above, but in the key of `Ds3`:

1. `Ds3:major-fifth`
2. `Ds3:C,E,G`
3. `Ds3:1,3,5`

The list of named chords can be found [here](https://github.com/Insulince/jtone/pkg/chord/chords.go#L74). You can play one like so:

```
jtone play chord E3:mionor-seventh
```

This will play the minor seventh chord in the key of E in the third octave.

You can play a chord from a list of tones like so:

```
jtone play chord B:C,E,G
```

This will play the major fifth in the key of B. Admittedly, this is a strange way to write this chord, since you are writing the notes already and you know the key, you should probably just specify your desired tones, but this is  noted here for completeness. Remember, the key is always optional.

Lastly you can play a chord from a list of degrees like so:

```
jtone play chord Gs5:3,1,9s,9
```

This will play a custom chord rooted at G sharp in the fifth octave using the degrees 3, 1, 9 sharp, and 9. This leads to these tones being played simultaneously: `C6`, `Gs5`, `B6`, and `As6`. Even in this format the key is optional. If you don't provide a key, the root tone will be assumed to be `C5`, and the degrees will populate relative to that tone.

### Playing Scales

Scales and chords are very similar structures sonically and thus have similar ways of being played. The only difference really is in how they are output, chords play their tones _simultaneously_ and scales play their tones _sequentially_.

```
jtone play scale <scale>
```

Scales respect the same 3 modes and methods of specifying a root note as chords do, but they have a different list of named scales. You can see the list [here](https://github.com/Insulince/jtone/pkg/scale/scales.go#L63).

Here is an example:

```
jtone play scale F4:major-pentatonic
```

This plays the major-pentatonic scale in the key of F on octave 4.

```
jtone play scale Cs3:1,5,9,12,13s
```

This plays a custom scale rooted at C sharp in the third octave using the degrees 1, 5, 9, 12, and 13 sharp. This leads to these tones being played in order: `Cs3`, `Gs3`, `Ds4`, `Gs4`, `B4`.

Note that you will often see degree notation written with degrees in ascending order. This is because scales are typically played in a specific order, but it isn't a requirement of this application; degrees can be provided in any order you wish, just know that the lowest degree value is where your key will be rooted to.

### Playing Progressions

Progressions are a series of chords played sequentially.

```
jtone play progression <progression>
```

Progressions build on chord/scale notation in that they are a series of chords put together separated by semicolons (`;`). In my environment, which is Windows, I am required to wrap the whole progression notation in quotes to prevent windows from misinterpreting the command.

```
jtone play progression "major-fifth;major-seventh;major-ninth"
```

This plays the major-fifth chord, then the major-seventh chord, then the major-ninth chord in order. All in the default key of C on the fifth octave. You can control this behavior with the same expressiveness as with chords and scales by leveraging the notation accordingly:

```
jtone play progression "F4:1,3,5;Gs3:1,5,7;E3:major-fifth"
```

This plays two custom chords followed by a major fifth. The first custom chord is the chord corresponding to the degrees `1`, `3`, and `5`, rooted at `F4`, and the second being the chord corresponding to the degrees `1`, `5`, and `7`, rooted at `Gs3`. Then the major-fifth chord is played rooted at `E3`.

Unlike chords and scales, there is no way to specify an entire progression should fall into some key, you have to provide the key with each chord.

Additionally, there are no named progressions at this point.

## BNF Grammar

This is a _rough_ BNF grammar of how the notation system works.

```
<progression> ::= <chords>

     <chords> ::= <chord> ";" <chords> | <chord>
      <chord> ::= <tone> ":" <chord-body> | <chord-body>
 <chord-body> ::= <named-chord> | <tones> | <degrees>
<named-chord> ::= {x ∈ *[1]}

      <scale> ::= <tone> ":" <scale-body> | <scale-body>
 <scale-body> ::= <named-scale> | <tones> | <degrees>
<named-scale> ::= {x ∈ *[2]}

      <tones> ::= <tone> "," <tones> | <tone>
       <tone> ::= <tone-body> "#" <degree> | <tone-body>
  <tone-body> ::= <note> <octave> | <note>

       <note> ::= <letter> <accidental>
     <letter> ::= "A" | "B" | "C" | "D" | "E" | "F" | "G"

 <accidental> ::= "s" | "b" | ""

     <octave> ::= {x ∈ Z | x >= 0}

    <degrees> ::= <degree> "," <degrees> | <degree>
     <degree> ::= {x ∈ Z | x >= 0}
```
[*[1] List of named chords](https://github.com/Insulince/jtone/pkg/chord/chords.go#L74)

[*[2] List of named scales](https://github.com/Insulince/jtone/pkg/scale/scales.go#L63)

## Further Notes

- Sharp accidentals are preferred over flats anytime a tone is to be output. This is simply a consequence of the fact that there are multiple ways to represent the same tone and one must be chosen. I chose sharps. Despite this, though, flats are always valid to use, so feel free if that notation is more natural for you.
- I have found that some sources put accidentals before the note when describing a tone while others put the accidental after. This is another case where I had to make a decision, and I chose for accidentals to _always_ appear after the note.
- The library which provides sound, https://github.com/faiface/beep, for some reason has a delay of some sort that seems to be tied to "sample rate". This delay in playback causes the sounds to be cut short, so if you were playing a one-second tone you would only hear about 2/3rds of a second of playback. After some fiddling with different delay values and sample rates, I found that forcing the streamers to stay alive and silent for an extra 300 milliseconds prevents the sounds from being cut-off early. But I do not know if it is specific to my system or if the delay should be tweaked by some other parameters. In any case, it is defined [here](https://github.com/Insulince/jtone/pkg/util/const.go#L8) and you can adjust it if needed.

## TODO

- [ ] Revamp chords package, it feels rather... haphazard. There must be some way to standardize it better.
- [ ] Roman numeral chord progressions.
- [ ] Internal documentation.
- [ ] More tests.
- [ ] Come up with more things to add.

## Thanks

- [Casey Connor](https://www.youtube.com/c/CaseyConnor) - For his excellent series, [Music Theory Distilled](https://www.youtube.com/watch?v=mdEcLQ_RQPY), which explains in very digestible terms how the basics of music theory works and served as my inspiration to create this project. And for answering some of my questions about the series.
- [faiface](https://github.com/faiface) - For helping me come up with a valid `beep.Streamer` which can play back sine waves of a given frequency and for creating the library which drives the playback of this system.
- [Bill Hilton](https://www.youtube.com/c/BillHilton) - For answering some of my questions on the finer details of how chord notation works.

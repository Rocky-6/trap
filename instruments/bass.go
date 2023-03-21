package instruments

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func MkBass(path string, key string, cp [4]string) {
	clock := smf.MetricTicks(96)
	s := smf.New()
	s.TimeFormat = clock
	tr := smf.Track{}
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(70))

	// start
	// 1
	c := bassNote(keyNoteBass(key), cp[0])
	tr.Add(0, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
	tr.Add(clock.Ticks64th()*23, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))

	// 2
	c = bassNote(keyNoteBass(key), cp[1])
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))

	// 3
	c = bassNote(keyNoteBass(key), cp[2])
	tr.Add(clock.Ticks64th()*7, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
	tr.Add(clock.Ticks64th()*23, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))

	// 4
	c = bassNote(keyNoteBass(key), cp[3])
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
	// loop
	for loop := 0; loop < 2; loop++ {
		c = bassNote(keyNoteBass(key), cp[0])
		tr.Add(clock.Ticks64th()*7, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
		tr.Add(clock.Ticks64th()*23, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))

		c = bassNote(keyNoteBass(key), cp[1])
		tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
		tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))

		c = bassNote(keyNoteBass(key), cp[2])
		tr.Add(clock.Ticks64th()*7, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
		tr.Add(clock.Ticks64th()*23, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))

		c = bassNote(keyNoteBass(key), cp[3])
		tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
		tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, c, 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, c))
	}

	// end
	tr.Close(0)
	s.Add(tr)
	s.WriteFile(path + "/bass.mid")
}

func bassNote(keyNoteChord uint8, degree_name string) uint8 {
	root := keyNoteChord

	switch true {
	case check_regexp(`bVII`, degree_name):
		root += 10
	case check_regexp(`VII`, degree_name):
		root += 11
	case check_regexp(`bVI`, degree_name):
		root += 8
	case check_regexp(`VI`, degree_name):
		root += 9
	case check_regexp(`#IV`, degree_name):
		root += 6
	case check_regexp(`IV`, degree_name):
		root += 5
	case check_regexp(`V`, degree_name):
		root += 7
	case check_regexp(`bIII`, degree_name):
		root += 3
	case check_regexp(`III`, degree_name):
		root += 4
	case check_regexp(`bII`, degree_name):
		root += 1
	case check_regexp(`II`, degree_name):
		root += 2
	default:
	}

	return root
}

func keyNoteBass(key string) uint8 {
	noteNames := [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

	var note uint8

	for i, noteName := range noteNames {
		if key == noteName {
			note = uint8(i) + 36
			break
		}
	}

	if note == 0 {
		fmt.Println("入力された音階名が不正です")
		return 0
	}

	return note
}

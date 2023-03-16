package instruments

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func mkHihat() {
	clock := smf.MetricTicks(96)
	s := smf.New()
	s.TimeFormat = clock
	tr := smf.Track{}
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(70))

	// start
	// 1
	tr.Add(0, midi.NoteOn(0, midi.C(3), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	for i := 0; i < 3; i++ {
		tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}
	// 2
	tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	for i := 0; i < 4; i++ {
		tr.Add(0, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}
	for i := 0; i < 2; i++ {
		tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}
	// 3
	for i := 0; i < 4; i++ {
		tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}
	// 4
	for i := 0; i < 3; i++ {
		tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}
	for i := 0; i < 4; i++ {
		tr.Add(0, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}

	// loop
	for loop := 0; loop < 4; loop++ {
		// 1
		tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		for i := 0; i < 3; i++ {
			tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
			tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		}
		// 2
		tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		for i := 0; i < 4; i++ {
			tr.Add(0, midi.NoteOn(0, midi.C(3), 100))
			tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		}
		for i := 0; i < 2; i++ {
			tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
			tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		}
		// 3
		for i := 0; i < 4; i++ {
			tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
			tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		}
		// 4
		for i := 0; i < 3; i++ {
			tr.Add(clock.Ticks64th()*3, midi.NoteOn(0, midi.C(3), 100))
			tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		}
		for i := 0; i < 4; i++ {
			tr.Add(0, midi.NoteOn(0, midi.C(3), 100))
			tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		}

	}

	// end
	tr.Close(0)
	s.Add(tr)
	s.WriteFile("/home/rocky/trap/hihat.mid")
}

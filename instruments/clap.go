package instruments

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func mkClap() {
	clock := smf.MetricTicks(96)
	s := smf.New()
	s.TimeFormat = clock
	tr := smf.Track{}
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(70))

	// start
	// 1
	tr.Add(clock.Ticks64th()*16, midi.NoteOn(0, midi.C(3), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	tr.Add(clock.Ticks64th()*31, midi.NoteOn(0, midi.C(3), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))

	// loop
	for loop := 0; loop < 4; loop++ {
		tr.Add(clock.Ticks64th()*31, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
		tr.Add(clock.Ticks64th()*31, midi.NoteOn(0, midi.C(3), 100))
		tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(3)))
	}

	// end
	tr.Close(0)
	s.Add(tr)
	s.WriteFile("/home/rocky/trap/clap.mid")
}
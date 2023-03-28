package instruments

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func MkKick(path string) {
	clock := smf.MetricTicks(96)
	s := smf.New()
	s.TimeFormat = clock
	tr := smf.Track{}
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(70))

	// start
	tr.Add(0, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*23, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*7, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*23, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks64th()*15, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	// end

	tr.Close(0)
	s.Add(tr)
	s.WriteFile(path + "/kick.mid")
}

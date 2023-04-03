package instruments

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func MkClap(path string) {
	clock := smf.MetricTicks(96)
	s := smf.New()
	s.TimeFormat = clock
	tr := smf.Track{}
	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(140))

	// start
	tr.Add(clock.Ticks4th()*2, midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks4th()*4-clock.Ticks64th(), midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks4th()*4-clock.Ticks64th(), midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	tr.Add(clock.Ticks4th()*4-clock.Ticks64th(), midi.NoteOn(0, midi.C(5), 100))
	tr.Add(clock.Ticks64th(), midi.NoteOff(0, midi.C(5)))
	// end

	tr.Close(0)
	s.Add(tr)
	s.WriteFile(path + "/clap.mid")
}

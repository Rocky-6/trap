package main

import (
	"flag"

	"github.com/Rocky-6/trap/database"
	"github.com/Rocky-6/trap/instruments"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	key := flag.Arg(1)
	instruments.MkKick(path)
	instruments.MkClap(path)
	instruments.MkHihat(path)
	cp := database.DispChord()
	instruments.MkBass(path, key, cp)
	instruments.MkChord(path, key, cp)
	instruments.MkMelody(path, key)
}

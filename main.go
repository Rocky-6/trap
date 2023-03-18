package main

import (
	"flag"

	"github.com/Rocky-6/trap/instruments"
)

func main() {
	flag.Parse()
	path := flag.Arg(0)
	instruments.MkKick(path)
	instruments.MkClap(path)
	instruments.MkHihat(path)
}

package instruments

import (
	"testing"

	"github.com/Rocky-6/trap/database"
)

func TestUnit(t *testing.T) {
	cp := database.DispChord()
	MkBass("..", "C", cp)
	MkChord("..", "C", cp)
	MkMelody("..", "C")
}

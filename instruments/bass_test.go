package instruments

import (
	"testing"

	"github.com/Rocky-6/trap/database"
)

func TestMkBass(t *testing.T) {
	MkBass("..", "C", database.DispChord())
}

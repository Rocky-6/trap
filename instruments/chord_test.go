package instruments

import (
	"testing"

	"github.com/Rocky-6/trap/database"
)

func TestMkChord(t *testing.T) {
	MkChord("..", "C", database.DispChord())
}

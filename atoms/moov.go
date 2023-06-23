package atoms

import (
	"fmt"
	"os"
)

type Moov struct {
	*Atom
}

func NewMoov(atom *Atom, file *os.File) *Moov {
	movie := Moov{Atom: atom}
	// Since moov atom contains nested atoms, we need to seek back for future parses.
	//_, err := file.Seek(-int64(atom.Size-8), io.SeekCurrent)
	//if err != nil {
	//	panic(err)
	//}
	return &movie
}

func (a Moov) String() string {
	return fmt.Sprintf(
		"MovieAtom { Type: %s, Size: %d}",
		a.Type, a.Size,
	)
}

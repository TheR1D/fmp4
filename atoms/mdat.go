package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Mdat struct {
	RawData []byte
	*Atom
}

func NewMdat(atom *Atom, file *os.File) *Mdat {
	data := make([]byte, atom.Size-8)
	mdat := Mdat{Atom: atom, RawData: data}
	err := binary.Read(file, binary.BigEndian, &mdat.RawData)
	if err != nil {
		panic(err)
	}
	return &mdat
}

func (a Mdat) String() string {
	return fmt.Sprintf(
		"FileTypeAtom {"+
			"Type: %s, "+
			"Size: %d, "+
			"Data: %b ... "+
			"}",
		a.Type, a.Size, a.RawData[0:10],
	)
}

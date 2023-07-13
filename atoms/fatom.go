package atoms

import (
	"encoding/binary"
	"os"
)

type FullAtom struct {
	BaseAtom
	Version uint8
	Flags   [3]byte
}

func NewFullAtom(file *os.File) *FullAtom {
	fatom := &FullAtom{}
	if err := fatom.Parse(file); err != nil {
		panic(err)
	}
	return fatom
}

func (a *FullAtom) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, a); err != nil {
		return err
	}
	return nil
}

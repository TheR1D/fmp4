package atoms

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Atom interface {
	GetSize() uint32
	GetType() string
	Parse(*os.File) error
}

type BaseAtom struct {
	// TODO: Might be large input uint64.
	Size uint32
	Type [4]byte
}

func NewBaseAtom(file *os.File) *BaseAtom {
	baseAtom := &BaseAtom{}
	if err := baseAtom.Parse(file); err != nil {
		panic(err)
	}
	return baseAtom

}

// Parse Expecting seek pointer to be on the beginning of the atom.
// Parses Size and Type of atom and stops seek pointer on the beginning of atom data.
func (a *BaseAtom) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, a); err != nil {
		return err
	}
	return nil
}

func (a *BaseAtom) GetType() string {
	return string(a.Type[:])
}

func (a *BaseAtom) GetSize() uint32 {
	return a.Size
}

func (a *BaseAtom) String() string {
	return fmt.Sprintf("BaseAtom: {Type: %s, Size: %d}", a.Type, a.Size)
}

// SkipUntil Searches for atom with given type and moves
// seek pointer of file to the beginning of the atom.
func SkipUntil(atomType string, file *os.File) error {
	// Expecting file seek pointer to be on the beginning of the atom.
	for a := NewBaseAtom(file); a.GetType() != atomType; a = NewBaseAtom(file) {
		fmt.Println("Type:", a.GetType())
		s := int64(a.GetSize() - 8)
		if _, err := file.Seek(s, io.SeekCurrent); err != nil {
			return err
		}
	}
	// Ignore error because we are sure we can seek back.
	_, _ = file.Seek(-8, io.SeekCurrent)
	return nil
}

func MustSkipUntil(atomType string, file *os.File) {
	if err := SkipUntil(atomType, file); err != nil {
		panic(err)
	}
}

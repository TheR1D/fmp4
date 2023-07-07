package atoms

import (
	"fmt"
	"io"
	"os"
)

type ByteRange interface {
	StartsAt() int64
	EndsAt() int64
	Length() uint32
}

type SAtom struct {
	// Seekable Atom.
	Atom     Atom
	startsAt int64
}

func (a *SAtom) GetType() string {
	return a.Atom.GetType()
}

func (a *SAtom) GetSize() uint32 {
	return a.Atom.GetSize()
}

// Parse parses atom data, type and size of atom should be already parsed.
// Expecting seek pointer to be on the beginning of the atom data.
func (a *SAtom) Parse(file *os.File) error {
	a.startsAt, _ = file.Seek(0, io.SeekCurrent)
	return a.Atom.Parse(file)
}

func (a *SAtom) StartsAt() int64 {
	return a.startsAt
}

func (a *SAtom) EndsAt() int64 {
	return a.startsAt + int64(a.GetSize())
}

func (a *SAtom) Length() uint32 {
	return a.GetSize()
}

func (a *SAtom) String() string {
	str := fmt.Sprintf(
		"SAtom: StartsAt: %d, EndsAt: %d, %s",
		a.StartsAt(), a.EndsAt(), a.Atom,
	)
	return str
}

package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Moov struct {
	BaseAtom
	Mvhd *Mvhd
	Trak *Trak
}

func NewMoov(file *os.File) *Moov {
	moov := &Moov{}
	if err := moov.Parse(file); err != nil {
		panic(err)
	}
	return moov
}

func (a *Moov) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, &a.Size); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.Type); err != nil {
		return err
	}
	a.Mvhd = &Mvhd{}
	if err := a.Mvhd.Parse(file); err != nil {
		return err
	}
	if err := SkipUntil("trak", file); err != nil {
		return err
	}
	a.Trak = &Trak{}
	if err := a.Trak.Parse(file); err != nil {
		return err
	}
	return nil
}

func (a *Moov) String() string {
	return fmt.Sprintf(
		"moov: {Type: %s, Size: %d, Mvhd: {%s}, Trak: {%s}}",
		a.Type, a.Size, a.Mvhd, a.Trak,
	)
}

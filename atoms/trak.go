package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Trak struct {
	BaseAtom
	Tkhd *Tkhd
	Mdia *Mdia
}

func NewTrak(file *os.File) *Trak {
	trak := &Trak{}
	if err := trak.Parse(file); err != nil {
		panic(err)
	}
	return trak
}

func (a *Trak) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, &a.Size); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.Type); err != nil {
		return err
	}
	a.Tkhd = &Tkhd{}
	// tkhd should first atom nested in trak.
	if err := a.Tkhd.Parse(file); err != nil {
		return err
	}
	if err := SkipUntil("mdia", file); err != nil {
		return err
	}
	a.Mdia = &Mdia{}
	if err := a.Mdia.Parse(file); err != nil {
		return err
	}
	return nil
}

func (a *Trak) String() string {
	return fmt.Sprintf(
		"trak: {Type: %s, Size: %d, tkhd: {%s}, mdia: {%s}}",
		a.Type, a.Size, a.Tkhd, a.Mdia,
	)
}

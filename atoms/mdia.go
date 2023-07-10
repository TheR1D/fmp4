package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Mdia struct {
	BaseAtom
	Hdlr *Hdlr
	Mdhd *Mdhd
}

func NewMdia(file *os.File) *Mdia {
	mdia := &Mdia{}
	if err := mdia.Parse(file); err != nil {
		panic(err)
	}
	return mdia
}

func (a *Mdia) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, &a.Size); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.Type); err != nil {
		return err
	}
	if err := SkipUntil("mdhd", file); err != nil {
		return err
	}
	a.Mdhd = &Mdhd{}
	if err := a.Mdhd.Parse(file); err != nil {
		return err
	}
	_ = SkipUntil("hdlr", file)
	a.Hdlr = &Hdlr{}
	if err := a.Hdlr.Parse(file); err != nil {
		return err
	}
	return nil
}

func (a *Mdia) String() string {
	return fmt.Sprintf(
		"mdia: {Type: %s, Size: %d, Hdlr: {%s}, Mdhd: {%s}}",
		a.Type, a.Size, a.Hdlr, a.Mdhd,
	)
}

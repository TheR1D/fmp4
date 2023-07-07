package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Mdia struct {
	BaseAtom
	Hdlr *Hdlr
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
	if err := SkipUntil("hdlr", file); err != nil {
		return err
	}
	a.Hdlr = NewHdlr(file)
	return nil
}

func (a *Mdia) String() string {
	return fmt.Sprintf(
		"mdia: {Type: %s, Size: %d, Hdlr: {%s}}",
		a.Type, a.Size, a.Hdlr,
	)
}

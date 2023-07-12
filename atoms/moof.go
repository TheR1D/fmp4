package atoms

import (
	"fmt"
	"os"
)

type Moof struct {
	BaseAtom
	Traf *Traf
}

func NewMoof(file *os.File) *Moof {
	moof := &Moof{}
	if err := moof.Parse(file); err != nil {
		panic(err)
	}
	return moof
}

func (a *Moof) Parse(file *os.File) error {
	if err := a.BaseAtom.Parse(file); err != nil {
		return err
	}
	if err := SkipUntil("traf", file); err != nil {
		return err
	}
	a.Traf = &Traf{}
	if err := a.Traf.Parse(file); err != nil {
		return err
	}
	return nil
}

func (a *Moof) String() string {
	return fmt.Sprintf(
		"moof: {Type: %s, Size: %d, Traf: {%s}}",
		a.Type, a.Size, a.Traf,
	)
}

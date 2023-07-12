package atoms

import (
	"fmt"
	"os"
)

type Traf struct {
	BaseAtom
	Tfhd *Tfhd
}

func NewTraf(file *os.File) *Traf {
	traf := &Traf{}
	if err := traf.Parse(file); err != nil {
		panic(err)
	}
	return traf
}

func (a *Traf) Parse(file *os.File) error {
	if err := a.BaseAtom.Parse(file); err != nil {
		return err
	}
	if err := SkipUntil("tfhd", file); err != nil {
		return err
	}
	a.Tfhd = &Tfhd{}
	if err := a.Tfhd.Parse(file); err != nil {
		return err
	}
	return nil
}

func (a *Traf) String() string {
	return fmt.Sprintf(
		"traf: {Type: %s, Size: %d, Tfhd: {%s}}",
		a.Type, a.Size, a.Tfhd,
	)
}

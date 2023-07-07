package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Ftyp struct {
	BaseAtom
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands []byte
}

func NewFtyp(file *os.File) *Ftyp {
	ftypAtom := &Ftyp{}
	if err := ftypAtom.Parse(file); err != nil {
		panic(err)
	}
	return ftypAtom
}

func (a *Ftyp) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, &a.Size); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.Type); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.MajorBrand); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.MinorVersion); err != nil {
		return err
	}

	a.CompatibleBrands = make([]byte, a.Size-16)
	if _, err := file.Read(a.CompatibleBrands); err != nil {
		return err
	}
	return nil
}

func (a *Ftyp) String() string {
	return fmt.Sprintf(
		"ftyp: {Type: %s, Size: %d, MajorBrand: %s, MinorVersion: %d, CompatibleBrands: %s}",
		a.GetType(), a.GetSize(), a.MajorBrand, a.MinorVersion, a.CompatibleBrands,
	)
}

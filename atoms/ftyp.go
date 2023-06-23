package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Ftyp struct {
	// File type atom (aka ftyp) are always first in the file.
	MajorBrand       [4]byte
	MinorVersion     uint32
	CompatibleBrands []byte
	*Atom
}

func NewFtyp(atom *Atom, file *os.File) *Ftyp {
	fileType := Ftyp{Atom: atom}
	err := binary.Read(file, binary.BigEndian, &fileType.MajorBrand)
	if err != nil {
		panic(err)
	}
	err = binary.Read(file, binary.BigEndian, &fileType.MinorVersion)
	if err != nil {
		panic(err)
	}
	fileType.CompatibleBrands = make([]byte, atom.Size-16)
	_, err = file.Read(fileType.CompatibleBrands)
	if err != nil {
		panic(err)
	}
	return &fileType
}

func (a Ftyp) String() string {
	return fmt.Sprintf(
		"FileTypeAtom {"+
			"Type: %s, "+
			"Size: %d, "+
			"MajorBrand: %s, "+
			"MinorVersion: %d, "+
			"CompatibleBrands: %s"+
			"}",
		a.Type, a.Size, a.MajorBrand, a.MinorVersion, a.CompatibleBrands,
	)
}

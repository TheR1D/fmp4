package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Mdhd struct {
	FullAtom
	CreationTime     uint32
	ModificationTime uint32
	Timescale        uint32
	Duration         uint32
	Language         uint16
	Reserved         uint16
}

func NewMdhd(file *os.File) *Mdhd {
	mdhd := &Mdhd{}
	if err := mdhd.Parse(file); err != nil {
		panic(err)
	}
	return mdhd
}

func (a *Mdhd) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, a); err != nil {
		return err
	}
	return nil
}

func (a *Mdhd) GetLanguage() string {
	var lang [3]uint16
	lang[0] = (a.Language >> 10) & 0x1F
	lang[1] = (a.Language >> 5) & 0x1F
	lang[2] = (a.Language) & 0x1F
	return fmt.Sprintf("%s%s%s",
		string(rune(lang[0]+0x60)),
		string(rune(lang[1]+0x60)),
		string(rune(lang[2]+0x60)))
}

func (a *Mdhd) String() string {
	return fmt.Sprintf(
		"mdhd: {Type: %s, Size: %d, Version: %d, Flags: %b, CreationTime: %d, ModificationTime: %d, TimeScale: %d, Duration: %d, Language: %s, Reserved: %d}",
		a.GetType(), a.GetSize(), a.Version, a.Flags, a.CreationTime, a.ModificationTime, a.Timescale, a.Duration, a.GetLanguage(), a.Reserved,
	)
}

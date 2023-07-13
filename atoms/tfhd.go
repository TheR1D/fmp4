package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Tfhd struct {
	FullAtom
	TrackId uint32
	// Optional fields
	//BaseDataOffset         uint64
	//SampleDescriptionIndex uint32
	DefaultSampleDuration uint32
	DefaultSampleSize     uint32
	DefaultSampleFlags    uint32
}

func NewTfhd(file *os.File) *Tfhd {
	tfhd := &Tfhd{}
	if err := tfhd.Parse(file); err != nil {
		panic(err)
	}
	return tfhd
}

func (a *Tfhd) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, a); err != nil {
		return err
	}
	return nil
}

//func (a *Tfhd) String() string {
//	return fmt.Sprintf(
//		"tfhd: {Type: %s, Size: %d, Version: %d, Flags: %b, TrackId: %d, BaseDataOffset: %d, SampleDescriptionIndex: %d, DefaultSampleDuration: %d, DefaultSampleSize: %d, DefaultSampleFlags: %d}",
//		a.Type, a.Size, a.Version, a.Flags, a.TrackId, a.BaseDataOffset, a.SampleDescriptionIndex, a.DefaultSampleDuration, a.DefaultSampleSize, a.DefaultSampleFlags,
//	)
//}

func (a *Tfhd) String() string {
	return fmt.Sprintf(
		"tfhd: {Type: %s, Size: %d, Version: %d, Flags: %b, TrackId: %d, DefaultSampleDuration: %d, DefaultSampleSize: %d, DefaultSampleFlags: %d}",
		a.Type, a.Size, a.Version, a.Flags, a.TrackId, a.DefaultSampleDuration, a.DefaultSampleSize, a.DefaultSampleFlags,
	)
}

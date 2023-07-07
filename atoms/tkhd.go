package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Tkhd struct {
	BaseAtom
	Version          uint8
	Flags            [3]byte
	CreationTime     uint32
	ModificationTime uint32
	TrackID          uint32
	Reserved1        [4]byte
	Duration         uint32
	Reserved2        [8]byte
	Layer            uint16
	AlternateGroup   uint16
	Volume           [2]byte
	Reserved3        [2]byte
	Matrix           [36]byte
	Width            uint32 // Fixed-point 16.16
	Height           uint32 // Fixed-point 16.16
}

func NewTkhd(file *os.File) *Tkhd {
	tkhd := &Tkhd{}
	if err := tkhd.Parse(file); err != nil {
		panic(err)
	}
	return tkhd
}

func (a *Tkhd) Parse(file *os.File) error {
	if err := binary.Read(file, binary.BigEndian, a); err != nil {
		return err
	}
	return nil
}

func (a *Tkhd) GetWidth() uint32 {
	return a.Width / (1 << 16)
}

func (a *Tkhd) GetHeight() uint32 {
	return a.Height / (1 << 16)
}

func (a *Tkhd) String() string {
	return fmt.Sprintf(
		"tkhd: {Type: %s, Size: %d, Version: %d, Flags: %b, CreationTime: %d, ModificationTime: %d, TrackID: %d, Duration: %d, Layer: %d, AlternateGroup: %d, Volume: %d, Matrix: %b, Width: %d, Height: %d}",
		a.GetType(), a.GetSize(), a.Version, a.Flags, a.CreationTime, a.ModificationTime, a.TrackID, a.Duration, a.Layer, a.AlternateGroup, a.Volume, a.Matrix, a.GetWidth(), a.GetHeight(),
	)
}

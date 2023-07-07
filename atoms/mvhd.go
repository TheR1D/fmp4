package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Mvhd struct {
	BaseAtom
	Version           uint8
	Flags             [3]byte
	CreationTime      uint32
	ModificationTime  uint32
	TimeScale         uint32
	Duration          uint32
	PreferredRate     uint32
	PreferredVolume   [2]byte
	Reserved          [10]byte
	MatrixStructure   [36]byte
	PreviewTime       uint32
	PreviewDuration   uint32
	PosterTime        uint32
	SelectionTime     uint32
	SelectionDuration uint32
	CurrentTime       uint32
	NextTrackId       uint32
}

func NewMvhd(file *os.File) *Mvhd {
	mvhd := &Mvhd{}
	if err := mvhd.Parse(file); err != nil {
		panic(err)
	}
	return mvhd
}

func (a *Mvhd) Parse(file *os.File) error {
	return binary.Read(file, binary.BigEndian, a)
}

func (a *Mvhd) String() string {
	return fmt.Sprintf(
		"mvhd: {Type: %s, Size: %d, Version: %d, Flags: %b, CreationTime: %d, ModificationTime: %d, TimeScale: %d, Duration: %d, PreferredRate: %d, PreferredVolume: %d, MatrixStructure: %b, NextTrackId: %d}",
		a.GetType(), a.GetSize(), a.Version, a.Flags, a.CreationTime, a.ModificationTime, a.TimeScale, a.Duration, a.PreferredRate, a.PreferredVolume, a.MatrixStructure, a.NextTrackId,
	)
}

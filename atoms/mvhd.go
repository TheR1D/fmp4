package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Mvhd struct {
	Version          uint8
	Flags            [3]uint8
	CreationTime     uint32
	ModificationTime uint32
	TimeScale        uint32
	// The timescale is the number of time units that pass per second.
	// For instance, if the timescale is 600, this means that one second
	// is represented as 600 time units. So, if the Duration field was,
	// for example, 1200, this would correspond to a duration of 2 seconds.
	Duration        uint32
	PreferredRate   uint32
	PreferredVolume [2]byte
	// Ten bytes reserved for use by Apple.
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

func NewMvhd(atom *Atom, file *os.File) *Mvhd {
	movieHeader := Mvhd{}
	_ = binary.Read(file, binary.BigEndian, &movieHeader)
	return &movieHeader
}

func (a Mvhd) String() string {
	return fmt.Sprintf(
		"MovieHeader {"+
			"Version: %v, "+
			"Flags: %v, "+
			"CreationTime: %v, "+
			"ModificationTime: %v, "+
			"TimeScale: %v, "+
			"Duration: %v, "+
			"PreferredRate: %v, "+
			"PreferredVolume: %v, "+
			"Reserved: %v, "+
			"MatrixStructure: %v, "+
			"PreviewTime: %v, "+
			"PreviewDuration: %v, "+
			"PosterTime: %v, "+
			"SelectionTime: %v, "+
			"SelectionDuration: %v, "+
			"CurrentTime: %v, "+
			"NextTrackId: %v"+
			"}",
		a.Version,
		a.Flags,
		a.CreationTime,
		a.ModificationTime,
		a.TimeScale,
		a.Duration,
		a.PreferredRate,
		a.PreferredVolume,
		a.Reserved,
		a.MatrixStructure,
		a.PreviewTime,
		a.PreviewDuration,
		a.PosterTime,
		a.SelectionTime,
		a.SelectionDuration,
		a.CurrentTime,
		a.NextTrackId,
	)
}

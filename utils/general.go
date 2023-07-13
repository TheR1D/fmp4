package utils

import (
	"fmp4/atoms"
	"fmt"
	"math"
	"os"
)

func GetFramerate(file *os.File) uint32 {
	atoms.MustSkipUntil("moov", file)
	moov := atoms.NewMoov(file)
	atoms.MustSkipUntil("moof", file)
	moof := atoms.NewMoof(file)
	timescale := moov.Trak.Mdia.Mdhd.Timescale
	duration := moof.Traf.Tfhd.DefaultSampleDuration
	secondsPerFrame := float32(duration) / float32(timescale)
	frameRate := 1 / secondsPerFrame
	fmt.Println(timescale, duration)
	// TODO: Check if there is a better way to calculate frame rate.
	return uint32(math.Round(float64(frameRate)))
}

func TotalSizeOfAtoms(atoms []*atoms.SAtom) uint64 {
	var totalSize uint64
	for _, atom := range atoms {
		totalSize += uint64(atom.Atom.GetSize())
	}
	return totalSize
}

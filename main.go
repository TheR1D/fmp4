package main

import (
	"fmp4/atoms"
	"fmt"
	"os"
)

const (
	segmentLength  = 6 // Segment length in seconds for manifest.
	mdatLength     = 2 // Single mdat length in seconds.
	segmentsLength = (segmentLength / mdatLength) * 2
)

//func getFrameRate(timescale uint32, fragmentDuration uint32) float32 {
//	return float32(timescale) / float32(fragmentDuration)
//}
//
//func DetectFrameRate(file *os.File) float32 {
//	if err := atoms.SkipUntil("mvhd", file); err != nil {
//		panic(err)
//	}
//	mvhd := atoms.NewMvhd(file)
//	if err := atoms.SkipUntil("moof", file); err != nil {
//		panic(err)
//	}
//	moof := atoms.NewMoof(file)
//}

func getFramerate(timescale uint32, duration uint32) uint32 {
	secondsPerFrame := float32(duration) / float32(timescale)
	frameRate := 1 / secondsPerFrame
	return uint32(frameRate)
}

func main() {
	// Expecting fragmented mp4 file FullHD 60fps, 6 seconds per manifest fragment.
	// TODO: Define manifest bandwidth, framerate, resolution based on mp4 metadata.
	// TODO: Add support for multiple video/audio/subtitle tracks.
	fileName := "main.mp4"
	file, err := os.Open("static/" + fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for iter := atoms.NewAtomIterator(file, true); iter.Next(); {
		satom := iter.Value()
		if moov, ok := satom.Atom.(*atoms.Moov); ok {
			fmt.Println("moov found", moov)
		}
		fmt.Println(satom)
	}
}

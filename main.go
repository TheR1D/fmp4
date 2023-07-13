package main

import (
	"fmp4/atoms"
	"fmp4/manifests"
	"fmp4/utils"
	"fmt"
	"io"
	"os"
)

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

	segmentLength := 6 // Segment length in seconds for manifest.
	mdatLength := 2    // Single mdat length in seconds.
	segmentsLength := (segmentLength / mdatLength) * 2

	file.Seek(0, io.SeekStart)
	var manifest *manifests.Hls

	mainBr := utils.ByteRange{}
	segmentAtoms := make([]*atoms.SAtom, segmentsLength)
	counter := 0
	manifestVidPath := "/files/" + fileName

	for iter := atoms.NewAtomIterator(file, true); iter.Next(); {
		satom := iter.Value()
		fmt.Println(satom)
		atom := satom.Atom
		// TODO: Could be optimized.
		if atom.GetType() == "moov" {
			// If we found first moov atom, the previous bytes are
			// main byte range which include ftype, moov atoms.
			mainBr.Length = uint64(satom.EndsAt())
			manifest = manifests.NewHls(segmentLength, manifestVidPath, mainBr)
		} else if atom.GetType() == "moof" || atom.GetType() == "mdat" {
			segmentAtoms[counter] = satom
			counter++
			if counter == segmentsLength {
				segmentStart := segmentAtoms[0].StartsAt()
				totalSize := utils.TotalSizeOfAtoms(segmentAtoms)
				segmentBr := utils.ByteRange{Start: uint64(segmentStart), Length: totalSize}
				manifest.AppendSegment(float64(segmentLength), segmentBr, manifestVidPath)
				counter = 0
			}
		}
	}
	manifest.Finalize()
	err = manifest.WriteToFile("static/manifest_fhd.m3u8")
	if err != nil {
		fmt.Println("Couldn't write to file:", err)
	}
}

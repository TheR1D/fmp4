package atoms

import (
	"io"
	"os"
)

type Iterator interface {
	Next() bool
	Value() Atom
}

type AtomIterator struct {
	file      *os.File
	parseBody bool
}

func NewAtomIterator(file *os.File, parseBody bool) *AtomIterator {
	return &AtomIterator{file: file, parseBody: parseBody}
}

func (it *AtomIterator) Next() bool {
	// Get the file size
	fileInfo, err := it.file.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size()
	currSeek, err := it.file.Seek(0, io.SeekCurrent)
	// If we can't read next 8 bytes (type + size), we are at the end of the file.
	return err == nil && currSeek+8 < fileSize
}

func (it *AtomIterator) Value() *SAtom {
	satom := &SAtom{Atom: NewBaseAtom(it.file)}
	// Since we read 8 bytes for base atom, we need to seek back.
	// Ignoring error, since we are sure that we can seek back.
	satom.startsAt, _ = it.file.Seek(-8, io.SeekCurrent)
	if it.parseBody {
		switch satom.Atom.GetType() {
		case "ftyp":
			satom.Atom = NewFtyp(it.file)
			return satom
		case "moov":
			satom.Atom = NewMoov(it.file)
			if _, err := it.file.Seek(satom.EndsAt(), io.SeekStart); err != nil {
				panic(err)
			}
			return satom
		case "moof":
			satom.Atom = NewMoof(it.file)
			if _, err := it.file.Seek(satom.EndsAt(), io.SeekStart); err != nil {
				panic(err)
			}
			return satom
		}
	}
	size := int64(satom.Atom.GetSize())
	if _, err := it.file.Seek(size, io.SeekCurrent); err != nil {
		panic(err)
	}
	return satom
}

package atoms

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type Atom struct {
	// If the size is set to 1, it is 64-bit and stored in the next 8 bytes.
	Size uint32 // TODO: Might be large input uint64.
	Type [4]byte
}

func NewAtom(file *os.File) (*Atom, error) {
	atom := Atom{}
	err := binary.Read(file, binary.BigEndian, &atom)
	if err != nil {
		return nil, err
	}
	return &atom, nil
}

func (a Atom) TypeStr() string {
	return string(a.Type[:])
}

func (a Atom) String() string {
	return fmt.Sprintf("Atom: {Type: %s, Size: %d}", a.Type, a.Size)
}

type AtomWithSeek struct {
	Atom *Atom
	Seek int64
}

func (a AtomWithSeek) String() string {
	return fmt.Sprintf(
		"AtomWithSeek: {Type: %s, Size: %d, Seek: %d}", a.Atom.Type, a.Atom.Size, a.Seek,
	)
}

func Generator(file *os.File) <-chan *AtomWithSeek {
	ch := make(chan *AtomWithSeek)

	go func() {
		defer close(ch)
		for {
			atom, err := NewAtom(file)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
			}
			// Skip to the next atom.
			_, err = file.Seek(int64(atom.Size-8), io.SeekCurrent)
			if err != nil {
				// If we can't seek, we can't continue, since we
				// don't know how much to skip to get to the next atom.
				panic(err)
			}
			// Should not throw error, since we are just getting the current seek.
			curSeek, _ := file.Seek(0, io.SeekCurrent)
			fmt.Println(atom, "Current seek:", curSeek)
			ch <- &AtomWithSeek{
				Atom: atom,
				Seek: curSeek,
			}
		}
	}()

	return ch
}

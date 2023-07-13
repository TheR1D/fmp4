package atoms

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Hdlr struct {
	FullAtom
	Reserved1   [4]byte
	HandlerType [4]byte // Should be cast to string.
	Reserved2   [12]byte
	HandlerName []byte // Should be cast to string.
}

func NewHdlr(file *os.File) *Hdlr {
	hdlr := &Hdlr{}
	if err := hdlr.Parse(file); err != nil {
		panic(err)
	}
	return hdlr
}

func (a *Hdlr) Parse(file *os.File) error {
	if err := a.FullAtom.Parse(file); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.Reserved1); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.HandlerType); err != nil {
		return err
	}
	if err := binary.Read(file, binary.BigEndian, &a.Reserved2); err != nil {
		return err
	}
	// Subtract 32 bytes, since at this point we read 32 bytes for fields above.
	a.HandlerName = make([]byte, a.Size-32)
	if err := binary.Read(file, binary.BigEndian, &a.HandlerName); err != nil {
		return err
	}
	return nil
}

func (a *Hdlr) String() string {
	return fmt.Sprintf(
		"hdlr: {Type: %s, Size: %d, Version: %d, Flags: %b, HandlerType: %s, HandlerName: %s}",
		a.GetType(), a.GetSize(), a.Version, a.Flags, a.HandlerType, a.HandlerName,
	)
}

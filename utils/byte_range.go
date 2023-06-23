package utils

type ByteRange struct {
	Start  uint64
	Length uint64
}

func (br ByteRange) End() uint64 {
	return br.Start + br.Length
}

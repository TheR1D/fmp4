package atoms

type Tfhd struct {
	BaseAtom
	Version        uint8
	Flags          [3]byte
	TrackId        uint32
	BaseDataOffset uint64
}

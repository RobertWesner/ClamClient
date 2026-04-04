package packets

type Packet interface {
	ID() uint8
	Bytes() ([]byte, error)
	Read(reader PacketReader) error
}

package packets

type Packet interface {
	Id() uint8
	Bytes() ([]byte, error)
	Read(reader PacketReader) error
}

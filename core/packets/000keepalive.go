package packets

type Packet0KeepAlive struct{}

func (p Packet0KeepAlive) Id() uint8 {
	return 0x00
}

func (p Packet0KeepAlive) Bytes() ([]byte, error) {
	return []byte{}, nil
}

func (p Packet0KeepAlive) Read(reader PacketReader) error {
	return nil
}

func NewPacket0KeepAlive() Packet0KeepAlive {
	return Packet0KeepAlive{}
}

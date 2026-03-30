package packets

import (
	"fmt"
)

type Packet255Disconnect struct {
	Reason string
}

func (p Packet255Disconnect) Id() uint8 {
	return 0xFF
}

func (p Packet255Disconnect) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.WriteString16(p.Reason); err != nil {
		return nil, fmt.Errorf("255 write reason: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet255Disconnect) Read(reader PacketReader) error {
	var err error

	if p.Reason, err = reader.String16(); err != nil {
		return fmt.Errorf("255 read reason: %w", err)
	}

	return nil
}

func NewPacket255Disconnect() Packet255Disconnect {
	return Packet255Disconnect{
		Reason: "Clam out! o7",
	}
}

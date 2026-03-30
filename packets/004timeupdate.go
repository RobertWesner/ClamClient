package packets

import (
	"fmt"
)

type Packet4TimeUpdate struct {
	Time int64
}

func (p Packet4TimeUpdate) Id() uint8 {
	return 0x04
}

func (p Packet4TimeUpdate) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.Write(p.Time)
	if err != nil {
		return nil, fmt.Errorf("006 write time: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet4TimeUpdate) Read(reader PacketReader) error {
	time, err := reader.Int64()
	if err != nil {
		return fmt.Errorf("006 read x: %w", err)
	}

	p.Time = time

	return nil
}

package packets

import (
	"errors"
	"fmt"
)

type Packet4TimeUpdate struct {
	Time int64
}

func (p Packet4TimeUpdate) ID() uint8 {
	return 0x04
}

func (p Packet4TimeUpdate) Bytes() ([]byte, error) {
	return []byte{}, errors.New("004 server->client packets should never be sent")
}

func (p Packet4TimeUpdate) Read(reader PacketReader) error {
	var err error

	if p.Time, err = reader.Int64(); err != nil {
		return fmt.Errorf("006 read x: %w", err)
	}

	return nil
}

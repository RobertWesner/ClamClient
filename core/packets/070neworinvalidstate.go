package packets

import (
	"errors"
	"fmt"
)

type Packet70NewOrInvalidState struct {
	Reason uint8
}

func (p Packet70NewOrInvalidState) ID() uint8 {
	return 0x46
}

func (p Packet70NewOrInvalidState) Bytes() ([]byte, error) {
	return []byte{}, errors.New("070 server->client packets should never be sent")
}

func (p Packet70NewOrInvalidState) Read(reader PacketReader) error {
	var err error

	if p.Reason, err = reader.Byte(); err != nil {
		return fmt.Errorf("070 read reason: %w", err)
	}

	return nil
}

package packets

import (
	"errors"
	"fmt"
)

type Packet31EntityRelativeMove struct {
	EntityID int32
	DX       uint8
	DY       uint8
	DZ       uint8
}

func (p Packet31EntityRelativeMove) ID() uint8 {
	return 0x1F
}

func (p Packet31EntityRelativeMove) Bytes() ([]byte, error) {
	return []byte{}, errors.New("031 server->client packets should never be sent")
}

func (p Packet31EntityRelativeMove) Read(reader PacketReader) error {
	var err error

	if p.EntityID, err = reader.Int32(); err != nil {
		return fmt.Errorf("031 read entityid: %w", err)
	}

	if p.DX, err = reader.Byte(); err != nil {
		return fmt.Errorf("031 read dx: %w", err)
	}

	if p.DY, err = reader.Byte(); err != nil {
		return fmt.Errorf("031 read dy: %w", err)
	}

	if p.DZ, err = reader.Byte(); err != nil {
		return fmt.Errorf("031 read dz: %w", err)
	}

	return nil
}

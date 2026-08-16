package packets

import (
	"errors"
	"fmt"
)

type Packet33LookAndRelativeMove struct {
	EntityID int32
	DX       uint8
	DY       uint8
	DZ       uint8
	Yaw      Uint8angle // 64 = 90°
	Pitch    Uint8angle // 64 = 90°
}

func (p *Packet33LookAndRelativeMove) ID() uint8 {
	return 0x21
}

func (p *Packet33LookAndRelativeMove) Bytes() ([]byte, error) {
	return []byte{}, errors.New("033 server->client packets should never be sent")
}

func (p *Packet33LookAndRelativeMove) Read(reader PacketReader) error {
	var err error

	if p.EntityID, err = reader.Int32(); err != nil {
		return fmt.Errorf("033 read entityid: %w", err)
	}

	if p.DX, err = reader.Byte(); err != nil {
		return fmt.Errorf("033 read dx: %w", err)
	}

	if p.DY, err = reader.Byte(); err != nil {
		return fmt.Errorf("033 read dy: %w", err)
	}

	if p.DZ, err = reader.Byte(); err != nil {
		return fmt.Errorf("033 read dz: %w", err)
	}

	yaw, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read yaw: %w", err)
	}
	p.Yaw = Uint8angle(yaw)

	pitch, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read pitch: %w", err)
	}
	p.Pitch = Uint8angle(pitch)

	return nil
}

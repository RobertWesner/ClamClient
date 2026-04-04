package packets

import (
	"errors"
	"fmt"
)

type Packet32EntityLook struct {
	EntityID int32
	Yaw      Uint8angle // 64 = 90°
	Pitch    Uint8angle // 64 = 90°
}

func (p Packet32EntityLook) ID() uint8 {
	return 0x20
}

func (p Packet32EntityLook) Bytes() ([]byte, error) {
	return []byte{}, errors.New("032 server->client packets should never be sent")
}

func (p Packet32EntityLook) Read(reader PacketReader) error {
	var err error

	if p.EntityID, err = reader.Int32(); err != nil {
		return fmt.Errorf("032 read entityid: %w", err)
	}

	yaw, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("032 read yaw: %w", err)
	}
	p.Yaw = Uint8angle(yaw)

	pitch, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("032 read pitch: %w", err)
	}
	p.Pitch = Uint8angle(pitch)

	return nil
}

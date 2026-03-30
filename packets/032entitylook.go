package packets

import (
	"fmt"
)

type Packet32EntityLook struct {
	EntityID int32
	Yaw      Uin8angle // 64 = 90°
	Pitch    Uin8angle // 64 = 90°
}

func (p Packet32EntityLook) Id() uint8 {
	return 0x20
}

func (p Packet32EntityLook) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.Write(p.EntityID)
	if err != nil {
		return nil, fmt.Errorf("032 write entityid: %w", err)
	}

	err = writer.Write(p.Yaw)
	if err != nil {
		return nil, fmt.Errorf("032 write yaw: %w", err)
	}

	err = writer.Write(p.Pitch)
	if err != nil {
		return nil, fmt.Errorf("032 write pitch: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet32EntityLook) Read(reader PacketReader) error {
	entityId, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("032 read entityid: %w", err)
	}

	yaw, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("032 read yaw: %w", err)
	}

	pitch, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("032 read pitch: %w", err)
	}

	p.EntityID = entityId
	p.Yaw = Uin8angle(yaw)
	p.Pitch = Uin8angle(pitch)

	return nil
}

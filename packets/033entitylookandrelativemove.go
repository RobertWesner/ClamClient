package packets

import (
	"fmt"
)

type Packet33LookAndRelativeMove struct {
	EntityID int32
	DX       uint8
	DY       uint8
	DZ       uint8
	Yaw      Uin8angle // 64 = 90°
	Pitch    Uin8angle // 64 = 90°
}

func (p Packet33LookAndRelativeMove) Id() uint8 {
	return 0x21
}

func (p Packet33LookAndRelativeMove) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.Write(p.EntityID)
	if err != nil {
		return nil, fmt.Errorf("033 write entityid: %w", err)
	}

	err = writer.Write(p.DX)
	if err != nil {
		return nil, fmt.Errorf("033 write dx: %w", err)
	}

	err = writer.Write(p.DY)
	if err != nil {
		return nil, fmt.Errorf("033 write dy: %w", err)
	}

	err = writer.Write(p.DZ)
	if err != nil {
		return nil, fmt.Errorf("033 write dz: %w", err)
	}

	err = writer.Write(p.Yaw)
	if err != nil {
		return nil, fmt.Errorf("033 write yaw: %w", err)
	}

	err = writer.Write(p.Pitch)
	if err != nil {
		return nil, fmt.Errorf("033 write pitch: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet33LookAndRelativeMove) Read(reader PacketReader) error {
	entityId, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("033 read entityid: %w", err)
	}

	dx, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read dx: %w", err)
	}

	dy, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read dy: %w", err)
	}

	dz, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read dz: %w", err)
	}

	yaw, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read yaw: %w", err)
	}

	pitch, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("033 read pitch: %w", err)
	}

	p.EntityID = entityId
	p.Yaw = Uin8angle(yaw)
	p.Pitch = Uin8angle(pitch)
	p.DX = dx
	p.DY = dy
	p.DZ = dz

	return nil
}

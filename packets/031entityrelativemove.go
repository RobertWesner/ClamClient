package packets

import (
	"fmt"
)

type Packet31EntityRelativeMove struct {
	EntityID int32
	DX       uint8
	DY       uint8
	DZ       uint8
}

func (p Packet31EntityRelativeMove) Id() uint8 {
	return 0x1F
}

func (p Packet31EntityRelativeMove) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.Write(p.EntityID)
	if err != nil {
		return nil, fmt.Errorf("031 write entityid: %w", err)
	}

	err = writer.Write(p.DX)
	if err != nil {
		return nil, fmt.Errorf("031 write dx: %w", err)
	}

	err = writer.Write(p.DY)
	if err != nil {
		return nil, fmt.Errorf("031 write dy: %w", err)
	}

	err = writer.Write(p.DZ)
	if err != nil {
		return nil, fmt.Errorf("031 write dz: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet31EntityRelativeMove) Read(reader PacketReader) error {
	entityId, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("031 read entityid: %w", err)
	}

	dx, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("031 read dx: %w", err)
	}

	dy, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("031 read dy: %w", err)
	}

	dz, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("031 read dz: %w", err)
	}

	p.EntityID = entityId
	p.DX = dx
	p.DY = dy
	p.DZ = dz

	return nil
}

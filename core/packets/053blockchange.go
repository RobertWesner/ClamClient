package packets

import (
	"fmt"
)

type Packet53BlockChange struct {
	X             int32
	Y             uint8
	Z             int32
	BlockType     uint8
	BlockMetadata uint8
}

func (p *Packet53BlockChange) ID() uint8 {
	return 0x35
}

func (p *Packet53BlockChange) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.Write(p.X); err != nil {
		return nil, fmt.Errorf("053 write x: %w", err)
	}

	if err = writer.Write(p.Y); err != nil {
		return nil, fmt.Errorf("053 write y: %w", err)
	}

	if err = writer.Write(p.Z); err != nil {
		return nil, fmt.Errorf("053 write z: %w", err)
	}

	if err = writer.Write(p.BlockType); err != nil {
		return nil, fmt.Errorf("053 write block type: %w", err)
	}

	if err = writer.Write(p.BlockMetadata); err != nil {
		return nil, fmt.Errorf("053 write block metadata: %w", err)
	}

	return writer.Bytes(), nil
}

func (p *Packet53BlockChange) Read(reader PacketReader) error {
	var err error

	if p.X, err = reader.Int32(); err != nil {
		return fmt.Errorf("053 read x: %w", err)
	}

	if p.Y, err = reader.Byte(); err != nil {
		return fmt.Errorf("053 read y: %w", err)
	}

	if p.Z, err = reader.Int32(); err != nil {
		return fmt.Errorf("053 read z: %w", err)
	}

	if p.BlockType, err = reader.Byte(); err != nil {
		return fmt.Errorf("053 read block type: %w", err)
	}

	if p.BlockMetadata, err = reader.Byte(); err != nil {
		return fmt.Errorf("053 read block metadata: %w", err)
	}

	return nil
}

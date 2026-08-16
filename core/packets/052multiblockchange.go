package packets

import (
	"fmt"
)

type Packet52MultiBlockChange struct {
	ChunkX          int32
	ChunkZ          int32
	ArraySize       int16
	CoordinateArray []int16
	TypeArray       []uint8
	MetadataArray   []uint8
}

func (p Packet52MultiBlockChange) ID() uint8 {
	return 0x34
}

func (p Packet52MultiBlockChange) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.Write(p.ChunkX); err != nil {
		return nil, fmt.Errorf("054 write chunk x: %w", err)
	}

	if err = writer.Write(p.ChunkZ); err != nil {
		return nil, fmt.Errorf("054 write chunk z: %w", err)
	}

	if err = writer.Write(p.ArraySize); err != nil {
		return nil, fmt.Errorf("054 write size: %w", err)
	}

	for it := range p.CoordinateArray {
		if err = writer.Write(it); err != nil {
			return nil, fmt.Errorf("054 write coords: %w", err)
		}
	}

	for it := range p.TypeArray {
		if err = writer.Write(it); err != nil {
			return nil, fmt.Errorf("054 write types: %w", err)
		}
	}

	for it := range p.MetadataArray {
		if err = writer.Write(it); err != nil {
			return nil, fmt.Errorf("054 write metadata: %w", err)
		}
	}

	return writer.Bytes(), nil
}

func (p Packet52MultiBlockChange) Read(reader PacketReader) error {
	var err error

	if p.ChunkX, err = reader.Int32(); err != nil {
		return fmt.Errorf("054 read chunk x: %w", err)
	}

	if p.ChunkZ, err = reader.Int32(); err != nil {
		return fmt.Errorf("054 read chunk z: %w", err)
	}

	if p.ArraySize, err = reader.Int16(); err != nil {
		return fmt.Errorf("054 read array size: %w", err)
	}

	for i := int16(0); i < p.ArraySize; i++ {
		var it int16
		if it, err = reader.Int16(); err != nil {
			return fmt.Errorf("054 read coords: %w", err)
		}

		p.CoordinateArray = append(p.CoordinateArray, it)
	}

	for i := int16(0); i < p.ArraySize; i++ {
		var it uint8
		if it, err = reader.Byte(); err != nil {
			return fmt.Errorf("054 read types: %w", err)
		}

		p.TypeArray = append(p.TypeArray, it)
	}

	for i := int16(0); i < p.ArraySize; i++ {
		var it uint8
		if it, err = reader.Byte(); err != nil {
			return fmt.Errorf("054 read metadata: %w", err)
		}

		p.MetadataArray = append(p.MetadataArray, it)
	}

	return nil
}

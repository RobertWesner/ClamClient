package packets

import (
	"fmt"
)

type Packet6SpawnPosition struct {
	X int32
	Y int32
	Z int32
}

func (p Packet6SpawnPosition) Id() uint8 {
	return 0x06
}

func (p Packet6SpawnPosition) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.Write(p.X)
	if err != nil {
		return nil, fmt.Errorf("006 write x: %w", err)
	}

	err = writer.Write(p.Y)
	if err != nil {
		return nil, fmt.Errorf("006 write y: %w", err)
	}

	err = writer.Write(p.Z)
	if err != nil {
		return nil, fmt.Errorf("006 write z: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet6SpawnPosition) Read(reader PacketReader) error {
	x, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("006 read x: %w", err)
	}

	y, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("006 read y: %w", err)
	}

	z, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("006 read z: %w", err)
	}

	p.X = x
	p.Y = y
	p.Z = z

	return nil
}

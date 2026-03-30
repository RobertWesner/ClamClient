package packets

import (
	"errors"
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
	return []byte{}, errors.New("006 server->client packets should never be sent")
}

func (p Packet6SpawnPosition) Read(reader PacketReader) error {
	var err error

	if p.X, err = reader.Int32(); err != nil {
		return fmt.Errorf("006 read x: %w", err)
	}

	if p.Y, err = reader.Int32(); err != nil {
		return fmt.Errorf("006 read y: %w", err)
	}

	if p.Z, err = reader.Int32(); err != nil {
		return fmt.Errorf("006 read z: %w", err)
	}

	return nil
}

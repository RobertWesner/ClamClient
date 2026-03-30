package packets

import (
	"errors"
	"fmt"
)

type Packet51MapChunk struct {
	X              int32
	Y              int16
	Z              int32
	SizeX          uint8 // SizeX is Actual X Size -1
	SizeY          uint8 // SizeY is Actual Y Size -1
	SizeZ          uint8 // SizeZ is Actual Zs Size -1
	CompressedSize int32
	CompressedData []byte
}

func (p Packet51MapChunk) Id() uint8 {
	return 0x33
}

func (p Packet51MapChunk) Bytes() ([]byte, error) {
	return []byte{}, errors.New("051 server->client packets should never be sent")
}

func (p Packet51MapChunk) Read(reader PacketReader) error {
	var err error

	if p.X, err = reader.Int32(); err != nil {
		return fmt.Errorf("051 read x: %w", err)
	}

	if p.Y, err = reader.Int16(); err != nil {
		return fmt.Errorf("051 read y: %w", err)
	}

	if p.Z, err = reader.Int32(); err != nil {
		return fmt.Errorf("051 read z: %w", err)
	}

	if p.SizeX, err = reader.Byte(); err != nil {
		return fmt.Errorf("051 read sx: %w", err)
	}

	if p.SizeY, err = reader.Byte(); err != nil {
		return fmt.Errorf("051 read sy: %w", err)
	}

	if p.SizeZ, err = reader.Byte(); err != nil {
		return fmt.Errorf("051 read sz: %w", err)
	}

	if p.CompressedSize, err = reader.Int32(); err != nil {
		return fmt.Errorf("051 read compSize: %w", err)
	}

	if p.CompressedData, err = reader.Bytes(int(p.CompressedSize)); err != nil {
		return fmt.Errorf("051 read compSize: %w", err)
	}

	return nil
}

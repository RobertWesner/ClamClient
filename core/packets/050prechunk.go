package packets

import (
	"errors"
	"fmt"
)

// Packet50PreChunk will be completely ignored
type Packet50PreChunk struct {
}

func (p Packet50PreChunk) Id() uint8 {
	return 0x32
}

func (p Packet50PreChunk) Bytes() ([]byte, error) {
	return []byte{}, errors.New("050 server->client packets should never be sent")
}

func (p Packet50PreChunk) Read(reader PacketReader) error {
	var err error

	if _, err = reader.Int32(); err != nil {
		return fmt.Errorf("050 read x: %w", err)
	}

	if _, err = reader.Int32(); err != nil {
		return fmt.Errorf("050 read z: %w", err)
	}

	if _, err = reader.Bool(); err != nil {
		return fmt.Errorf("050 read mode: %w", err)
	}

	return nil
}

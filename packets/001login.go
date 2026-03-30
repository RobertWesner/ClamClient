package packets

import (
	"fmt"
	"log/slog"
)

type Packet1Login struct {
	ProtocolVersionOrEntityId int32
	Username                  string
	MapSeed                   int64
	Dimension                 byte
}

func (p Packet1Login) Id() uint8 {
	return 0x01
}

func (p Packet1Login) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.Write(p.ProtocolVersionOrEntityId)
	if err != nil {
		return nil, fmt.Errorf("001 write protocol version: %w", err)
	}

	err = writer.WriteString16(p.Username)
	if err != nil {
		return nil, fmt.Errorf("001 write username: %w", err)
	}

	err = writer.Write(p.MapSeed)
	if err != nil {
		return nil, fmt.Errorf("001 write map seed: %w", err)
	}

	err = writer.Write(p.Dimension)
	if err != nil {
		return nil, fmt.Errorf("001 write dimension: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet1Login) Read(reader PacketReader) error {
	version, err := reader.Int32()
	if err != nil {
		return fmt.Errorf("001 read protocol version: %w", err)
	}

	username, err := reader.String16()
	if err != nil {
		return fmt.Errorf("001 read username: %w", err)
	}

	seed, err := reader.Int64()
	if err != nil {
		return fmt.Errorf("001 read username: %w", err)
	}

	dimension, err := reader.Byte()
	if err != nil {
		return fmt.Errorf("001 read username: %w", err)
	}

	p.ProtocolVersionOrEntityId = version
	p.Username = username
	p.MapSeed = seed
	p.Dimension = dimension

	return nil
}

func NewPacket1Login(
	protocolVersion int32,
	username string,
	mapSeed int64,
	dimension byte,
) Packet1Login {
	if len(username) > 16 {
		slog.Warn("001 username is too long, truncating", "username", username)
		username = username[:16]
	}

	return Packet1Login{
		ProtocolVersionOrEntityId: protocolVersion,
		Username:                  username,
		MapSeed:                   mapSeed,
		Dimension:                 dimension,
	}
}

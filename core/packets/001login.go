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
	var err error

	writer := NewWriter()

	if err = writer.Write(p.ProtocolVersionOrEntityId); err != nil {
		return nil, fmt.Errorf("001 write protocol version: %w", err)
	}

	if err = writer.WriteString16(p.Username); err != nil {
		return nil, fmt.Errorf("001 write username: %w", err)
	}

	if err = writer.Write(p.MapSeed); err != nil {
		return nil, fmt.Errorf("001 write map seed: %w", err)
	}

	if err = writer.Write(p.Dimension); err != nil {
		return nil, fmt.Errorf("001 write dimension: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet1Login) Read(reader PacketReader) error {
	var err error

	if p.ProtocolVersionOrEntityId, err = reader.Int32(); err != nil {
		return fmt.Errorf("001 read protocol version: %w", err)
	}

	if p.Username, err = reader.String16(); err != nil {
		return fmt.Errorf("001 read username: %w", err)
	}

	if p.MapSeed, err = reader.Int64(); err != nil {
		return fmt.Errorf("001 read username: %w", err)
	}

	if p.Dimension, err = reader.Byte(); err != nil {
		return fmt.Errorf("001 read username: %w", err)
	}

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

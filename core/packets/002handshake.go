package packets

import (
	"fmt"
	"log/slog"
)

type Packet2Handshake struct {
	UsernameOrConnectionHash string
}

func (p *Packet2Handshake) ID() uint8 {
	return 0x02
}

func (p *Packet2Handshake) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.WriteString16(p.UsernameOrConnectionHash); err != nil {
		return nil, fmt.Errorf("002 write username: %w", err)
	}

	return writer.Bytes(), nil
}

func (p *Packet2Handshake) Read(reader PacketReader) error {
	var err error

	if p.UsernameOrConnectionHash, err = reader.String16(); err != nil {
		return fmt.Errorf("002 read username: %w", err)
	}

	return nil
}

func NewPacket2Handshake(
	username string,
) *Packet2Handshake {
	if len(username) > 16 {
		slog.Warn("002 username is too long, truncating", "username", username)
		username = username[:16]
	}

	return &Packet2Handshake{
		UsernameOrConnectionHash: username,
	}
}

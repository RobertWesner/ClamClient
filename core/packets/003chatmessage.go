package packets

import (
	"fmt"
	"log/slog"
)

type Packet3ChatMessage struct {
	Message string
}

func (p Packet3ChatMessage) ID() uint8 {
	return 0x03
}

func (p Packet3ChatMessage) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.WriteString16(p.Message); err != nil {
		return nil, fmt.Errorf("003 write message: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet3ChatMessage) Read(reader PacketReader) error {
	var err error

	if p.Message, err = reader.String16(); err != nil {
		return fmt.Errorf("003 read message: %w", err)
	}

	return nil
}

func NewPacket3ChatMessage(
	message string,
) Packet3ChatMessage {
	if len(message) > 16 {
		slog.Warn("003 message is too long, truncating", "message", message)
		message = message[:100]
	}

	return Packet3ChatMessage{
		Message: message,
	}
}

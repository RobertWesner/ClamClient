package packets

import (
	"fmt"
	"log/slog"
)

type Packet3ChatMessage struct {
	Message string
}

func (p Packet3ChatMessage) Id() uint8 {
	return 0x02
}

func (p Packet3ChatMessage) Bytes() ([]byte, error) {
	writer := NewWriter()

	err := writer.WriteString16(p.Message)
	if err != nil {
		return nil, fmt.Errorf("003 write message: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet3ChatMessage) Read(reader PacketReader) error {
	message, err := reader.String16()
	if err != nil {
		return fmt.Errorf("003 read message: %w", err)
	}

	p.Message = message

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

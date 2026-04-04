package packets

import (
	"fmt"
)

type Packet106Transaction struct {
	WindowID     uint8
	ActionNumber int16
	Accepted     bool
}

func (p Packet106Transaction) ID() uint8 {
	return 0x6A
}

func (p Packet106Transaction) Bytes() ([]byte, error) {
	var err error

	writer := NewWriter()

	if err = writer.Write(p.WindowID); err != nil {
		return nil, fmt.Errorf("106 write windowid: %w", err)
	}

	if err = writer.Write(p.ActionNumber); err != nil {
		return nil, fmt.Errorf("106 write actionnumber: %w", err)
	}

	if err = writer.Write(p.Accepted); err != nil {
		return nil, fmt.Errorf("106 write accepted: %w", err)
	}

	return writer.Bytes(), nil
}

func (p Packet106Transaction) Read(reader PacketReader) error {
	var err error

	if p.WindowID, err = reader.Byte(); err != nil {
		return fmt.Errorf("106 read windowid: %w", err)
	}

	if p.ActionNumber, err = reader.Int16(); err != nil {
		return fmt.Errorf("106 read actionnumber: %w", err)
	}

	if p.Accepted, err = reader.Bool(); err != nil {
		return fmt.Errorf("106 read accepted: %w", err)
	}

	return nil
}

func NewPacket106Transaction(
	windowID uint8,
	actionNumber int16,
	accepted bool,
) Packet106Transaction {
	return Packet106Transaction{
		WindowID:     windowID,
		ActionNumber: actionNumber,
		Accepted:     accepted,
	}
}

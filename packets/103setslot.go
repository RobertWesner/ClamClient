package packets

import (
	"errors"
	"fmt"
)

type Packet103SetSlot struct {
	WindowId  uint8
	Slot      int16
	ItemId    int16
	ItemCount uint8
	ItemUses  uint8
}

func (p Packet103SetSlot) Id() uint8 {
	return 0x67
}

func (p Packet103SetSlot) Bytes() ([]byte, error) {
	return []byte{}, errors.New("103 server->client packets should never be sent")
}

func (p Packet103SetSlot) Read(reader PacketReader) error {
	var err error

	if p.WindowId, err = reader.Byte(); err != nil {
		return fmt.Errorf("103 read windowid: %w", err)
	}

	if p.Slot, err = reader.Int16(); err != nil {
		return fmt.Errorf("103 read slot: %w", err)
	}

	if p.ItemId, err = reader.Int16(); err != nil {
		return fmt.Errorf("103 read itemid: %w", err)
	}

	if p.ItemId > -1 {
		if p.ItemCount, err = reader.Byte(); err != nil {
			return fmt.Errorf("103 read itemid: %w", err)
		}

		if p.ItemUses, err = reader.Byte(); err != nil {
			return fmt.Errorf("103 read itemid: %w", err)
		}
	}

	return nil
}

package packets

import (
	"errors"
	"fmt"
)

type Packet104WindowItems struct {
	WindowId uint8
	Count    int16
	Payload  InventoryData
}

func (p Packet104WindowItems) Id() uint8 {
	return 0x68
}

func (p Packet104WindowItems) Bytes() ([]byte, error) {
	return []byte{}, errors.New("104 server->client packets should never be sent")
}

func (p Packet104WindowItems) Read(reader PacketReader) error {
	var err error

	if p.WindowId, err = reader.Byte(); err != nil {
		return fmt.Errorf("104 read windowid: %w", err)
	}

	if p.Count, err = reader.Int16(); err != nil {
		return fmt.Errorf("104 read count: %w", err)
	}

	if p.Payload, err = parseInventoryPayload(reader, p.Count); err != nil {
		return fmt.Errorf("104 read count: %w", err)
	}

	return nil
}

type ItemData struct {
	Id    int16
	Count uint8
	Uses  int16
}

type InventoryData []ItemData

func parseInventoryPayload(reader PacketReader, count int16) (InventoryData, error) {
	var err error

	data := InventoryData{}

	for range count {
		item := ItemData{}
		item.Id, err = reader.Int16()
		if err != nil {
			return InventoryData{}, fmt.Errorf("window items id: %w", err)
		}

		if item.Id > -1 {
			item.Count, err = reader.Byte()
			if err != nil {
				return InventoryData{}, fmt.Errorf("window items count: %w", err)
			}

			item.Uses, err = reader.Int16()
			if err != nil {
				return InventoryData{}, fmt.Errorf("window items uses: %w", err)
			}
		}

		data = append(data, item)
	}

	return data, nil
}

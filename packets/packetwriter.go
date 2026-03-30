package packets

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type PacketWriter struct {
	buffer *bytes.Buffer
}

func NewWriter() *PacketWriter {
	return &PacketWriter{
		buffer: &bytes.Buffer{},
	}
}

func (pw *PacketWriter) Write(value any) error {
	switch x := value.(type) {
	case byte:
		return pw.buffer.WriteByte(x)
	case int8:
		return pw.buffer.WriteByte(byte(x))
	case int16:
		var buf [2]byte
		binary.BigEndian.PutUint16(buf[:], uint16(x))
		_, err := pw.buffer.Write(buf[:])
		return err
	case int32:
		var buf [4]byte
		binary.BigEndian.PutUint32(buf[:], uint32(x))
		_, err := pw.buffer.Write(buf[:])
		return err
	case int64:
		var buf [8]byte
		binary.BigEndian.PutUint64(buf[:], uint64(x))
		_, err := pw.buffer.Write(buf[:])
		return err
	case string:
		return errors.New("use WriteString16() for strings")
	case []byte:
		_, err := pw.buffer.Write(x)
		return err
	default:
		return fmt.Errorf("unsupported type %T", value)
	}
}

func (pw *PacketWriter) WriteString16(s string) error {
	runes := []rune(s)

	// length = number of UCS-2 characters
	if len(runes) > 32767 {
		return fmt.Errorf("string16 too long")
	}

	// write length (int16)
	if err := pw.Write(int16(len(runes))); err != nil {
		return err
	}

	for _, r := range runes {
		if r > 0xFFFF {
			return fmt.Errorf("string16: rune out of UCS-2 range: %U", r)
		}

		if err := pw.Write(int16(r)); err != nil {
			return err
		}
	}

	return nil
}

func (pw *PacketWriter) Bytes() []byte {
	return pw.buffer.Bytes()
}

func (pw *PacketWriter) Reset() {
	pw.buffer.Reset()
}

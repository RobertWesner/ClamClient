package packets

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type PacketReader struct {
	r io.Reader
}

func NewReader(r io.Reader) *PacketReader {
	return &PacketReader{r: r}
}

func NewReaderFromBytes(b []byte) *PacketReader {
	return &PacketReader{r: bytes.NewReader(b)}
}

func (pr *PacketReader) Skip(n int) error {
	if n < 0 {
		return fmt.Errorf("skip: negative count %d", n)
	}

	var buf [512]byte
	for n > 0 {
		chunk := n
		if chunk > len(buf) {
			chunk = len(buf)
		}
		if _, err := io.ReadFull(pr.r, buf[:chunk]); err != nil {
			return err
		}
		n -= chunk
	}
	return nil
}

func (pr *PacketReader) Bool() (bool, error) {
	v, err := pr.Byte()
	if err != nil {
		return false, err
	}

	return v != 0, nil
}

func (pr *PacketReader) Int16() (int16, error) {
	var buf [2]byte
	if _, err := io.ReadFull(pr.r, buf[:]); err != nil {
		return 0, err
	}

	return int16(binary.BigEndian.Uint16(buf[:])), nil
}

func (pr *PacketReader) Int32() (int32, error) {
	var buf [4]byte
	if _, err := io.ReadFull(pr.r, buf[:]); err != nil {
		return 0, err
	}

	return int32(binary.BigEndian.Uint32(buf[:])), nil
}

func (pr *PacketReader) Int64() (int64, error) {
	var buf [8]byte
	if _, err := io.ReadFull(pr.r, buf[:]); err != nil {
		return 0, err
	}

	return int64(binary.BigEndian.Uint64(buf[:])), nil
}

func (pr *PacketReader) Float32() (float32, error) {
	bits, err := pr.Int32()
	if err != nil {
		return 0, err
	}

	return math.Float32frombits(uint32(bits)), nil
}

func (pr *PacketReader) Float64() (float64, error) {
	bits, err := pr.Int64()
	if err != nil {
		return 0, err
	}

	return math.Float64frombits(uint64(bits)), nil
}

func (pr *PacketReader) Byte() (byte, error) {
	var buf [1]byte
	if _, err := io.ReadFull(pr.r, buf[:]); err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (pr *PacketReader) Bytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, fmt.Errorf("bytes: negative length %d", n)
	}

	buf := make([]byte, n)
	if _, err := io.ReadFull(pr.r, buf); err != nil {
		return nil, err
	}

	return buf, nil
}

func (pr *PacketReader) String16() (string, error) {
	length, err := pr.Int16()
	if err != nil {
		return "", err
	}

	if length < 0 {
		return "", fmt.Errorf("string16: negative length %d", length)
	}

	buf := make([]byte, int(length)*2)

	if _, err := io.ReadFull(pr.r, buf); err != nil {
		return "", err
	}

	runes := make([]rune, length)

	for i := 0; i < int(length); i++ {
		code := binary.BigEndian.Uint16(buf[i*2:])

		// Reject surrogate range (not valid UCS-2)
		if code >= 0xD800 && code <= 0xDFFF {
			return "", fmt.Errorf("string16: surrogate detected (invalid UCS-2)")
		}

		runes[i] = rune(code)
	}

	return string(runes), nil
}

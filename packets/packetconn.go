package packets

import (
	"fmt"
	"io"
	"net"
)

type PacketConn struct {
	conn   net.Conn
	reader PacketReader
}

func NewPacketConn(conn net.Conn) PacketConn {
	return PacketConn{
		conn:   conn,
		reader: PacketReader{conn},
	}
}

func (pc PacketConn) Write(packet Packet) error {
	bytes, err := packet.Bytes()
	if err != nil {
		return fmt.Errorf("get packet bytes: %w", err)
	}

	_, err = pc.conn.Write(
		append(
			[]byte{packet.Id()},
			bytes...,
		),
	)
	if err != nil {
		return fmt.Errorf("PacketConn write: %w", err)
	}

	return nil
}

func (pc PacketConn) Read() (Packet, error) {
	var buf [1]byte
	if _, err := io.ReadFull(pc.conn, buf[:]); err != nil {
		return nil, fmt.Errorf("PacketConn read id: %w", err)
	}

	packet := func() Packet {
		switch buf[0] {
		case 0x00:
			return Packet0KeepAlive{}
		case 0x01:
			return Packet1Login{}
		case 0x02:
			return Packet2Handshake{}
		case 0x03:
			return Packet3ChatMessage{}
		case 0x04:
			return Packet4TimeUpdate{}
		case 0x05:
			// TODO
		case 0x06:
			return Packet6SpawnPosition{}
		case 0x07:
			// TODO
		case 0x08:
			// TODO
		case 0x09:
			// TODO
		case 0x0A:
			// TODO
		case 0x0B:
			// TODO
		case 0x0C:
			// TODO
		case 0x0D:
			// TODO
		case 0x0E:
			// TODO
		case 0x0F:
			// TODO
		case 0x10:
			// TODO
		case 0x11:
			// TODO
		case 0x12:
			// TODO
		case 0x13:
			// TODO
		case 0x14:
			// TODO
		case 0x15:
			// TODO
		case 0x16:
			// TODO
		case 0x17:
			// TODO
		case 0x18:
			// TODO
		case 0x19:
			// TODO
		case 0x1B:
			// TODO
		case 0x1C:
			// TODO
		case 0x1D:
			// TODO
		case 0x1E:
			// TODO
		case 0x1F:
			return Packet31EntityRelativeMove{}
		case 0x20:
			return Packet32EntityLook{}
		case 0x21:
			return Packet33LookAndRelativeMove{}
		case 0x22:
			// TODO
		case 0x26:
			// TODO
		case 0x27:
			// TODO
		case 0x28:
			// TODO
		case 0x32:
			// TODO
		case 0x33:
			// TODO
		case 0x34:
			// TODO
		case 0x35:
			// TODO
		case 0x36:
			// TODO
		case 0x3C:
			// TODO
		case 0x3D:
			// TODO
		case 0x46:
			// TODO
		case 0x47:
			// TODO
		case 0x64:
			// TODO
		case 0x65:
			// TODO
		case 0x66:
			// TODO
		case 0x67:
			// TODO
		case 0x68:
			// TODO
		case 0x69:
			// TODO
		case 0x6A:
			// TODO
		case 0x82:
			// TODO
		case 0x83:
			// TODO
		case 0xC8:
			// TODO
		case 0xFF:
			// TODO
		}

		return nil
	}()

	if packet == nil {
		return nil, fmt.Errorf("unknown packet 0x%X", buf[0])
	}

	err := packet.Read(pc.reader)
	if err != nil {
		return nil, fmt.Errorf("PacketConn read packet: %w", err)
	}

	return packet, nil
}

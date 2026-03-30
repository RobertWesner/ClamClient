package client

import (
	"clamclient/game"
	"clamclient/packets"
	"fmt"
)

func (c *Client) handlePackets() {
	for {
		packet, err := c.packetConn.Read()
		if err != nil {
			c.fail(err)

			return
		}

		switch p := packet.(type) {
		case *packets.Packet0KeepAlive:
			break
		case *packets.Packet1Login:
			fmt.Println(p) // TODO: remove

			err = c.packetConn.Write(packets.NewPacket2Handshake(c.username))
			if err != nil {
				c.fail(err)

				return
			}
		case packets.Packet2Handshake:
			fmt.Println(p) // TODO: remove

			c.state.ConnectionHash = p.UsernameOrConnectionHash
		case packets.Packet3ChatMessage:
			c.events.on.chat <- p.Message
		case packets.Packet4TimeUpdate:
			c.state.Time = p.Time
		case packets.Packet6SpawnPosition:
			fmt.Println(p) // TODO: remove

			c.state.SpawnPosition = game.NewVec3(float64(p.X), float64(p.Y), float64(p.Z))
			close(c.ready)
		case packets.Packet31EntityRelativeMove:
			select {
			case c.events.on.entityMove <- EntityMoveEvent{p.EntityID, game.NewVec3(float64(p.DX), float64(p.DY), float64(p.DZ))}:
			default:
			}
		case packets.Packet32EntityLook:
			select {
			case c.events.on.entityLook <- EntityLookEvent{p.EntityID, game.NewAngle(float64(p.Pitch.To360()), float64(p.Yaw.To360()))}:
			default:
			}
		case packets.Packet33LookAndRelativeMove:
			select {
			case c.events.on.entityMove <- EntityMoveEvent{p.EntityID, game.NewVec3(float64(p.DX), float64(p.DY), float64(p.DZ))}:
			default:
			}
			select {
			case c.events.on.entityLook <- EntityLookEvent{p.EntityID, game.NewAngle(float64(p.Pitch.To360()), float64(p.Yaw.To360()))}:
			default:
			}
		}
	}
}

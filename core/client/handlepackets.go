package client

import (
	"fmt"

	"github.com/RobertWesner/ClamClient/core/game"
	"github.com/RobertWesner/ClamClient/core/material"
	"github.com/RobertWesner/ClamClient/core/packets"
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
			c.events.on.entityMove <- EntityMoveEvent{int(p.EntityID), game.NewVec3(float64(p.DX), float64(p.DY), float64(p.DZ))}
		case packets.Packet32EntityLook:
			c.events.on.entityLook <- EntityLookEvent{int(p.EntityID), game.NewAngle(float64(p.Pitch.To360()), float64(p.Yaw.To360()))}
		case packets.Packet33LookAndRelativeMove:
			c.events.on.entityMove <- EntityMoveEvent{int(p.EntityID), game.NewVec3(float64(p.DX), float64(p.DY), float64(p.DZ))}
			c.events.on.entityLook <- EntityLookEvent{int(p.EntityID), game.NewAngle(float64(p.Pitch.To360()), float64(p.Yaw.To360()))}
		case packets.Packet51MapChunk:
			// TODO TODO TODO !!!
			HandleChunkData(p)
		case packets.Packet52MultiBlockChange:
			for i := int16(0); i < p.ArraySize; i++ {
				coords := p.CoordinateArray[i]
				x := (coords >> 12) & 0x0F
				z := (coords >> 8) & 0x0F
				y := coords & 0xFF

				c.events.on.blockChange <- BlockChangeEvent{
					int(p.ChunkX)*16 + int(x), int(y), int(p.ChunkZ)*16 + int(z),
					material.FromID(int(p.TypeArray[i])),
					p.MetadataArray[i],
				}
			}
		case packets.Packet53BlockChange:
			c.events.on.blockChange <- BlockChangeEvent{int(p.X), int(p.Y), int(p.Z), material.FromID(int(p.BlockType)), p.BlockMetadata}
		case packets.Packet70NewOrInvalidState:
			switch p.Reason {
			case 0:
				// ignore
			case 1:
				c.events.on.rain <- true
			case 2:
				c.events.on.rain <- false
			default:
				// ignore
			}
		case packets.Packet103SetSlot:
			c.events.on.setSlot <- SetSlotEvent{int(p.WindowID), int(p.Slot), material.FromID(int(p.ItemID)), int(p.ItemCount), int(p.ItemUses)}
		case packets.Packet104WindowItems:
			c.events.on.windowItems <- WindowItemsEvent{int(p.WindowID), int(p.Count), p.Payload}
		case packets.Packet106Transaction:
			c.events.on.transaction <- TransactionEvent{int(p.WindowID), int(p.ActionNumber), p.Accepted}
		case packets.Packet255Disconnect:
			c.events.on.disconnect <- p.Reason
		}
	}
}

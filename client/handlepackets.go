package client

import (
	"ClamClient/game"
	"ClamClient/packets"
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

			c.state.SpawnPosition = game.Vec3{X: float64(p.X), Y: float64(p.Y), Z: float64(p.Z)}
			close(c.ready)
		}
	}
}

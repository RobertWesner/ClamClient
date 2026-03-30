package client

import (
	"log/slog"

	"github.com/RobertWesner/ClamClient/core/packets"
)

func (c *Client) loopControls() {
	for {
		select {
		case message := <-c.commands.chatChan():
			err := c.packetConn.Write(packets.NewPacket3ChatMessage(message))
			if err != nil {
				slog.Error("could not send chat", "err", err)
			}
		case <-c.commands.disconnectChan():
			err := c.packetConn.Write(packets.NewPacket255Disconnect())
			if err != nil {
				slog.Error("could not send disconnect", "err", err)
			}
			c.Close()
		}
	}
}

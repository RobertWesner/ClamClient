package client

func (c *Client) Chat(message string) {
	c.commands.doChat(message)
}

func (c *Client) Disconnect() {
	c.commands.doDisconnect()
}

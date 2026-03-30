package client

func (c *Client) Chat(message string) {
	c.commands.doChat(message)
}

package client

type Commands struct {
	chat       chan string
	disconnect chan struct{}
}

func NewCommands() *Commands {
	return &Commands{
		chat: make(chan string, 32),
	}
}

func (c *Commands) doChat(message string) {
	c.chat <- message
}

func (c *Commands) chatChan() <-chan string {
	return c.chat
}

func (c *Commands) doDisconnect() {
	close(c.disconnect)
}

func (c *Commands) disconnectChan() <-chan struct{} {
	return c.disconnect
}

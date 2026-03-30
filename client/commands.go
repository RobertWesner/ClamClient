package client

type Commands struct {
	chat chan string
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

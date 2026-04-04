package client

type Commands struct {
	chat        chan string
	disconnect  chan struct{}
	transaction chan TransactionCommand
}

func NewCommands() *Commands {
	return &Commands{
		chat:        make(chan string, 32),
		disconnect:  make(chan struct{}),
		transaction: make(chan TransactionCommand),
	}
}

type TransactionCommand struct {
	WindowID     int
	ActionNumber int
	Accepted     bool
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

func (c *Commands) doTransaction(
	windowID int,
	actionNumber int,
	accepted bool,
) {
	c.transaction <- TransactionCommand{
		WindowID:     windowID,
		ActionNumber: actionNumber,
		Accepted:     accepted,
	}
}

func (c *Commands) transactionChan() <-chan TransactionCommand {
	return c.transaction
}

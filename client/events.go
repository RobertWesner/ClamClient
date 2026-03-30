package client

type Events struct {
	chat chan string

	on eventsOn
}

type eventsOn struct {
	chat chan<- string
}

func NewEvents() *Events {
	chat := make(chan string, 32)

	return &Events{
		chat: chat,
		on: eventsOn{
			chat: chat,
		},
	}
}

func (e *Events) Chat() <-chan string {
	return e.chat
}

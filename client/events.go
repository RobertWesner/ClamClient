package client

import "clamclient/game"

type EntityMoveEvent struct {
	EntityId int32
	Vector   game.Vec3
}

type EntityLookEvent struct {
	EntityId int32
	Angle    game.Angle
}

type Events struct {
	chat       chan string
	entityMove chan EntityMoveEvent
	entityLook chan EntityLookEvent

	on eventsOn
}

type eventsOn struct {
	chat       chan<- string
	entityMove chan<- EntityMoveEvent
	entityLook chan<- EntityLookEvent
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

func (e *Events) EntityMove() <-chan EntityMoveEvent {
	return e.entityMove
}

func (e *Events) EntityLook() <-chan EntityLookEvent {
	return e.entityLook
}

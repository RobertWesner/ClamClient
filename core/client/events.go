package client

import (
	"github.com/RobertWesner/ClamClient/core/game"
	"github.com/RobertWesner/ClamClient/core/packets"
)

type PlayerMoveAndLookEvent struct {
	Vector   game.Vec3
	Angle    game.Angle
	Stance   float64
	OnGround bool
}

type EntityMoveEvent struct {
	EntityId int
	Vector   game.Vec3
}

type EntityLookEvent struct {
	EntityId int
	Angle    game.Angle
}

type MapChunkEvent struct {
	X int
	Y int
	Z int
	// TODO
}

type SetSlotEvent struct {
	WindowId  int
	Slot      int
	ItemId    int
	ItemCount int
	ItemUses  int
}

type WindowItemsEvent struct {
	WindowId int
	Count    int
	Payload  packets.InventoryData
}

type TransactionEvent struct {
	WindowId     int
	ActionNumber int
	Accepted     bool
}

type Events struct {
	chat              <-chan string
	playerMoveAndLook <-chan PlayerMoveAndLookEvent
	entityMove        <-chan EntityMoveEvent
	entityLook        <-chan EntityLookEvent
	rain              <-chan bool
	setSlot           <-chan SetSlotEvent
	windowItems       <-chan WindowItemsEvent
	transaction       <-chan TransactionEvent
	disconnect        <-chan string

	on eventsOn
}

type eventsOn struct {
	chat              chan<- string
	playerMoveAndLook chan<- PlayerMoveAndLookEvent
	entityMove        chan<- EntityMoveEvent
	entityLook        chan<- EntityLookEvent
	rain              chan<- bool
	setSlot           chan<- SetSlotEvent
	windowItems       chan<- WindowItemsEvent
	transaction       chan<- TransactionEvent
	disconnect        chan<- string
}

func NewEvents() *Events {
	chat := make(chan string, 32)
	playerMoveAndLook := make(chan PlayerMoveAndLookEvent)
	entityMove := make(chan EntityMoveEvent)
	entityLook := make(chan EntityLookEvent)
	rain := make(chan bool)
	setSlot := make(chan SetSlotEvent)
	windowItems := make(chan WindowItemsEvent)
	transaction := make(chan TransactionEvent)
	disconnect := make(chan string)

	return &Events{
		chat:              chat,
		playerMoveAndLook: playerMoveAndLook,
		entityMove:        entityMove,
		entityLook:        entityLook,
		rain:              rain,
		setSlot:           setSlot,
		windowItems:       windowItems,
		transaction:       transaction,
		disconnect:        disconnect,
		on: eventsOn{
			chat:              chat,
			playerMoveAndLook: playerMoveAndLook,
			entityMove:        entityMove,
			entityLook:        entityLook,
			setSlot:           setSlot,
			rain:              rain,
			windowItems:       windowItems,
			transaction:       transaction,
			disconnect:        disconnect,
		},
	}
}

func (e *Events) Chat() <-chan string {
	return e.chat
}

func (e *Events) PlayerLookAndMove() <-chan PlayerMoveAndLookEvent {
	return e.playerMoveAndLook
}

func (e *Events) EntityMove() <-chan EntityMoveEvent {
	return e.entityMove
}

func (e *Events) EntityLook() <-chan EntityLookEvent {
	return e.entityLook
}

func (e *Events) Rain() <-chan bool {
	return e.rain
}

func (e *Events) SetSlot() <-chan SetSlotEvent {
	return e.setSlot
}

func (e *Events) WindowItems() <-chan WindowItemsEvent {
	return e.windowItems
}

func (e *Events) Transaction() <-chan TransactionEvent {
	return e.transaction
}

func (e *Events) Disconnect() <-chan string {
	return e.disconnect
}

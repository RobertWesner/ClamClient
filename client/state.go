package client

import "clamclient/game"

type State struct {
	EntityId       int32
	ConnectionHash string
	SpawnPosition  game.Vec3
	Time           int64
}

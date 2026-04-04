package client

import "github.com/RobertWesner/ClamClient/core/game"

type State struct {
	EntityID       int32
	ConnectionHash string
	SpawnPosition  game.Vec3
	Time           int64
}

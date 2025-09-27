package main

type GameStatus int
const (
	StatusIntro GameStatus = iota
	StatusOverworld
)

type GameState struct {
	Status GameStatus
	OverPosX, OverPosY int
}

var State GameState
func InitGameState() {
	State.Status = StatusIntro
	State.OverPosX = 64
	State.OverPosY = 64
}

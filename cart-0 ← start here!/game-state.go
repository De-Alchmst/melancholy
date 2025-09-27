package main

type GameStatus int
const (
	StatusIntro GameStatus = iota
	StatusOverworld
)

type GameState struct {
	Status GameStatus
	CurrentRoom *Room
}

var State GameState
func InitGameState() {
	State.Status = StatusIntro
	State.CurrentRoom = GetRoomAtID(0)
}

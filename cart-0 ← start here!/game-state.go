package main

type GameStatus int
const (
	StatusMessage GameStatus = iota
	StatusOverworld
)

type GameState struct {
	Status GameStatus
	CurrentRoom *Room
	Events map[string]byte
	CurrentMessage Message
}

var State GameState
func InitGameState() {
	State.Status = StatusMessage
	State.CurrentRoom = GetRoomAtID(0)
	State.Events = make(map[string]byte)
	State.CurrentMessage = IntroMessage
}

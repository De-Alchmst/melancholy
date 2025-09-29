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
	// sometimes I need entities with and somwtimes without the player
	OverworldEntsWithPlayer OverworldEntityList
}

var State GameState
func InitGameState() {
	State.Status = StatusMessage
	State.CurrentRoom = GetRoomAtID(0)
	State.Events = make(map[string]byte)
	State.CurrentMessage = IntroMessage
	SwitchRoom(0)
}


func RegisterEvent(name string, value byte) {
	State.Events[name] = value
}

func EventRegistered(name string) bool {
	_, ok := State.Events[name]
	return ok
}

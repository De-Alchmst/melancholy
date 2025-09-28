package main

//go:export start
func start() { // gruvbox light
	SetPallete(PalleteGruvboxLight)
	InitGameState()
	UpdatePressed()

	State.Status = StatusOverworld
	State.CurrentRoom = GetRoomAtID(2)
}

//go:export update
func update() {
	switch State.Status {
	case StatusMessage:
		UpdateMessage()
	case StatusOverworld:
		UpdateOverworld()
	}

	UpdatePressed()
}

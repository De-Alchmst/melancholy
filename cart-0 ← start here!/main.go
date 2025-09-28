package main

//go:export start
func start() { // gruvbox light
	SetPallete(PalleteGruvboxLight)
	InitGameState()

	State.Status = StatusOverworld
	State.CurrentRoom = GetRoomAtID(2)
}

//go:export update
func update() {
	switch State.Status {
	case StatusIntro:
		UpdateIntro()
	case StatusOverworld:
		UpdateOverworld()
	}
}

package main

//go:export start
func start() { // gruvbox light
	SetPallete(PalleteGruvboxLight)
	InitGameState()
	InitOverworldData()
	State.Status = StatusOverworld
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

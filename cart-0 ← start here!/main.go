package main

//go:export start
func start() { // gruvbox light
	SetPallete(PalleteGruvboxLight)
	InitGameState()
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

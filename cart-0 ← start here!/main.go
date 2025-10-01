package main

//go:export start
func start() { // gruvbox light
	SetPallete(PalleteGruvboxLight)
	InitGameState()
	UpdatePressed()

	State.Status = StatusOverworld
	SwitchRoom(6)
	// Player.Hitbox.X = tileToPos(8) + 3
	// Player.Hitbox.Y = tileToPos(4) + 3
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

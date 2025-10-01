package main

//go:export start
func start() { // gruvbox light
	SetPallete(PalleteGruvboxLight)
	InitGameState()
	UpdatePressed()

	State.Status = StatusOverworld
	SwitchRoom(5)
	Player.Hitbox.X = tileToPos(1) + 3
	Player.Hitbox.Y = tileToPos(2) + 3
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

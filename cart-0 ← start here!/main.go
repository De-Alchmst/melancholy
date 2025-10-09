package main

//go:export start
func start() { // gruvbox light
	InitGameState()
	UpdatePressed()

	 // here is where I would put my 'load save function...'
	//  if I had any!     

	State.Status = StatusOverworld
	State.Status = StatusBoss
	SwitchRoom(16)
	SetPallete(PalleteRustGold)
	RegisterEvent("got_second_key", 1)
	Player.Hitbox.X = tileToPos(8) + 3
	Player.Hitbox.Y = tileToPos(4) + 3
}

//go:export update
func update() {
	switch State.Status {
	case StatusMessage:
		UpdateMessage()
	case StatusOverworld:
		UpdateOverworld()
	case StatusBoss:
		UpdateBoss()
	}

	UpdatePressed()
}

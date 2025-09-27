package main

const (
	playerSpeed = 1.7
	playerAnimationDelay = 10
)

func (e *OverworldEntity) NextFrame() {
	e.AnimationIndex = (e.AnimationIndex + 1) % len(e.AnimationFrames)
}


func playerCollides() bool {
	box := Player.Hitbox
	px  := int(Player.X) + box.OffsetX
	py  := int(Player.Y) + box.OffsetY

	xStart := max( px                  / TileSize, 0)
	xEnd   := min((px + box.Width)    / TileSize, 9)
	yStart := max( py                / TileSize, 0)
	yEnd   := min((py + box.Height) / TileSize, 9)

	// This does not make me happy, but since player hitbox size is constant
	// it should be O(1)?
	for x := xStart; x <= xEnd; x++ {
		Y: // might look better one line above, but it still breaks the overall flow
		for y := yStart; y <= yEnd; y++ {
			for _, passthrough := range PassthroughTiles {
				if State.CurrentRoom.Tiles[y][x] == passthrough {
					continue Y
				}
			}
			return true
		}
	}
	return false
}


func movePlayer() {
	if pressed(KeyUp) {
		Player.Y -= playerSpeed
		if playerCollides() {
			Player.Y += playerSpeed
		}
		Player.Direction = DirUp
	}
	if pressed(KeyDown) {
		Player.Y += playerSpeed
		if playerCollides() {
			Player.Y -= playerSpeed
		}
		Player.Direction = DirDown
	}
	if pressed(KeyLeft) {
		Player.X -= playerSpeed
		if playerCollides() {
			Player.X += playerSpeed
		}
		Player.Direction = DirLeft
	}
	if pressed(KeyRight) {
		Player.X += playerSpeed
		if playerCollides() {
			Player.X -= playerSpeed
		}
		Player.Direction = DirRight
	}

	if pressed(KeyMovement) {
		Player.AnimationCountdown -= playerSpeed
		if Player.AnimationCountdown <= 0 {
			Player.AnimationCountdown = playerAnimationDelay
			Player.NextFrame()
		}

	} else {
		Player.AnimationIndex = 0
		Player.AnimationCountdown = 0
	}
}


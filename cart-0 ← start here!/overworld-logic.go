package main

const (
	playerSpeed = 1.7
	playerAnimationDelay = 10
)

func (e *OverworldEntity) NextFrame() {
	e.AnimationIndex = (e.AnimationIndex + 1) % len(e.AnimationFrames)
}


func SwitchRoom(newRoomID RoomID) {
	State.CurrentRoom = GetRoomAtID(newRoomID)
}


func playerCollides() bool {
	box := Player.Hitbox
	xStart := int(max( box.X                  / TileSize, 0))
	xEnd   := int(min((box.X + box.Width)    / TileSize, 9))
	yStart := int(max( box.Y                / TileSize, 0))
	yEnd   := int(min((box.Y + box.Height) / TileSize, 9))

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


func movePlayerDirections() {
	if pressed(KeyUp) {
		Player.Hitbox.Y -= playerSpeed
		if playerCollides() {
			Player.Hitbox.Y += playerSpeed
		}
		Player.Direction = DirUp
	}
	if pressed(KeyDown) {
		Player.Hitbox.Y += playerSpeed
		if playerCollides() {
			Player.Hitbox.Y -= playerSpeed
		}
		Player.Direction = DirDown
	}
	if pressed(KeyLeft) {
		Player.Hitbox.X -= playerSpeed
		if playerCollides() {
			Player.Hitbox.X += playerSpeed
		}
		Player.Direction = DirLeft
	}
	if pressed(KeyRight) {
		Player.Hitbox.X += playerSpeed
		if playerCollides() {
			Player.Hitbox.X -= playerSpeed
		}
		Player.Direction = DirRight
	}
}


func movePlayerAnimation() {
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


func movePlayerRooms() {
	box := &Player.Hitbox
	if box.X < 0 {
		SwitchRoom(State.CurrentRoom.Left)
		box.X = 10*TileSize - box.Width
	} else if box.X + box.Width > 10*TileSize {
		SwitchRoom(State.CurrentRoom.Right)
		box.X = 0
	} else if box.Y < 0 {
		SwitchRoom(State.CurrentRoom.Up)
		box.Y = 10*TileSize - box.Height
	} else if box.Y + box.Height > 10*TileSize {
		SwitchRoom(State.CurrentRoom.Down)
		box.Y = 0
	}
}


func MovePlayer() {
	movePlayerDirections()
	movePlayerAnimation()
	movePlayerRooms()
}


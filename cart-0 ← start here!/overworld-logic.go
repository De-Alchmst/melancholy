package main

const (
	playerSpeed = 1.7
	playerAnimationDelay = 10
)

func (e *OverworldEntity) NextFrame() {
	e.AnimationIndex = (e.AnimationIndex + 1) % len(e.AnimationFrames)
}


func movePlayer() {
	if pressed(KeyUp) {
		Player.Y -= playerSpeed
		Player.Direction = DirUp
	}
	if pressed(KeyDown) {
		Player.Y += playerSpeed
		Player.Direction = DirDown
	}
	if pressed(KeyLeft) {
		Player.X -= playerSpeed
		Player.Direction = DirLeft
	}
	if pressed(KeyRight) {
		Player.X += playerSpeed
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


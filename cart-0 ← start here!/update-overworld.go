package main

import "cart/w4"

const (
	playerSpeed = 2
)

func UpdateOverworld() {
	*w4.DRAW_COLORS = 0x22
	updatePlayer()
}


func updatePlayer() {
	movePlayer()
	drawPlayer()
}


func movePlayer() {
	if pressed(KeyUp) {
		State.OverPosY -= playerSpeed
	}
	if pressed(KeyDown) {
		State.OverPosY += playerSpeed
	}
	if pressed(KeyLeft) {
		State.OverPosX -= playerSpeed
	}
	if pressed(KeyRight) {
		State.OverPosX += playerSpeed
	}
}


func drawPlayer() {
	w4.Rect(State.OverPosX, State.OverPosY, 8, 8)
}

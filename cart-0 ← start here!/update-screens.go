package main

import "cart/w4"

func UpdateIntro() {
	*w4.DRAW_COLORS = 2
	w4.Text("And so it begins", 5, 10)
	w4.Text("Like many times\n         before", 20, 35)
	w4.Text("Like many times\n          after", 20, 70)

	if pressed(KeyAction) {
		State.Status = StatusOverworld
	}
}

package main

import "w4"

// what did you expected to find here

draw_game_over :: proc "c" () {
	w4.DRAW_COLORS^ = 0x41
	w4.text("Or was thre", 20, 42)
	w4.text("another way?", 30, 50)
}

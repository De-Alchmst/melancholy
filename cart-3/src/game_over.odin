package main

import "w4"

// what did you expected to find here

draw_game_over :: proc "c" () {
	w4.DRAW_COLORS^ = 0x41
	w4.text("Or was there", 20, 42)
	w4.text("another way?", 30, 50)
}

draw_game_won :: proc "c" () {
	w4.DRAW_COLORS^ = 0x14
	w4.text("You won", 20, 40)
	w4.text("The End.", 30, 50)
}

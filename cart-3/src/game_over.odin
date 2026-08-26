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
	w4.text("Not today,", 5, 40)
	w4.text("Yaldabaoth", 50, 50)
	w4.text("Not today", 70, 90)

	w4.blit(&YALDABAOTH_SPRITE[0], 10, 60, YALDABAOTH_WIDTH, YALDABAOTH_HEIGHT, YALDABAOTH_FLAGS)
}

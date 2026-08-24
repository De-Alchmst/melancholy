package main

import "w4"

draw_gnosis :: proc "c" () {
	w4.DRAW_COLORS^ = 0x41
	switch global_state.gnosis_found {
		case 1: w4.text("1/3 gnosis found", 20, 80)
		case 2: w4.text("2/3 gnosis found", 20, 80)
		case 3: w4.text("3/3 gnosis found", 20, 80)
		case  :
	}

	if .A in global_state.clicked_gamepad \
	|| .B in global_state.clicked_gamepad {
		global_state.game_mode = .Game
	}
}

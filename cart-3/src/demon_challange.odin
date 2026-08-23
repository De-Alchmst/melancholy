package main

import "w4"

// let's do random numbers the good old way
random_starts : []i32 = {
	5, 17, 0, 23, 16, 30, 24, 7, 12, 26, 12, 15, 0-9-0, FIREBALL_MOVE_DELAY
}

draw_demon_challange :: proc "c" () {
	w4.DRAW_COLORS^= 0x32
	w4.oval(80 - 70, 80 - 25, 140, 50)
	w4.DRAW_COLORS^= 0x31
	w4.oval(80 - 15, 80 - 25, 30, 50)
	w4.DRAW_COLORS^= 0x33
	w4.oval(80 - 10, 80 - 10, 20, 20)

	global_state.challange_pos += 3
	if global_state.challange_pos > 160 {
		play_death()
		switch_mode(.Game)
		return
	}

	w4.DRAW_COLORS^= 0x4
	w4.rect(global_state.challange_pos, 80 - 30, 5, 60)

	if .A in global_state.clicked_gamepad \
  || .B in global_state.clicked_gamepad {
		// hit
		if global_state.challange_pos > 80-20 && global_state.challange_pos < 80+15 {
			play_challange_success()
			global_state.game_mode = .Game
		// miss
		} else {
			play_death()
			switch_mode(.Game)
		}
	}
}

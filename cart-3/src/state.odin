package main

import "w4"

Game_Mode :: enum {
		Map,
		Text, // <-- you are here
		Game,
}

State :: struct {
	game_mode: Game_Mode,
	map_ticks: i32,
	map_index: i32,
	text_index: i32,
	held_gamepad, clicked_gamepad: w4.Buttons
}

global_state := State {
	game_mode       = .Text,
	map_ticks       = 0    ,
	map_index       = 0    ,
	held_gamepad    = {}   ,
	clicked_gamepad = {}   ,
}

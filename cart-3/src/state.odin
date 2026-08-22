package main

import "w4"

Game_Mode :: enum {
		Map,
		Text, // <-- you are here
		Game,
}

State :: struct {
	game_mode  : Game_Mode,
	map_ticks  : i32,
	map_index  : u8,
	text_index : u8,
	level_index: u8,

	hero_pos: Point,
	  // I have considered putting this on the heap, but adding all the code
	 // needed for allocation (`core:runtime`) took way more space than this did
	// so stack-only it is
	game_entities: [9]Game_Entity,

	held_gamepad, clicked_gamepad: w4.Buttons,
}

global_state := State {
	game_mode       = .Text,
	map_ticks       = 0    ,
	map_index       = 0    ,
	level_index     = 0    ,

	held_gamepad    = { /* THERE'S NOBODY HERE */ },
	clicked_gamepad = { /* there's nobody here */ },

	hero_pos = { -1, -1 },
}

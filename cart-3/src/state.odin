package main

import "w4"

Game_Mode :: enum {
		GG,
		Map,
		Text, // <-- you are here
		Game,
		Demon_Challange,
}

State :: struct {
	game_mode             : Game_Mode,
	map_ticks             : i32      ,
	game_ticks            : i32      ,
	map_index             : u8       ,
	text_index            : u8      ,
	level_index           : u8     ,
	challange_pos         : i32   ,
	challange_start_index : u8   ,
	level_completed       : bool,

	hero_state: Hero_State,
	  // I have considered putting this on the heap, but adding all the code
	 // needed for allocation (`core:runtime`) took way more space than this did
	// so stack-only it is
	game_entities: [15]Game_Entity,

	held_gamepad, clicked_gamepad: w4.Buttons,
}

global_state := State {
	game_mode             = .Text,
	map_ticks             = 0    ,
	game_ticks            = 0    ,
	map_index             = 0    ,
	level_index           = 0    ,
	challange_start_index = 0    ,
	challange_pos         = 0    ,
	level_completed       = false,

	held_gamepad    = { /* THERE'S NOBODY HERE */ },
	clicked_gamepad = { /* there's nobody here */ },
}

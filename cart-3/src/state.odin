package main

import "w4"

Game_Mode :: enum {
		GG,
		Map,
		Text, // <-- you are here
		Game,
		Gnosis,
		Secret_Maze,
		Demon_Challange,
}

State :: struct {
	game_mode             : Game_Mode,
	secret_pos            : Point    ,
	game_ticks            : i32      ,
	map_index             : u8       ,
	text_index            : u8      ,
	level_index           : u8     ,
	level_completed       : bool  ,
	challange_pos         : i32  ,
	challange_start_index : u8  ,
	map_ticks             : i32,
	gnosis_found          : u8,


	hero_state: Hero_State,
	     // I have considered putting this on the heap, but adding all the code
	    // needed for allocation (`base:runtime`) took way more space than this
     // did so stack-only it is
 	  // well, maybe not a 50, but it acts a bit weird and I still don't even know
   // if the 64k of RAM is in addition to the cartrige, it is all added together
	// ... well, it's on the stack now, so deal with it
	game_entities: [50]Game_Entity,

	held_gamepad, clicked_gamepad: w4.Buttons,
}

global_state := State {
	game_mode             = .Text,
	map_ticks             = 0    ,
	game_ticks            = 0    ,
	map_index             = 0    ,
	text_index            = 0    ,
	level_index           = 0    ,
	challange_start_index = 0    ,
	challange_pos         = 0    ,
	gnosis_found          = 0    ,
	level_completed       = false,

	held_gamepad    = { /* THERE'S NOBODY HERE */ },
	clicked_gamepad = { /* there's nobody here */ },
}

package main

State :: struct {
	global_mode: enum {
		Map
	},
	map_ticks: i32,
	map_stage: i32,
}

global_state := State {
	global_mode = .Map,
	map_ticks   = 0   ,
	map_stage   = 0   ,
}

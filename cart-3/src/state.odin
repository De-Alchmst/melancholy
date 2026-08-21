package main

State :: struct {
	global_mode: enum {
		Map
	}
}

global_state := State {
	global_mode = .Map,
}

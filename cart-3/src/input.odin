package main

import "w4"

read_input :: proc "c" () {
	// code promptly stolen from wasm4.org
	global_state.clicked_gamepad =  w4.GAMEPAD1^ \
															 & (w4.GAMEPAD1^ ~ global_state.held_gamepad)
	global_state.held_gamepad    = w4.GAMEPAD1^
}

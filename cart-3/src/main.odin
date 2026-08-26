package main

@export
start :: proc "c" () {
	switch_mode(.Main_Menu)
}

@export
update :: proc "c" () {
	read_input()

	switch global_state.game_mode {
		case .Map:
			draw_map()
		case .Text:
			draw_text()
		case .Game:
			update_game()
			draw_game()
		case .GG:
			draw_game_over()
		case .End:
			draw_game_won()
		case .Demon_Challange:
			draw_demon_challange()
		case .Secret_Maze:
			draw_secret_maze()
		case .Gnosis:
			draw_gnosis()
		case .Main_Menu:
			draw_main_menu()
	}
}


import "w4"


draw_main_menu :: proc "c" () {
	// continue ...
	if .A in global_state.clicked_gamepad {
		switch_mode(.Text)
	}

	// colors are fun
	if   .B in global_state.held_gamepad do w4.PALETTE^ = HIS_PALETTE
	else                                 do w4.PALETTE^ = L3_PALETTE

	w4.DRAW_COLORS^ = 0x1324
	w4.blit(&MAIN_MENU_SPRITE[0], 0, 0, MAIN_MENU_WIDTH, MAIN_MENU_HEIGHT, MAIN_MENU_FLAGS)
}

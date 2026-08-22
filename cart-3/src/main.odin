package main

@export
start :: proc "c" () {
	switch_mode(.Game)
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
	}
}

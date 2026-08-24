package main

@export
start :: proc "c" () {
	switch_mode(.End)
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
	}
}

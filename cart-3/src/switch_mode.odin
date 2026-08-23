package main

import "w4"

// switches game mode and changes colors
switch_mode :: proc "c" (mode: Game_Mode) {
	global_state.game_mode  = mode
	switch mode {
	case .Map:  w4.PALETTE^ = MAP_PALETTE
	case .Text: w4.PALETTE^ = ALEXANDRIA[global_state.text_index].palette
	case .GG:   w4.PALETTE^ = MAP_PALETTE
	case .Game:
		w4.PALETTE^ = LEVELS[global_state.level_index].palette
		global_state.game_ticks = 0
		global_state.level_completed = false
		flush_entities()
		load_entities()
	}
}


flush_entities :: proc "c" () {
	for i in 0..<(len(global_state.game_entities)) {
		global_state.game_entities[i].type = .Nihil 
	}
}

load_entities :: proc "c" () {
	free_index := 0

	for y : i32 = 0; y < 20; y += 1 {
		for x : i32 = 0; x < 20; x += 1 {
			switch LEVELS[global_state.level_index].layout[y][x] {
				case TILE_PLAYER:
					global_state.hero_state = {
						pos = { x, y },
						direction = .Nihil,
					}
					// also skip, he's aside
					free_index -= 1

				case TILE_CULTIST:
					global_state.game_entities[free_index] = {
						pos = { x, y },
						type = .Cultist,
					}

				case TILE_DEAD:
					global_state.game_entities[free_index] = {
						pos = { x, y },
						type = .Dead,
					}

				case TILE_MAGA_UP:
					global_state.game_entities[free_index] = {
						pos = { x, y },
						type = .Maga_Up,
					}

				case TILE_MAGA_RIGHT:
					global_state.game_entities[free_index] = {
						pos = { x, y },
						type = .Maga_Right,
					}

				case TILE_MAGA_DOWN:
					global_state.game_entities[free_index] = {
						pos = { x, y },
						type = .Maga_Down,
					}

				case TILE_MAGA_LEFT:
					global_state.game_entities[free_index] = {
						pos = { x, y },
						type = .Maga_Left,
					}

				case:
					// skip
					free_index -= 1
			}

			free_index += 1
		}
	}
}

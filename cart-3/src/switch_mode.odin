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

				case:
					// skip
					free_index -= 1
			}

			free_index += 1
		}
	}
}


// https://colorhunt.co/palette/5e00069b0f06d53e0feed9b9
MAP_PALETTE :: w4.Palette {
	0x9B0F06,
	0xD53E0F,
	0x5E0006,
	0xEED9B9,
}

// https://colorhunt.co/palette/77bef0ffcb61ff894fea5b6f
L1_PALETTE :: w4.Palette {
	0xFF894F,
	0xFFCB61,
	0xEA5B6F,
	0x77BEF0,
}

// https://lospec.com/palette-list/wish-gb
L5_PALETTE :: w4.Palette {
	0x608fcf,
	0x7550e8,
	0x622e4c,
	0x8be5ff,
}

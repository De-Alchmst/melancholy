package main

import "w4"

draw_game :: proc "c" () {
	// skip if game already ended
	if global_state.game_mode != .Game do return

	// draw the map
	for y : i32 = 0; y < 20; y += 1 {
		for x : i32 = 0; x < 20; x += 1 {
			switch LEVELS[global_state.level_index].layout[y][x] {
				case TILE_SUSE:
					w4.DRAW_COLORS^ = 0x21
					w4.blit(&SUSE_SPRITE[0], x*8, y*8, 8, 8)

				case TILE_BOX:
					w4.DRAW_COLORS^ = 0x31
					w4.blit(&BOX_SPRITE[0], x*8, y*8, 8, 8)

				case TILE_ROCK:
					w4.DRAW_COLORS^ = 0x31
					w4.blit(&ROCK_SPRITE[0], x*8, y*8, 8, 8)

				case TILE_BRICK:
					w4.DRAW_COLORS^ = 0x23
					w4.blit(&BRICK_SPRITE[0], x*8, y*8, 8, 8)
				case:
			}
		}
	}

	// draw the living things

	for ent in global_state.game_entities {
		switch ent.type {
			case .Cultist:
				w4.DRAW_COLORS^ = 0x31
				w4.blit(&CULTIST_SPRITE[0], ent.pos.x*8, ent.pos.y*8, 8, 8)

			case .Dead:
				w4.DRAW_COLORS^ = 0x21
				w4.blit(&DEAD_SPRITE   [0], ent.pos.x*8, ent.pos.y*8, 8, 8)

			case .Maga_Up, .Maga_Right, .Maga_Down, .Maga_Left:
				w4.DRAW_COLORS^ = 0x21
				w4.blit(&MAGA_SPRITE[0], ent.pos.x*8, ent.pos.y*8, 8, 8)

			case .Fireball_Up, .Fireball_Right, .Fireball_Down, .Fireball_Left:
				w4.DRAW_COLORS^ = 0x21
				w4.blit(&FIREBALL_SPRITE[0], ent.pos.x*8, ent.pos.y*8, 8, 8)

			case .Nihil:
		}
	}

	w4.DRAW_COLORS^ = 0x41
	w4.blit(&HERO_GAME_SPRITE[0],
	      	global_state.hero_state.pos.x*8, global_state.hero_state.pos.y*8,
					8, 8)
}

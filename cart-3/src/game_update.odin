package main

import "w4"

Game_Entity :: struct {
	pos: Point,
	type: enum {
		Cultist,
		Nihil  ,
		Dead   ,
	}
}

Direction :: enum {
	Up, Right, Down, Left, Nihil
}

Hero_State :: struct {
	pos: Point,
	direction: Direction,
}

update_game :: proc "c" () {
	update_player()
	update_entities() // and end_the_gamemode()
}


update_player :: proc "c" () {
	// if not moving
	if global_state.hero_state.direction == .Nihil {
		
		  /**/ if .UP    in global_state.held_gamepad {
			global_state.hero_state.direction = .Up
		} else if .RIGHT in global_state.held_gamepad {
			global_state.hero_state.direction = .Right
		} else if .DOWN  in global_state.held_gamepad {
			global_state.hero_state.direction = .Down
		} else if .LEFT  in global_state.held_gamepad {
			global_state.hero_state.direction = .Left
		}

	// if in motion
	} else {
		global_state.game_ticks += 1

		if global_state.game_ticks == 2 {
			global_state.game_ticks = 0
			new_pos := global_state.hero_state.pos

			// try moving forward
			switch global_state.hero_state.direction {
				case .Up:    new_pos.y -= 1
				case .Right: new_pos.x += 1
				case .Down:  new_pos.y += 1
				case .Left:  new_pos.x -= 1
				case .Nihil: unreachable()
			}

			// if there's a wall in the way
			if solid_tile_pos_p(new_pos) {
				global_state.hero_state.direction = .Nihil

			// else
			} else {
				global_state.hero_state.pos = new_pos
				kill_entities_at(new_pos)
			}
		}
	}
}


kill_entities_at :: proc "c" (pos: Point) {
	for i in 0..<(len(global_state.game_entities)) {
		if (global_state.game_entities[i].pos == pos) {
			global_state.game_entities[i].type = .Dead 
		}
	}
}


update_entities :: proc "c" () {
	some_alive := false

	for i in 0..<(len(global_state.game_entities)) {
		ent := global_state.game_entities[i]
		// skip nonexistent and dead
		if ent.type == .Nihil || ent.type == .Dead do continue

		some_alive = true
	}

	// end game if all dead
	if !some_alive {
		global_state.level_index += 1
		switch_mode(.Map)
	}
}

package main

HERO_MOVE_DELAY      :: 2
FIREBALL_MOVE_DELAY  :: 5
FIREBALL_SPAWN_DELAY :: 51
GAME_TICKS_LOOP      :: 102

Game_Entity_Type :: enum {
	Cultist,
	Nihil   ,
	Dead     ,
	Maga_Right,
	Maga_Down ,
	Maga_Left  ,
	Maga_Up     ,
	Fireball_Up  ,
	Fireball_Right,
	Fireball_Down ,
	Fireball_Left,
}

Game_Entity :: struct {
	pos: Point,
	type: Game_Entity_Type,
}

Direction :: enum {
	Up, Right, Down, Left, Nihil
}

Hero_State :: struct {
	pos: Point,
	direction: Direction,
}

update_game :: proc "c" () {
	global_state.game_ticks = 1 + (global_state.game_ticks % GAME_TICKS_LOOP)

	update_player()
	update_entities() // and end_the_gamemode()

	// reset if needed
	if .B in global_state.clicked_gamepad {
		switch_mode(.Game)
	}

	// end after 1s/3 as to not seem too scuffed
	if global_state.level_completed && global_state.game_ticks == 20 {
		global_state.level_index += 1
		switch_mode(.Map)
	}
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

		if global_state.hero_state.direction != .Nihil {
			play_hero_launch()
		}

	// if in motion
	} else {
		if global_state.game_ticks % HERO_MOVE_DELAY == 0 {
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
		ent := global_state.game_entities[i]
		if (ent.pos == pos) {

			// if the killed entity is a fireball, kills the player instead
			if ent.type == .Fireball_Up   || ent.type == .Fireball_Right \
			|| ent.type == .Fireball_Down || ent.type == .Fireball_Left { \
				switch_mode(.Game)
			}
			else do global_state.game_entities[i].type = .Dead 

			// either way, something dies today
				play_death()
		}
	}
}


update_entities :: proc "c" () {
	some_alive := false

	for i in 0..<(len(global_state.game_entities)) {
		ent := &global_state.game_entities[i]
		// skip non-existent and dead
		if ent.type == .Nihil || ent.type == .Dead do continue

		// MAGÆ
		if global_state.game_ticks % FIREBALL_SPAWN_DELAY == 0 {
			play_fireball_spawn()
			#partial switch ent.type {
				case .Maga_Up    : spawn_entity(.Fireball_Up   , ent.pos.x  , ent.pos.y-1)
				case .Maga_Right : spawn_entity(.Fireball_Right, ent.pos.x+1, ent.pos.y  )
				case .Maga_Down  : spawn_entity(.Fireball_Down , ent.pos.x  , ent.pos.y+1)
				case .Maga_Left  : spawn_entity(.Fireball_Left , ent.pos.x-1, ent.pos.y  )
			}
		}

		// fire is stored in the BALLZ
		if ent.type == .Fireball_Up   || ent.type == .Fireball_Right \
		|| ent.type == .Fireball_Down || ent.type == .Fireball_Left { \
			if global_state.game_ticks % FIREBALL_MOVE_DELAY == 0 {
				#partial switch ent.type {
					case .Fireball_Up    : ent.pos.y -= 1
					case .Fireball_Right : ent.pos.x += 1
					case .Fireball_Down  : ent.pos.y += 1
					case .Fireball_Left  : ent.pos.x -= 1
				}

				if (solid_tile_pos_p(ent.pos))  do  ent.type = .Nihil
			}

		// rest is living
		} else {
			some_alive = true
		}
	}

	// end game if all dead
	// mark game as complete
	// and prepare timeout
	if !(some_alive || global_state.level_completed) {
		global_state.level_completed = true
		global_state.game_ticks = 0
	}
}


spawn_entity :: proc "c" (type: Game_Entity_Type, x, y: i32) {
	// replace first non-entity in our entity list
	for i in 0..<(len(global_state.game_entities)) {
		if global_state.game_entities[i].type == .Nihil {
			global_state.game_entities[i] = {
				type = type,
				pos = { x, y },
			}
			return
		}
	}

	// if no empty space, just ignore, lel
	return
}

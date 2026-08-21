package main

import "w4"
import "base:intrinsics"

The_Path :: [6]Point


draw_map :: proc "c" () {
	CENTER :: Point { 80, 80 }
	RADIUS :: 60

	// LinkError: WebAssembly.instantiate(): Import #3 "env" "sinf": function import requires a callable
	// so magic numbers it is!
	// STEP   :: 2 * math.π / 5
	head := Point {
		CENTER.x,
		CENTER.y - RADIUS,
	}
	right_arm := Point {
		CENTER.x - (i32)(intrinsics.constant_floor(RADIUS * -0.951056)),
		CENTER.y - (i32)(intrinsics.constant_floor(RADIUS * 0.309017)),
	}
	right_leg := Point {
		CENTER.x - (i32)(intrinsics.constant_floor(RADIUS * -0.587785)),
		CENTER.y - (i32)(intrinsics.constant_floor(RADIUS * -0.809017)),
	}
	left_leg := Point {
		CENTER.x - (i32)(intrinsics.constant_floor(RADIUS * 0.587785)),
		CENTER.y - (i32)(intrinsics.constant_floor(RADIUS * -0.809017)),
	}
	left_arm := Point {
		CENTER.x - (i32)(intrinsics.constant_floor(RADIUS * 0.951056)),
		CENTER.y - (i32)(intrinsics.constant_floor(RADIUS * 0.309017)),
	}

	the_path : The_Path = {
		head, right_leg, left_arm, right_arm, left_leg, CENTER
	}
	
	/// first draw the circle
	w4.DRAW_COLORS^= 0x31
	/// add a bit of radius, just in case
	w4.oval(CENTER.x - RADIUS, CENTER.y - RADIUS, RADIUS*2+1, RADIUS*2+1)

	/// then draw The Path
	w4.DRAW_COLORS^= 0x3
	for i := 0; i < 5; i += 1 {
		from := the_path[    i    ]
		to   := the_path[(i+1) % 5] // last line to head instead of center
		w4.line(from.x, from.y, to.x, to.y)
	}

	/// and the center...
	w4.DRAW_COLORS^= 0x32
	w4.oval(CENTER.x-9, CENTER.y-5, 20, 10)
	w4.DRAW_COLORS^= 0x31
	w4.oval(CENTER.x-4, CENTER.y-5, 10, 10)
	w4.DRAW_COLORS^= 0x33
	w4.oval(CENTER.x-1, CENTER.y-2, 4, 4)

	/// and finally, the hero

	// update the state
	global_state.map_ticks += 1
	hero_percent := global_state.map_ticks
	if hero_percent > 100 do hero_percent = 100

	// unsmooth hero progression
	hero_percent = (hero_percent / 3) * 3

	// get progress position
	total_Δx := the_path[global_state.map_stage+1].x - the_path[global_state.map_stage].x
	total_Δy := the_path[global_state.map_stage+1].y - the_path[global_state.map_stage].y

	Δx := i32((f16(total_Δx * hero_percent) / 100))
	Δy := i32((f16(total_Δy * hero_percent) / 100))

	// and apply it
	w4.DRAW_COLORS^= 0x41
	w4.blit(&hero_map_sprite[0],
		the_path[global_state.map_stage].x + Δx - 8,
		the_path[global_state.map_stage].y + Δy - 16,
		16, 16)

	// progress onto the next stage
	if global_state.map_ticks > 150 {
		global_state.map_ticks = 0
		global_state.map_stage = ((global_state.map_stage + 1) % 5)
	}
}

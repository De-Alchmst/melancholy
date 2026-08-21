package main

import "w4"
import "core:math"

draw_map :: proc() {
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
		CENTER.x - (i32)(math.floor_f32(RADIUS * -0.951056)),
		CENTER.y - (i32)(math.floor_f32(RADIUS * 0.309017)),
	}
	right_leg := Point {
		CENTER.x - (i32)(math.floor_f32(RADIUS * -0.587785)),
		CENTER.y - (i32)(math.floor_f32(RADIUS * -0.809017)),
	}
	left_leg := Point {
		CENTER.x - (i32)(math.floor_f32(RADIUS * 0.587785)),
		CENTER.y - (i32)(math.floor_f32(RADIUS * -0.809017)),
	}
	left_arm := Point {
		CENTER.x - (i32)(math.floor_f32(RADIUS * 0.951056)),
		CENTER.y - (i32)(math.floor_f32(RADIUS * 0.309017)),
	}

	the_path : [6]Point = {
		head, right_leg, left_arm, right_arm, left_leg, CENTER
	}
	
	// draw the circle
	w4.DRAW_COLORS^= 0x31
	// add a bit of radius, just in case
	w4.oval(CENTER.x - RADIUS, CENTER.y - RADIUS, RADIUS*2+1, RADIUS*2+1)

	// draw The Path
	w4.DRAW_COLORS^= 0x3
	for i := 0; i < 5; i += 1 {
		from := the_path[    i    ]
		to   := the_path[(i+1) % 5] // last line to head instead of center
		w4.line(from.x, from.y, to.x, to.y)
	}

	// and the center
	w4.DRAW_COLORS^= 0x32
	w4.oval(CENTER.x-9, CENTER.y-5, 20, 10)
	w4.DRAW_COLORS^= 0x31
	w4.oval(CENTER.x-4, CENTER.y-5, 10, 10)
	w4.DRAW_COLORS^= 0x33
	w4.oval(CENTER.x-1, CENTER.y-2, 4, 4)
}

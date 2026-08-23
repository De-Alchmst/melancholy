package main

import "w4"

play_footstep :: proc "c" () {
	w4.tone(210 | (190 << 16), 7, 60, .Noise)
}

play_fireball_spawn :: proc "c" () {
	w4.tone(300 | (560 << 16), 10 | (25 << 8), 21, .Noise)
}

play_hero_launch :: proc "c" () {
	w4.tone(800 | (500 << 16), 15 | (5 << 8), 24, .Pulse1)
}

play_death :: proc "c" () {
	w4.tone(290 | (210 << 16), (5 << 8) | (25 << 16), 7, .Noise)
}

play_challange_success :: proc "c" () {
	w4.tone(750, 3 | (2 << 8) | (15 << 16), 22, .Pulse1)
}

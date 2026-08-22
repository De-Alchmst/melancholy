package main

import "w4"

Text_Line :: struct {
	str: cstring,
	pos: Point,
}

Text_Data :: struct {
	lines: []Text_Line,
	next_game_mode: Game_Mode,
	palette: w4.Palette,
}

draw_text :: proc "c" () {
	text := &ALEXANDRIA[global_state.text_index]

	w4.DRAW_COLORS^= 0x14
	for line in text.lines {
		w4.text(line.str, line.pos.x, line.pos.y)
	}

	// continue
	if .A in global_state.clicked_gamepad \
	|| .B in global_state.clicked_gamepad {
		global_state.text_index += 1
		switch_mode(text.next_game_mode)
	}
}

ALEXANDRIA : []Text_Data = {
	{
		lines = {
			{
				str = "You are a cowboy",
				pos = Point{ 30, 20 },
			},
			{
				str = "hunting down",
				pos = Point{ 20, 40 },
			},
			{
				str = "cultists",
				pos = Point{ 5, 60 },
			},
		},
		next_game_mode = .Text,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "You might be out",
				pos = Point{ 5, 60 },
			},
			{
				str = "of shells, but",
				pos = Point{ 20, 80 },
			},
			{
				str = "you still have",
				pos = Point{ 30, 100 },
			},
			{
				str = "your trusty",
				pos = Point{ 45, 120 },
			},
			{
				str = "-|k=a=t=a=n=a>",
				pos = Point{ 20, 140 },
			},
		},
		next_game_mode = .Map,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "Kill them all!",
				pos = Point{ 20, 30 },
			},
		},
		next_game_mode = .Game,
		palette = L1_PALETTE,
	},
	{
		lines = {
			{
				str = "GG?",
				pos = Point{ 10, 10 },
			},
		},
		next_game_mode = .GG,
		palette = MAP_PALETTE,
	},
}

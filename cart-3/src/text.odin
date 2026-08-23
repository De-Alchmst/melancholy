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
				str = "The mag\xE6 have",
				pos = Point{ 10, 30 },
			},
			{
				str = "arrived!",
				pos = Point{ 50, 40 },
			},
			{
				str = "Beware the ballz",
				pos = Point{ 20, 100 },
			},
			{
				str = "...",
				pos = Point{ 10, 115 },
			},
			{
				str = "of FIRE",
				pos = Point{ 40, 130 },
			},
		},
		next_game_mode = .Game,
		palette = L2_PALETTE,
	},
	{
		lines = {
			{
				str = "You enter the cave",
				pos = Point{ 5, 80 },
			},
		},
		next_game_mode = .Text,
		palette = L3_PALETTE,
	},
	{
		lines = {
			{
				str = "It leads to victory",
				pos = Point{ 5, 80 },
			},
		},
		next_game_mode = .Text,
		palette = L3_PALETTE,
	},
	{
		lines = {
			{
				str = "You are sure of it!",
				pos = Point{ 5, 80 },
			},
		},
		next_game_mode = .Game,
		palette = L3_PALETTE,
	},
	{
		lines = {
			{
				str = "!! D E E E M O N !!",
				pos = Point{ 5, 80 },
			},
		},
		next_game_mode = .Text,
		palette = L4_PALETTE,
	},
	{
		lines = {
			{
				str = "Beware the demonic",
				pos = Point{ 10, 50 },
			},
			{
				str = "menace!",
				pos = Point{ 90, 65 },
			},
		},
		next_game_mode = .Text,
		palette = L4_PALETTE,
	},
	{
		lines = {
			{
				str = "They do not go",
				pos = Point{ 10, 90 },
			},
			{
				str = "down so easily",
				pos = Point{ 30, 115 },
			},
		},
		next_game_mode = .Game,
		palette = L4_PALETTE,
	},
	{
		lines = {
			{
				str = "You are this close",
				pos = Point{ 10, 30 },
			},
			{
				str = "Just a little more",
				pos = Point{ 10, 90 },
			},
			{
				str = "Almost done",
				pos = Point{ 60, 140 },
			},
		},
		next_game_mode = .Game,
		palette = L5_PALETTE,
	},
	{
		lines = {
			{
				str = "The ritual has",
				pos = Point{ 25, 80 },
			},
			{
				str = "b e g o n e",
				pos = Point{ 35, 90 },
			},
		},
		next_game_mode = .Text,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "END",
				pos = Point{ 60, 70 },
			},
			{
				str = "IT",
				pos = Point{ 70, 80 },
			},
			{
				str = "NOW",
				pos = Point{ 71, 90 },
			},
		},
		next_game_mode = .Game,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "You have defeated",
				pos = Point{ 10, 50 },
			},
			{
				str = "the archmagus",
				pos = Point{ 30, 80 },
			},
			{
				str = "...",
				pos = Point{ 70, 100 },
			},
		},
		next_game_mode = .Text,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "But this is",
				pos = Point{ 30, 60 },
			},
			{
				str = "not the end",
				pos = Point{ 20, 70 },
			},
		},
		next_game_mode = .Text,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "it is never",
				pos = Point{ 40, 80 },
			},
			{
				str = "The End",
				pos = Point{ 100, 140 },
			},
		},
		next_game_mode = .Text,
		palette = MAP_PALETTE,
	},
	{
		lines = {
			{
				str = "There",
				pos = Point{ 10, 30 },
			},
			{
				str = "Is",
				pos = Point{ 50, 60 },
			},
			{
				str = "Never",
				pos = Point{ 70, 90 },
			},
			{
				str = "Peace",
				pos = Point{ 110, 120 },
			},
		},
		next_game_mode = .GG,
		palette = MAP_PALETTE,
	},
}

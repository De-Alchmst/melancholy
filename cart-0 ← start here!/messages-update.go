package main

import "cart/w4"

type MessageText struct {
	Text string
	X, Y int
	DrawColors uint16
}

type Message struct {
	Texts []MessageText
	after func()
}

var (
	IntroMessage = Message{
		Texts: []MessageText{
			{Text: "And so it begins", X: 5, Y: 10, DrawColors: 0x2},
			{Text: "Like many times\n         before", X: 20, Y: 35, DrawColors: 0x2},
			{Text: "Like many times\n          after", X: 20, Y: 70, DrawColors: 0x2},
			{Text: "What is one\n  more time?", X:52, Y: 130, DrawColors: 0x4},
		},
		after: func() {State.Status = StatusOverworld},
	}
)

func UpdateMessage() {
	for _, text := range IntroMessage.Texts {
		*w4.DRAW_COLORS = text.DrawColors
		w4.Text(text.Text, text.X, text.Y)
	}

	if pressed(KeyAction) {
		IntroMessage.after()
	}
}

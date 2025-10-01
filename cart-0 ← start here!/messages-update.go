package main

import "cart/w4"

type MessageText struct {
	Text string
	X, Y int
	DrawColors uint16
}

type MessageImage struct {
	Sprite Sprite
	X, Y int
}

type Message struct {
	Texts []MessageText
	Images []MessageImage
	After func()
}

var (
	IntroMessage = Message {
		Texts: []MessageText {
			{Text: "And so it begins", X: 5, Y: 10, DrawColors: 0x2},
			{Text: "Like many times\n         before", X: 20, Y: 35, DrawColors: 0x2},
			{Text: "Like many times\n          after", X: 20, Y: 70, DrawColors: 0x2},
			{Text: "What is one\n  more time?", X:52, Y: 130, DrawColors: 0x4},
		},
		Images: []MessageImage{},
		After: BackToOverworld,
	}
)


func BackToOverworld() {
	State.Status = StatusOverworld
}


func UpdateMessage() {
	msg := State.CurrentMessage

	for _, img := range msg.Images {
		sprite := &img.Sprite
		*w4.DRAW_COLORS = sprite.DrawColors
		w4.Blit(&sprite.Data[0], img.X, img.Y, sprite.PiceWidth, sprite.PiceHeight, sprite.Flags)
	}

	for _, text := range msg.Texts {
		*w4.DRAW_COLORS = text.DrawColors
		w4.Text(text.Text, text.X, text.Y)
	}

	if Pressed(KeyAction) {
		msg.After()
	}
}

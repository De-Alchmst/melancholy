package main

import "cart/w4"

type Pallete int
const (
	PalleteGruvboxLight Pallete = iota
)

func SetPallete(p Pallete) {
	switch p {
	case PalleteGruvboxLight:
		w4.PALETTE[0] = 0xfbf1c7 // bg
		w4.PALETTE[1] = 0x3c3836 // fg
		w4.PALETTE[2] = 0x98971a // green
		w4.PALETTE[3] = 0x458588 // blue
	}
}

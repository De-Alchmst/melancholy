package main

import "cart/w4"

type Pallete int
const (
	PalleteGruvboxLight Pallete = iota // https://github.com/morhetz/gruvbox?tab=readme-ov-file
	PalleteBlessing                   // https://lospec.com/palette-list/blessing
	PalleteRustGold                  // https://lospec.com/palette-list/rust-gold-8
)

func SetPallete(p Pallete) {
	switch p {
	case PalleteGruvboxLight:
		w4.PALETTE[0] = 0xfbf1c7 // bg
		w4.PALETTE[1] = 0x3c3836 // fg
		w4.PALETTE[2] = 0x98971a // green
		w4.PALETTE[3] = 0x458588 // blue
	
	case PalleteBlessing:
		w4.PALETTE[0] = 0xf7ffae // yellow
		w4.PALETTE[1] = 0x74569b // dark purple
		w4.PALETTE[2] = 0xd8bfd8 // light purple
		w4.PALETTE[3] = 0x96fbc7 // green

	case PalleteRustGold:
		w4.PALETTE[0] = 0xf6cd26 // gold
		w4.PALETTE[1] = 0x202020 // black
		w4.PALETTE[2] = 0x393939 // gray
		w4.PALETTE[3] = 0xac6b26 // bronze
	}
}

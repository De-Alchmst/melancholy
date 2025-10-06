package main

import "cart/w4"

func DrawRectLines(x, y int, width, height uint) {
	w4.HLine(x, y, width)
	w4.VLine(x, y, height)
	w4.HLine(x, y+int(height), width+1)
	w4.VLine(x+int(width), y, height)
}

package main

import "cart/w4"

func DrawRectLines(x, y int, width, height uint) {
	w4.HLine(x, y, width)
	w4.VLine(x, y, height)
	w4.HLine(x, y+int(height), width+1)
	w4.VLine(x+int(width), y, height)
}


func (s *Soul) Draw() {
	*w4.DRAW_COLORS = s.Sprite.DrawColors
	x := int  (s.Hitbox.X)
	y := int  (s.Hitbox.Y)
	w := uint (s.Hitbox.Width)
	h := uint (s.Hitbox.Height)

	flags := s.Sprite.Flags
	switch (s.Direction) {
	case DirUp:
		flags |= w4.BLIT_FLIP_Y
	case DirRight:
		flags |= w4.BLIT_ROTATE
	case DirLeft:
		flags |= w4.BLIT_ROTATE | w4.BLIT_FLIP_Y
	}

	w4.BlitSub(&s.Sprite.Data[0], x, y, w, h, 0, 0, s.Sprite.ArchWidth, flags)
}


func (l SoulShotList) Draw() {
	var ( // one must respect the Pascal legacy...
		x, y  int
		w, h uint
	)

	*w4.DRAW_COLORS = 0x4
	for _, s := range l {
		x = int  (s.Hitbox.X)
		y = int  (s.Hitbox.Y)
		w = uint (s.Hitbox.Width)
		h = uint (s.Hitbox.Height)

		if s.Direction == DirUp || s.Direction == DirDown {
			w4.VLine(x, y, h)
		} else {
			w4.HLine(x, y, w)
		}
	}
}

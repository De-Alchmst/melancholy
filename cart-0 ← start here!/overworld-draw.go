package main

import "cart/w4"
import "sort" // takes 4k, might write myself if run out of memory

func (l OverworldEntityList) Draw() {
	sort.Slice(l, func(i, j int) bool {
		return l[i].Hitbox.Y < l[j].Hitbox.Y
	})

	for i := range l {
		l[i].Draw()
	}
}

func (e *OverworldEntity) Draw() {
	s := e.Sprite
	x := int(e.Hitbox.X) + e.DrawOffsetX
	y := int(e.Hitbox.Y) + e.DrawOffsetY
	srcX := e.AnimationFrames[e.AnimationIndex] * s.PiceWidth
	srcY := uint(e.Direction) * s.PiceHeight

	*w4.DRAW_COLORS = s.DrawColors
	w4.BlitSub(&s.Data[0], x, y,s.PiceWidth, s.PiceHeight, srcX, srcY, s.ArchWidth, s.Flags)
}


func DrawBackground() {
	*w4.DRAW_COLORS = State.CurrentRoom.DrawColors
	
	for y := range 10 {
		for x := range 10 {
			tileID := State.CurrentRoom.Tiles[y][x]
			tile := TileAtlas[tileID]

			w4.Blit(&tile.Data[0], x*TileSize, y*TileSize, TileSize, TileSize, tile.Flags)
		}
	}
}


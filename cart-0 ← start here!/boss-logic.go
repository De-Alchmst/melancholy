package main

import (
	"cart/w4"
	"math"
)

const(
	soulSpeed = 1.4
	soulShotSpeed = 1.6
	handSpeed = 0.02
)

// a lot of garbage here...
func removeAtIndex[T any](s []T, i int) []T {
	return append(s[:i], s[i+1:]...)
}


 // let's consider just the current box dimensions and rewrite as needed
//  Moore approved! 
func soulCollidesBox(s *Soul) bool {
	x := int(s.Hitbox.X)
	y := int(s.Hitbox.Y)
	w := int(s.Hitbox.Width)
	h := int(s.Hitbox.Height)
	// ask not of the wizard to reveal his secrets...
	return x <= 40 || y <= 40 || x+w > 120 || y+h > 120
}


// probably should be factored
// anyways...
func UpdateSoulDirections(s *Soul) {
		if Held(KeyUp) {
		s.Hitbox.Y -= soulSpeed
		if soulCollidesBox(s) {
			s.Hitbox.Y += soulSpeed
		}
		s.Direction = DirUp
	}
	if Held(KeyDown) {
		s.Hitbox.Y += soulSpeed
		if soulCollidesBox(s) {
			s.Hitbox.Y -= soulSpeed
		}
		s.Direction = DirDown
	}
	if Held(KeyLeft) {
		s.Hitbox.X -= soulSpeed
		if soulCollidesBox(s) {
			s.Hitbox.X += soulSpeed
		}
		s.Direction = DirLeft
	}
	if Held(KeyRight) {
		s.Hitbox.X += soulSpeed
		if soulCollidesBox(s) {
			s.Hitbox.X -= soulSpeed
		}
		s.Direction = DirRight
	}
}


func newSoulShot(s *Soul) SoulShot {
	shot := SoulShot {
		Direction: s.Direction,
		Damage: 1,
	}

	h := s.Hitbox
	switch (s.Direction) {
	case DirUp:
		shot.Hitbox = Hitbox {
			Width: 1, Height: 2,
			X: h.X + h.Width/2, Y: h.Y-2,
		}

	case DirRight:
		shot.Hitbox = Hitbox {
			Width: 2, Height: 1,
			X: h.X+h.Width, Y: h.Y + h.Height/2,
		}

	case DirDown:
		shot.Hitbox = Hitbox {
			Width: 1, Height: 2,
			X: h.X + h.Width/2, Y: h.Y + h.Height,
		}

	case DirLeft:
		shot.Hitbox = Hitbox {
			Width: 2, Height: 1,
			X: h.X-2, Y: h.Y + h.Height/2,
		}
	}

	return shot
}


func UpdateSoulShooting(s *Soul, b *BossConfig) {
	const cooldown = 10
	if s.Cooldown > 0 {
		s.Cooldown -= 1
	}

	if Held(KeyX) && s.Cooldown <= 0 {
		s.Cooldown = cooldown
		b.SoulShots = append(b.SoulShots, newSoulShot(s))
	}
}


func MoveSoulShots(sl *SoulShotList) {
	s := *sl
	for i := range s {
		switch (s[i].Direction) {
		case DirUp:
			s[i].Hitbox.Y -= soulShotSpeed
		case DirRight:
			s[i].Hitbox.X += soulShotSpeed
		case DirDown:
			s[i].Hitbox.Y += soulShotSpeed
		case DirLeft:
			s[i].Hitbox.X -= soulShotSpeed
		}
	}
}


func KillSoulShots(l *SoulShotList) {
	for i, s := range *l {
		h := s.Hitbox
		if h.X < 0 || h.Y < 0 || h.X+h.Width > 160 || h.Y+h.Height > 160 {
			*l = removeAtIndex(*l, i)
		}
	}
}


func newHand() BossPart {
	var (
		width, height, x, y  float32
		counter float64 = 0
	)

	sprite    := BossHandSprite
	flags     := sprite.Flags
	randomDir := [4]Direction {DirUp, DirRight, DirDown, DirLeft}[GetRandomN(4)]
	switch (randomDir) {
	case DirUp:
		flags |= w4.BLIT_FLIP_Y | w4.BLIT_FLIP_X
		width  = float32(sprite.PiceWidth)
		height = float32(sprite.PiceHeight)
		x      = (160 - float32(sprite.PiceWidth)) / 2
		y      = (160)
	case DirRight:
		flags |=  w4.BLIT_ROTATE
		width  = float32(sprite.PiceHeight)
		height = float32(sprite.PiceWidth)
		x      = (    - float32(sprite.PiceHeight)) 
		y      = (160 - float32(sprite.PiceWidth)) / 2
	case DirDown:
		width  = float32(sprite.PiceWidth)
		height = float32(sprite.PiceHeight)
		x      = (160 - float32(sprite.PiceWidth)) / 2
		y      = (    - float32(sprite.ArchHeight) )
	case DirLeft:
		flags |= w4.BLIT_FLIP_Y | w4.BLIT_ROTATE | w4.BLIT_FLIP_X
		width  = float32(sprite.PiceHeight)
		height = float32(sprite.PiceWidth)
		x      = (160)
		y      = (160 - float32(sprite.PiceWidth)) / 2
	}

	part := BossPart {
		Sprite: BossHandSprite,
		Hitbox: Hitbox {Width: width, Height: height, X: x, Y:y},
		DrawOffsetX: 0, DrawOffsetY: 0,
		Flags: flags,
		DrawColors: sprite.DrawColors,

		Update: func(self *BossPart) bool {
			counter += handSpeed
			sinc := float32(math.Sin(counter)) * float32(sprite.ArchHeight)
			cosc := float32(math.Cos(counter)) * float32(sprite.ArchWidth) / 2

			switch (randomDir) {
			case DirUp:
				self.Hitbox.X = x - cosc
				self.Hitbox.Y = y - sinc
			case DirRight:
				self.Hitbox.X = x + sinc
				self.Hitbox.Y = y - cosc
			case DirDown:
				self.Hitbox.X = x + cosc
				self.Hitbox.Y = y + sinc
			case DirLeft:
				self.Hitbox.X = x - sinc
				self.Hitbox.Y = y + cosc
			}
			return counter > math.Pi
		},
	}

	return part
}


func HandleHandsPopulation(parts *BossPartList) {
	if len(*parts) == 0 {
		*parts = append(*parts, newHand())
	}
}


func UpdateHands(b *BossConfig) {
	parts := &b.BossParts

	HandleHandsPopulation(parts)

	for i := range *parts {
		if (*parts)[i].Update(&(*parts)[i]) {
			*parts = removeAtIndex(*parts, i)
		}
	}
}

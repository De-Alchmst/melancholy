package main

import (
	"cart/w4"
	"math"
)

type Vect struct {
	X, Y float32
}
		
const(
	soulSpeed = 1.4
	soulShotSpeed = 1.6
	handSpeed = 0.02
)

var (
	HandsTakenUp    = false
	HandsTakenRight = false
	HandsTakenDown  = false
	HandsTakenLeft  = false
)

// a lot of garbage here...
func RemoveAtIndex[T any](s []T, i int) []T {
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
		PlayShootingSound()
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
	i := 0
	for i < len(*l) {
		h := (*l)[i].Hitbox
		if h.X < 0 || h.Y < 0 || h.X+h.Width > 160 || h.Y+h.Height > 160 {
			*l = RemoveAtIndex(*l, i)
		} else {
			i++
		}
	}
}


func newHand(b *BossConfig) BossPart {
	var (
		width, height, x, y  float32
		counter float64 = 0
		randomDir Direction
	)

	sprite    := BossHandSprite
	flags     := sprite.Flags

	for true {
		randomDir = [4]Direction {DirUp, DirRight, DirDown, DirLeft}[GetRandomN(4)]
		switch (randomDir) {
		case DirUp:
			if HandsTakenUp { continue }
			HandsTakenUp = true

			flags |= w4.BLIT_FLIP_Y | w4.BLIT_FLIP_X
			width  = float32(sprite.PiceWidth)
			height = float32(sprite.PiceHeight)
			x      = (160 - float32(sprite.PiceWidth)) / 2
			y      = (160)

		case DirRight:
			if HandsTakenRight { continue }
			HandsTakenRight = true

			flags |=  w4.BLIT_ROTATE
			width  = float32(sprite.PiceHeight)
			height = float32(sprite.PiceWidth)
			x      = (    - float32(sprite.PiceHeight)) 
			y      = (160 - float32(sprite.PiceWidth)) / 2

		case DirDown:
			if HandsTakenDown { continue }
			HandsTakenDown = true

			width  = float32(sprite.PiceWidth)
			height = float32(sprite.PiceHeight)
			x      = (160 - float32(sprite.PiceWidth)) / 2
			y      = (    - float32(sprite.ArchHeight) )

		case DirLeft:
			if HandsTakenLeft { continue }
			HandsTakenLeft = true

			flags |= w4.BLIT_FLIP_Y | w4.BLIT_ROTATE | w4.BLIT_FLIP_X
			width  = float32(sprite.PiceHeight)
			height = float32(sprite.PiceWidth)
			x      = (160)
			y      = (160 - float32(sprite.PiceWidth)) / 2
		}
		break
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

			if counter > math.Pi {
				switch (randomDir) {
				case DirUp:    HandsTakenUp     = false
				case DirRight: HandsTakenRight  = false
				case DirDown:  HandsTakenDown   = false
				case DirLeft:  HandsTakenLeft   = false
				}
				return true
			}

			return false
		},
	}

	spawnBossAttacks(x, y, randomDir, &b.BossAttacks)

	return part
}


func HandleHandsPopulation(parts *BossPartList, b *BossConfig) {
	var handsNum int = 0
	if b.HP < 37 {
		handsNum = 3
	} else if (b.HP < 75) {
		handsNum = 2
	}

	if len(*parts) == handsNum {
		*parts = append(*parts, newHand(b))
	}
}


func UpdateHands(b *BossConfig) {
	parts := &b.BossParts

	HandleHandsPopulation(parts, b)

	i := 0
	for i < len(*parts) {
		if (*parts)[i].Update(&(*parts)[i]) {
			*parts = RemoveAtIndex(*parts, i)
		} else {
			i++
		}
	}
}


func getVectOffset(start, end float32) float32 {
	return start + float32(GetRandomN(int((end-start)*100)))/100
}


func spawnBossAttacks(startX, startY float32, direction Direction, list *BossAttackList) {
	for range 15 + GetRandomN(10) {
		primary   := -1   - getVectOffset(0, 2)
		secondary := -0.5 + getVectOffset(0, 1)

		var vect Vect
		switch (direction) {
		case DirUp:    vect = Vect { X:  secondary , Y:  primary   }
		case DirRight: vect = Vect { X: -primary   , Y:  secondary }
		case DirDown:  vect = Vect { X:  secondary , Y: -primary   }
		case DirLeft:  vect = Vect { X:  primary   , Y:  secondary }
		}


		*list = append(*list, BossAttack {
			Hitbox: Hitbox{
				X: startX - float32(BossShotSprite.ArchWidth)  /2 + float32(GetRandomN(20)),
				Y: startY - float32(BossShotSprite.ArchHeight) /2 + float32(GetRandomN(20)),
				Width:      float32(BossShotSprite.PiceWidth),
				Height:     float32(BossShotSprite.PiceHeight),
			},
		})

		attack := &(*list)[len(*list)-1]

		attack.Update = func(self *BossAttack) bool {
			attack.Hitbox.X += vect.X
			attack.Hitbox.Y += vect.Y

			return attack.Hitbox.X < -attack.Hitbox.Width  ||
						 attack.Hitbox.X > 160                   ||
						 attack.Hitbox.Y < -attack.Hitbox.Height ||
						 attack.Hitbox.Y > 160
		}

		attack.Draw = func(self *BossAttack) {
			*w4.DRAW_COLORS = BossShotSprite.DrawColors
			w4.Blit(&BossShotSprite.Data[0], int(attack.Hitbox.X), int(attack.Hitbox.Y), BossShotSprite.PiceWidth, BossFaceSprite.PiceHeight, BossShotSprite.Flags)
		}
	}
}


func BossAttackCollision(a BossAttack, b *BossConfig) bool {
	if b.Soul.Hitbox.Collides(a.Hitbox) {
		b.HP += BossHealFactor
		if b.HP > BossMaxHealth {
			b.HP = BossMaxHealth
		}
		return true
	}
	return false
}

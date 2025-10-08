package main

const(
	soulSpeed = 1.4
	soulShotSpeed = 1.6
)


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


func MoveSoulShots(s SoulShotList) {
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

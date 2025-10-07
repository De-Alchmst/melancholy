package main

const(
	soulSpeed = 1.4
)

func (s *Soul) Update() {
	updateSoulDirections(s)
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
func updateSoulDirections(s *Soul) {
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

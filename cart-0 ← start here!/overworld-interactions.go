package main

func HandleInteractions() {
	if Pressed(KeyX) {
		box := GetPlayerInteractionBox()

		// here we will also use entities from the room directly.
		// we wouldn't need to necesereeceraly, but I do feel like it...
		for _, ent := range State.CurrentRoom.Entities {
			if box.Collides(ent.Hitbox) {
				ent.OnInteract(ent)
			}
		}

		for _, evnt := range State.CurrentRoom.Events {
			if box.Collides(evnt.Hitbox) {
				evnt.OnInteract(/*no arg here I presume...*/)
			}
		}
	}
}


func GetPlayerInteractionBox() Hitbox {
	const reach = 5
	var box Hitbox
	switch Player.Direction {
	case DirUp:
		box = Hitbox {
			X: Player.Hitbox.X,
			Y: Player.Hitbox.Y - reach,
			Width: Player.Hitbox.Width,
			Height: reach,
		}
	case DirDown:
		box = Hitbox {
			X: Player.Hitbox.X,
			Y: Player.Hitbox.Y + Player.Hitbox.Height,
			Width: Player.Hitbox.Width,
			Height: reach,
		}
	case DirLeft:
		box = Hitbox {
			X: Player.Hitbox.X - reach,
			Y: Player.Hitbox.Y,
			Width: reach,
			Height: Player.Hitbox.Height,
		}
	case DirRight:
		box = Hitbox {
			X: Player.Hitbox.X + Player.Hitbox.Width,
			Y: Player.Hitbox.Y,
			Width: reach,
			Height: Player.Hitbox.Height,
		}
	}

	return box
}

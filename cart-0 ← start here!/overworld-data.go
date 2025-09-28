package main

type Direction uint
const (
	DirDown Direction = iota
	DirUp
	DirLeft
	DirRight
)

type Hitbox struct {
	X, Y, Width, Height float32
}

type OverworldEntity struct {
	Hitbox Hitbox
	DrawOffsetX, DrawOffsetY int
	AnimationFrames []uint // Order of columns in sprite sheet
	AnimationIndex int
	AnimationCountdown float32
	Sprite Sprite
	Direction Direction
}


type RoomID int
type Room struct {
	ID RoomID
	Tiles [10][10]TileAtlasTile
	Left, Right, Up, Down RoomID
	DrawColors uint16
}

const (
	TileSize = 16
)


var Player = OverworldEntity{
	DrawOffsetX: -2, DrawOffsetY: -9,
	Hitbox: Hitbox {
		X: 64, Y: 64,
		Width: 12, Height: 7,
	},
	AnimationFrames: []uint{1, 0, 1, 2},
	AnimationIndex: 0,
	AnimationCountdown: 0,
	Sprite: PlayerSprite,
	Direction: DirDown,
}

func GetRoomAtID(id RoomID) *Room {
	for i := range Rooms {
		if Rooms[i].ID == id {
			return &Rooms[i]
		}
	}
	return &Rooms[0]
}

var (
	Rooms = [...]Room {
		Room {
			ID: 0,
			Tiles: [10][10]TileAtlasTile {
				{0,0,0,0,0,1,0,0,0,0},
				{0,0,0,0,0,1,0,0,0,0},
				{0,0,1,1,1,1,1,1,1,1},
				{0,0,1,1,1,1,1,1,1,1},
				{0,0,1,1,1,1,1,1,1,1},
				{0,0,1,1,1,1,1,1,1,1},
				{0,0,1,1,1,1,1,1,1,1},
				{0,0,1,1,1,1,1,1,1,1},
				{0,0,0,0,0,0,0,0,0,0},
				{0,0,0,0,0,0,0,0,0,0},
			},
			Left: 0, Right: 1, Up: 0, Down: 0,
			DrawColors: 0x41,
		},
		Room {
			ID: 1,
			Tiles: [10][10]TileAtlasTile {
				{0,0,0,0,0,0,0,0,0,0},
				{0,0,0,0,0,0,0,0,0,0},
				{1,1,1,1,1,1,0,0,0,0},
				{1,1,1,1,1,1,0,1,1,0},
				{1,1,1,1,1,1,0,1,1,1},
				{1,1,1,1,1,1,0,1,1,0},
				{1,1,1,1,1,1,0,0,0,0},
				{1,1,1,1,1,1,0,0,0,0},
				{0,0,0,0,1,1,0,0,0,0},
				{0,0,0,0,1,1,0,0,0,0},
			},
			Left: 0, Right: 1, Up: 0, Down: 0,
			DrawColors: 0x41,
		},
	}
)

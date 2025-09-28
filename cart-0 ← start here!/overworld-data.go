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

////////////////////

// hear me out
// This implements a list of RoomEntries, each of which holds a room ID and
// Either a Room, or a function that generates a Room.
// This allows to have rooms generate dynamically based on game state, but see
// the ID without generating the room first.

type TilesMap [10][10]TileAtlasTile
type RoomID int
type Room struct {
	Tiles TilesMap
	Left, Right, Up, Down RoomID
	DrawColors uint16
}

type RoomMaker func() *Room
type RoomGetter interface {
	GetRoom() *Room
}

func (e *Room)     GetRoom() *Room { return e   }
func (f RoomMaker) GetRoom() *Room { return f() }

type RoomListEntry struct {
	ID RoomID
	Value RoomGetter
}
func (e *RoomListEntry) Room() *Room { return e.Value.GetRoom() }


////////////////////


func GetRoomAtID(id RoomID) *Room {
	for i := range RoomEntries {
		if RoomEntries[i].ID == id {
			return RoomEntries[i].Room()
		}
	}
	return RoomEntries[0].Room()
}


var (
	RoomEntries = [...]RoomListEntry {
		RoomListEntry {
			ID: 0,
			Value: &Room {
				Tiles: TilesMap {
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
		},
		RoomListEntry {
			ID: 1,
			Value: &Room {
				Tiles: TilesMap {
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
				Left: 0, Right: 2, Up: 0, Down: 2,
				DrawColors: 0x41,
			},
		},
		RoomListEntry {
			ID: 2,
			Value: &Room {
				Tiles: TilesMap {
					{0,0,0,0,1,1,0,0,0,0},
					{0,0,1,1,1,1,0,0,0,0},
					{0,0,1,1,1,1,0,0,0,0},
					{0,0,1,1,1,1,0,0,0,0},
					{1,1,1,1,1,1,1,1,0,0},
					{0,0,0,0,0,1,1,1,1,0},
					{0,0,0,0,0,1,1,1,1,0},
					{0,0,0,0,0,1,1,1,1,0},
					{0,0,0,0,0,1,1,1,0,0},
					{0,0,0,0,0,0,0,0,0,0},
				},
				Left: 1, Right: 0, Up: 1, Down: 0,
				DrawColors: 0x41,
			},
		},
	}
)

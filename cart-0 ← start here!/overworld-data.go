package main

type Direction uint
const (
	DirDown Direction = iota
	DirUp
	DirLeft
	DirRight
)

const (
	TileSize = 16
)

type Hitbox struct {
	X, Y, Width, Height float32
}

func (b1 Hitbox) Collides(b2 Hitbox) bool {
	return b1.X < b2.X+b2.Width && b1.X+b1.Width > b2.X && b1.Y < b2.Y+b2.Height && b1.Y+b1.Height > b2.Y
}

type OverworldEntityList []*OverworldEntity
type OverworldEntity struct {
	Hitbox Hitbox
	DrawOffsetX, DrawOffsetY int
	AnimationFrames []uint // Order of columns in sprite sheet
	AnimationIndex int
	AnimationCountdown float32
	Sprite Sprite
	Direction Direction
	OnInteract func(self *OverworldEntity)
	Data any
}

func (e *OverworldEntity) NextFrame() {
	e.AnimationIndex = (e.AnimationIndex + 1) % len(e.AnimationFrames)
}

func EntityDoNothing (self *OverworldEntity) {}


var (
	Player = OverworldEntity{
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
		OnInteract: EntityDoNothing,
		Data: nil,
	}
)

////////////////////

// hear me out
// This implements a list of RoomEntries, each of which holds a room ID and
// Either a Room, or a function that generates a Room.
// This allows to have rooms generate dynamically based on game state, but see
// the ID without generating the room first.

type PositionalEventList []PositionalEvent
type PositionalEvent struct {
	Hitbox Hitbox
	OnInteract func()
}

type TilesMap [10][10]TileAtlasTile
type RoomID int
type Room struct {
	Tiles TilesMap
	Left, Right, Up, Down RoomID
	DrawColors uint16
	Entities OverworldEntityList
	Events PositionalEventList
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


func tileToPos(tile int) float32 {
	return float32(tile * TileSize)
}

func TileToHitbox(x, y float32) Hitbox {
	return Hitbox {
		X     : x*TileSize,
		Y     : y*TileSize,
		Width :   TileSize,
		Height:   TileSize,
	}
}


var (
	RoomEntries = [...]RoomListEntry {
		RoomListEntry {
			ID: 0,
			Value: RoomMaker(func() *Room {
				return &Room {
					Tiles: TilesMap {
						{0,0,0,0,0,1,0,0,0,0},
						{0,0,0,0,0,2,0,0,0,0},
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
					Entities: OverworldEntityList{},
					Events: PositionalEventList{},
				}
			}),
		},

		RoomListEntry {
			ID: 1,
			Value: RoomMaker(func() *Room {
				gotKey, ok := State.Events["got_first_key"]
				if !ok {
					gotKey = 0
				}
			
				return &Room {
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
					Entities : OverworldEntityList{
						&OverworldEntity {
							DrawOffsetX: -1, DrawOffsetY: -9,
							Hitbox: Hitbox {
								X: tileToPos(7)+1, Y: tileToPos(4)+9,
								Width: 14, Height: 7,
							},
							AnimationFrames: []uint{0, 1},
							AnimationIndex: int(gotKey),
							AnimationCountdown: 0,
							Sprite: KeyholderSprite,
							Direction: DirDown,
							OnInteract: func(self *OverworldEntity) {
								self.AnimationIndex = 1
								State.Events["got_first_key"] = 1
							},
							Data: nil,
						},
					},
					Events: PositionalEventList{},
				}
			}),
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
				Entities : OverworldEntityList {},
				Events: PositionalEventList {
					PositionalEvent {
						Hitbox: TileToHitbox(7, 6),
						OnInteract: func() {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "Do you remember\n what happened\n last week?",
							  		X: 20, Y: 20,  DrawColors: 0x2 },
									{ Text: "Did it even\n happen then?",
								  	X: 50, Y: 100, DrawColors: 0x4 },
								},
								Images: []MessageImage{},
								After: BackToOverworld,
							}//, <- I hate you!
						},
					},
				},
			},
		},
	}
)

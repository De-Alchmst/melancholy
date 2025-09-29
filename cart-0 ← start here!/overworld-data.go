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
	Pallete Pallete
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
				room := &Room {
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
					Left: 0, Right: 1, Up: 3, Down: 0,
					Pallete: PalleteGruvboxLight,
					DrawColors: 0x41,
					Entities: OverworldEntityList{},
					Events: PositionalEventList{},
				}

				// cannot refer to 'room' when defined inside as it's part.
				// Go's FP could certainly be improved...
				room.Events = PositionalEventList {
					PositionalEvent {
						Hitbox: TileToHitbox(5,1),
						OnInteract: func () {
							if EventRegistered("got_first_key") {
								RegisterEvent("first_unlocked", 1)
								room.Tiles[1][5] = 1

							} else {
								State.Status = StatusMessage
								State.CurrentMessage = Message {
									Texts: []MessageText {
										{ Text: "You are locked here",
											X: 5, Y: 20,  DrawColors: 0x2 },
										{ Text: "Why do you want\n to leave anyways?",
										  // It's not like there's anything to see anyways
										 // it's good here as is, you don't need to go anywhere
										// you have here everything you'llllllllllllllllllllll
									 // ever need
									// no really, don't!
								 //             don't!1!
											X: 5, Y: 100, DrawColors: 0x4 },
									},
									Images: []MessageImage{},
									After: BackToOverworld,
								}
							}
						},
					},
				}
				if EventRegistered("first_unlocked") {
					room.Tiles[1][5] = 1
				}
				return room
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
						{1,1,1,4,1,1,0,0,0,0},
						{1,1,1,1,1,1,0,1,1,0},
						{1,1,1,1,1,1,0,1,1,1},
						{1,1,1,1,1,1,0,1,1,0},
						{1,1,1,1,1,1,0,0,0,0},
						{1,1,1,1,1,1,0,0,0,0},
						{0,0,0,0,1,1,0,0,0,0},
						{0,0,0,0,1,1,0,0,0,0},
					},
					Left: 0, Right: 2, Up: 0, Down: 2,
					Pallete: PalleteGruvboxLight,
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
								RegisterEvent("got_first_key", 1)
							},
							Data: nil,
						},
					},
					Events: PositionalEventList {
						PositionalEvent {
							Hitbox: TileToHitbox(3,2),
							OnInteract: func() {
								State.Status = StatusMessage
								State.CurrentMessage = Message {
									Texts: []MessageText {
										{ Text: "This is your\n computer.",
											X: 5, Y: 20,  DrawColors: 0x2 },
										{ Text: "It makes you happy.",
											X: 10, Y: 60, DrawColors: 0x2 },
										{ Text: "It makes YOU able\n to feel this\n\n  [EXPERIENCE]!",
											X: 5, Y: 90, DrawColors: 0x3 },
									},
									Images: []MessageImage{},
									After: BackToOverworld,
								}
							},
						},
					},
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
				Pallete: PalleteGruvboxLight,
				DrawColors: 0x41,
				Entities: OverworldEntityList {},
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

		RoomListEntry {
			ID: 3, // the outside
			Value: &Room {
				Tiles: TilesMap {
					{0,0,0,0,0,3,0,0,0,0},
					{1,1,1,1,1,1,1,0,0,0},
					{1,1,1,1,1,1,1,0,0,0},
					{0,0,0,0,1,1,1,0,0,0},
					{0,0,0,0,1,1,1,1,1,2},
					{0,0,0,0,1,1,1,0,0,0},
					{0,0,0,0,1,1,1,0,0,0},
					{0,0,0,0,1,1,1,0,0,0},
					{0,0,0,0,0,1,0,0,0,0},
					{0,0,0,0,0,1,0,0,0,0},
				},
				Left: 0, Right: 0, Up: 4, Down: 0,
				Pallete: PalleteGruvboxLight,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		RoomListEntry {
			ID: 4, // outside what?
			Value: &Room {
				Tiles: TilesMap {
					{0,0,0,0,0,0,0,0,3,0},
					{0,3,3,3,3,0,3,3,3,0},
					{0,3,0,3,0,3,3,3,0,0},
					{0,0,3,3,0,3,0,3,3,3},
					{0,0,3,3,3,3,0,3,0,0},
					{0,0,3,0,0,3,0,3,3,0},
					{0,3,3,3,0,3,3,0,3,0},
					{0,3,0,3,3,3,3,0,3,0},
					{0,3,3,3,0,3,0,0,3,1},
					{0,0,0,0,0,1,0,0,0,0},
				},
				Left: 0, Right: 5, Up: 5, Down: 3,
				Pallete: PalleteRustGold,
				DrawColors: 0x21,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		RoomListEntry {
			ID: 5, // outside what?
			Value: &Room {
				Tiles: TilesMap {
					{0,0,0,0,0,0,0,0,0,0},
					{0,0,0,0,0,1,1,1,0,0},
					{0,3,0,0,1,1,1,1,1,0},
					{3,3,0,1,1,1,1,1,1,0},
					{0,0,0,1,1,1,1,1,1,0},
					{0,0,0,1,1,1,1,1,1,0},
					{0,0,0,1,1,1,1,1,1,0},
					{0,0,0,0,1,1,1,1,1,0},
					{1,1,0,0,0,1,1,0,1,0},
					{0,0,0,0,0,0,0,0,1,0},
				},
				Left: 4, Right: 4, Up: 4, Down: 4,
				Pallete: PalleteRustGold,
				DrawColors: 0x21,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{
					PositionalEvent {
						Hitbox: TileToHitbox(1,1),
						OnInteract: func() {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "Are you Lost?",
							  		X: 20, Y: 20,  DrawColors: 0x2 },
								},
								Images: []MessageImage{},
								After: BackToOverworld,
							}
						},
					},
				},
			},
		},
	}
)

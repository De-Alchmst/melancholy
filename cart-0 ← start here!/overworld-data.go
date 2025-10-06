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

 // Oh, look! Math!
//  https://developer.mozilla.org/en-US/docs/Games/Techniques/2D_collision_detection#Axis-Aligned_Bounding_Box
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
    // this is so called /dev/null function
   //  it takes a bunch of data, pretends to be super important, but it actually
  //   does absolutely nothing!
 //
//     very relatable


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
		{
			ID: 0,
			Value: RoomMaker(func() *Room {
				room := &Room {
					Tiles: TilesMap {
						{0,3,0,0,0,1,0,0,0,0},
						{3,3,0,0,0,2,0,0,0,0},
						{0,0,1,1,1,1,1,1,1,1},
						{0,0,1,1,1,1,1,1,1,1},
						{0,0,1,1,1,1,1,1,1,1},
						{0,0,1,1,1,1,1,1,1,1},
						{0,0,1,1,1,1,1,1,1,1},
						{0,0,1,1,1,1,1,1,1,1},
						{0,0,0,0,0,0,0,0,0,0},
						{0,0,0,0,0,0,0,0,0,0},
					},
					Left: 11, Right: 1, Up: 3, Down: 0,
					Pallete: PalleteGruvboxLight,
					DrawColors: 0x41,
					Entities: OverworldEntityList{},
					Events: PositionalEventList{},
				}

				// cannot refer to 'room' when defined inside as it's part.
				// Go's FP could certainly be improved...
				room.Events = PositionalEventList {
					{
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

		{
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
						{
							DrawOffsetX: -1, DrawOffsetY: -13,
							Hitbox: Hitbox {
								X: tileToPos(7)+1, Y: tileToPos(4)+9,
								Width: 14, Height: 3,
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
						{
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

		{
			ID: 2,
			Value: &Room {
				Tiles: TilesMap {
					{0,0,0,0,1,1,0,0,0,0},
					{0,0,5,1,1,1,0,0,0,0},
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
					{
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
					}, {
						Hitbox: TileToHitbox(2, 1),
						OnInteract: func() {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "This is your\n fridge.",
										X: 5, Y: 20,  DrawColors: 0x2 },
									{ Text: "It has probably\n been runing since\n the last regime.",
										X: 5, Y: 55,  DrawColors: 0x2 },
									{ Text: "It's been\n refilling itself\n less and less\n lately tho...",
										X: 5, Y: 100, DrawColors: 0x2 },
								},
								Images: []MessageImage{},
								After: BackToOverworld,
							}
						},
					},
				},
			},
		},

		{
			ID: 3, // the outside
			Value: &Room {
				Tiles: TilesMap {
					{0,0,0,0,0,3,0,0,0,0},
					{1,1,1,1,1,1,1,0,0,0},
					{1,1,1,1,1,1,1,0,0,0},
					{0,0,0,0,1,1,1,0,0,0},
					{0,0,0,0,1,1,1,1,1,2},
					{0,0,0,0,1,1,1,0,0,0},
					{3,3,0,0,1,1,1,0,0,0},
					{3,3,0,0,1,1,1,0,0,0},
					{3,3,0,0,0,1,0,0,0,0},
					{0,3,0,0,0,1,0,0,0,0},
				},
				Left: 6, Right: 0, Up: 4, Down: 0,
				Pallete: PalleteGruvboxLight,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
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

		{
			ID: 5, // outside of what?
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
				Entities: OverworldEntityList{
					{
						DrawOffsetX: -2, DrawOffsetY: -10,
						Hitbox: Hitbox {
							X: tileToPos(6)+2, Y: tileToPos(4)+10,
							Width: 11, Height: 6,
						},
						AnimationFrames: []uint{0},
						AnimationIndex: 0,
						AnimationCountdown: 0,
						Sprite: MagicianSprite,
						Direction: DirDown,
						OnInteract: func(self *OverworldEntity) {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "I am the magician.\n You seem lost\n out there child.",
										X: 5, Y: 55,  DrawColors: 0x2 },
									{ Text: "//HE// is always\n watching.\n //HE// will take\n care of you.",
										X: 5, Y: 100, DrawColors: 0x2 },
								},
								Images: []MessageImage{
									{
										Sprite: MagicianFaceSprite,
										X: 5,
										Y: 10,
									},
								},
								After: func() {
									State.CurrentMessage = Message {
										Texts: []MessageText {
											{ Text: "You have not lost\n //HIS// favor, did\n you?",
										   X: 5, Y: 55, DrawColors: 0x2,},
											{ Text: "Terrible fate\n awaits those that\n do, so make sure\n to keep //HIM//\n on your side.",
												X: 5, Y: 100, DrawColors: 0x2 },
										},
										Images: []MessageImage{
											{
												Sprite: MagicianFaceSprite,
												X: 5,
												Y: 10,
											},
										},
										After: func() {
											State.CurrentMessage = Message {
												Texts: []MessageText {
													{ Text: "I am the magician.",
													  X: 10, Y: 65, DrawColors: 0x2,},
													{ Text: "I am the darkness.",
														X: 10, Y: 80, DrawColors: 0x3 },
													{ Text: "I am //HIS// will.",
														X: 10, Y: 95, DrawColors: 0x4 },
												},
												Images: []MessageImage{
													{
														Sprite: MagicianFaceSprite,
														X: 64,
														Y: 20,
													},
												},
												After: BackToOverworld,
											}
										},
									}
								},
							}
						},
						Data: nil,
					},
				},
				Events: PositionalEventList{
					{
						Hitbox: TileToHitbox(1,1),
						OnInteract: func() {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "Are you Lost?",
							  		X: 25, Y: 20,  DrawColors: 0x2 },
									{ Text: "Shouldn't you be\n  somewhere else?",
							  		X: 10, Y: 130,  DrawColors: 0x2 },
								},
								Images: []MessageImage{
									{
										Sprite: TalkToTheHandSprite,
										X: 56,
										Y: 65,
									},
								},
								After: BackToOverworld,
							}
						},
					},
				},
			},
		},

		{
			ID: 6,
			Value: RoomMaker(func() *Room {
				room := &Room{
					Tiles: TilesMap {
						{0,0,0,2,0,0,2,0,0,0},
						{1,1,1,1,1,1,1,1,1,1},
						{1,1,1,1,1,1,1,1,1,1},
						{0,0,0,0,1,1,0,0,0,0},
						{0,0,0,0,1,1,0,0,0,0},
						{0,0,0,0,1,1,0,0,0,0},
						{1,0,0,0,1,1,0,0,0,1},
						{1,1,2,1,1,1,1,2,1,1},
						{1,0,0,0,1,1,0,0,0,1},
						{0,0,0,0,0,0,0,0,0,0},
						 // ↑
						// but what is here tho???
					},
					Left: 7, Right: 3, Up: 0, Down: 0,
					Pallete: PalleteGruvboxLight,
					DrawColors: 0x31,
					Entities: OverworldEntityList{},
					Events: PositionalEventList{},
				}

				room.Events = PositionalEventList{
					{
					Hitbox: TileToHitbox(4, 8),
						OnInteract: func() {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "It feels safe here.",
										X: 5, Y: 80,  DrawColors: 0x2 },
								},
								Images: []MessageImage{},
								After: BackToOverworld,
							}
						},
					}, {
						Hitbox: TileToHitbox(7,7),
						OnInteract: func () {
							if EventRegistered("got_second_key") {
								RegisterEvent("second_6_unlocked", 1)
								room.Tiles[7][7] = 1

							} else {
								State.Status = StatusMessage
								State.CurrentMessage = Message {
									Texts: []MessageText {
										{ Text: "It's a door...",
											X: 20, Y: 40,  DrawColors: 0x2 },
										{ Text: "What else can\n I say?",
											X: 10, Y: 100, DrawColors: 0x2 },
									},
									Images: []MessageImage{},
									After: BackToOverworld,
								}
							}
						},
					},
				}
				if EventRegistered("second_6_unlocked") {
					room.Tiles[7][7] = 1
				}
				return room
			}),
		},

		{
			ID: 7,
			Value: &Room{       // but who does live here?
				Tiles: TilesMap {//↓ 
					{0,0,0,0,2,0,0,0,2,0},
					{0,0,0,1,1,1,1,1,1,1},
					{0,0,0,1,1,1,1,1,1,1},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
				},
				Left: 0, Right: 11, Up: 0, Down: 8,
				Pallete: PalleteGruvboxLight,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 8,
			Value: &Room{
				Tiles: TilesMap {
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,6,6,6,0,0,0,0},
					{0,0,0,1,1,1,1,1,1,1},
					{3,3,3,1,1,1,1,1,1,1},
					{0,0,0,0,2,0,0,0,2,0},
				},
				Left: 6, Right: 9, Up: 7, Down: 0,
				Pallete: PalleteGruvboxLight,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 9,
			Value: &Room{
				Tiles: TilesMap {
					{0,0,0,0,6,6,6,0,0,0},
					{0,0,0,0,6,6,6,0,0,0},
					{0,0,0,0,6,6,6,0,0,0},
					{0,0,0,0,6,6,6,0,0,0},
					{0,0,0,0,6,6,6,0,0,0},
					{0,0,0,0,6,6,6,0,0,0},
					{0,0,0,0,6,6,6,0,0,0},
					{1,1,1,1,1,1,1,0,0,0},
					{1,1,1,1,1,1,1,0,0,0},
					{0,2,0,0,0,2,0,0,0,0},
				}, //        ↑ you never was there, now were you?
				Left: 8, Right: 0, Up: 10, Down: 0,
				Pallete: PalleteGruvboxLight,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 10,
			Value: RoomMaker(func() *Room {
				room := &Room{
					Tiles: TilesMap {
						{0,2,0,0,0,2,0,0,0,0},
						{1,1,1,1,1,1,1,0,0,0},
						{1,1,1,1,1,1,1,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
						{0,0,0,0,6,6,6,0,0,0},
					},
					Left: 7, Right: 0, Up: 13, Down: 9,
					Pallete: PalleteGruvboxLight,
					DrawColors: 0x31,
					Entities: OverworldEntityList{},
					Events: PositionalEventList{ },
				}

				room.Events = PositionalEventList {
						{
						Hitbox: TileToHitbox(5,0),
						OnInteract: func () {
							if EventRegistered("got_second_key") {
								RegisterEvent("second_10_unlocked", 1)
								room.Tiles[0][5] = 1
							}
						},
					},
				}
				if EventRegistered("second_10_unlocked") {
					room.Tiles[0][5] = 1
				}
				return room
			}),
		},

		{
			ID: 11,
			Value: RoomMaker(func() *Room {
				gotKey, ok := State.Events["got_second_key"]
				if !ok {
					gotKey = 0
				}

				room := &Room {
					Tiles: TilesMap {
						{0,0,0,0,0,0,0,0,0,0},
						{1,1,1,1,1,1,1,0,1,1},
						{1,1,1,1,1,1,1,0,6,0},
						{0,0,0,0,6,6,0,0,6,0},
						{0,0,0,0,6,6,0,0,6,0},
						{0,0,0,0,6,6,6,0,1,0},
						{0,0,0,0,0,6,6,0,0,0},
						{0,0,0,0,0,6,0,0,0,0},
						{0,0,0,0,6,6,0,0,0,0},
						{0,0,0,0,0,2,0,0,0,0},
					},
					Left: 7, Right: 0, Up: 0, Down: 12,
					Pallete: PalleteGruvboxLight,
					DrawColors: 0x31,
					Entities: OverworldEntityList{
						{
							DrawOffsetX: -1, DrawOffsetY: -13,
							Hitbox: Hitbox {
								X: tileToPos(8)+1, Y: tileToPos(5)+7,
								Width: 14, Height: 3,
							},
							AnimationFrames: []uint{0, 1},
							AnimationIndex: int(gotKey),
							AnimationCountdown: 0,
							Sprite: KeyholderSpriteRev,
							Direction: DirDown,
							OnInteract: func(self *OverworldEntity) {
								self.AnimationIndex = 1
								RegisterEvent("got_second_key", 1)
							},
							Data: nil,
						},
					},
					Events: PositionalEventList{},
				}

				room.Events = PositionalEventList {
					{
						Hitbox: TileToHitbox(5,9),
						OnInteract: func () {
							if EventRegistered("got_second_key") {
								RegisterEvent("second_11_unlocked", 1)
								room.Tiles[9][5] = 1

							} else {
								State.Status = StatusMessage
								State.CurrentMessage = Message {
									Texts: []MessageText {
										{ Text: "The door doesn't\n want to let \n you out.",
											X: 15, Y: 40,  DrawColors: 0x4 },
										{ Text: "It's the door\n I swear...",
											X: 15, Y: 100, DrawColors: 0x4 },
									},
									Images: []MessageImage{},
									After: BackToOverworld,
								}
							}
						},
					},
				}
				if EventRegistered("second_11_unlocked") {
					room.Tiles[9][5] = 1
				}
				return room
			}),
		},

		{
			ID: 12,
			Value: &Room{ /// outside!!
				Tiles: TilesMap { // you finally touched grass!!!
					{7,7,7,7,7,1,7,7,7,7},
					{7,7,7,7,7,1,7,7,7,7},
					{8,8,8,8,8,9,8,8,8,0},
					{9,9,9,9,9,9,8,8,8,0},
					{9,9,9,9,9,9,8,8,8,0},
					{8,8,8,8,8,8,8,8,8,0},
					{8,8,8,8,8,8,8,8,8,0},
					{8,8,8,8,8,8,8,8,8,0},
					{7,8,8,8,8,8,8,8,8,0},
					{7,0,0,0,0,0,0,0,0,0},
				},
				Left: 14, Right: 0, Up: 11, Down: 0,
				Pallete: PalleteBlessing,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 13,
			Value: &Room{
				Tiles: TilesMap {
					{0,0,0,0,0,0,0,0,0,0},
					{0,0,0,0,8,8,8,0,0,0},
					{0,0,0,8,8,8,8,8,0,0},
					{0,0,0,8,8,8,8,8,0,0},
					{0,0,8,8,8,8,8,8,8,0},
					{0,0,8,8,8,8,8,8,8,0},
					{0,0,8,8,8,8,8,8,8,0},
					{0,0,0,8,8,8,8,8,0,0},
					{0,0,0,8,8,1,8,8,0,0},
					{0,0,0,0,0,1,0,0,0,0},
				},
				Left: 0, Right: 0, Up: 0, Down: 10,
				Pallete: PalleteRustGold,
				DrawColors: 0x21,
				Entities: OverworldEntityList{
					{
						DrawOffsetX: 0, DrawOffsetY: -6,
						Hitbox: Hitbox {
							X: tileToPos(5)+1, Y: tileToPos(4)+9,
							Width: 14, Height: 6,
						},
						AnimationFrames: []uint{0},
						AnimationIndex: 0,
						AnimationCountdown: 0,
						Sprite: BenchSprite,
						Direction: DirDown,
						OnInteract: EntityDoNothing,
						Data: nil,
					},
				},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 14,
			Value: &Room{
				Tiles: TilesMap {
					{7,7,7,7,7,7,7,7,7,7},
					{7,7,7,7,7,7,7,7,7,7},
					{8,8,8,8,8,8,8,8,8,8},
					{9,9,9,9,9,9,9,9,9,9},
					{9,9,9,9,9,9,9,9,9,9},
					{8,8,8,9,9,8,8,8,8,8},
					{8,8,8,9,9,8,8,8,8,8},
					{8,8,8,9,9,8,8,8,8,8},
					{7,7,8,9,9,8,8,7,7,7},
					{7,7,8,9,9,8,8,7,7,7},
				},
				Left: 14, Right: 12, Up: 0, Down: 15,
				Pallete: PalleteBlessing,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 15,
			Value: &Room{
				Tiles: TilesMap {
					{7,7,8,9,9,8,8,7,7,7},
					{7,7,8,9,9,8,8,7,7,7},
					{0,8,8,9,9,8,8,7,7,7},
					{0,8,8,9,9,9,8,8,8,0},
					{0,8,8,9,9,9,8,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{7,7,7,8,9,9,8,8,8,0},
					{7,7,7,8,9,9,8,8,8,0},
					{7,7,7,8,9,9,8,8,8,0},
				},
				Left: 0, Right: 0, Up: 14, Down: 16,
				Pallete: PalleteBlessing,
				DrawColors: 0x31,
				Entities: OverworldEntityList{},
				Events: PositionalEventList{},
			},
		},

		{
			ID: 16,
			Value: &Room{
				Tiles: TilesMap {
					{7,7,7,8,9,9,8,8,8,0},
					{7,7,7,8,9,9,8,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{0,8,8,9,9,9,9,8,8,0},
					{0,8,8,9,9,9,9,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{0,8,8,8,9,9,8,8,8,0},
					{0,0,0,0,2,2,0,0,0,0},
				},
				Left: 0, Right: 0, Up: 15, Down: 16,
				Pallete: PalleteBlessing,
				DrawColors: 0x31,
				Entities: OverworldEntityList{
					{
						DrawOffsetX: -3, DrawOffsetY: -6,
						Hitbox: Hitbox {
							X: tileToPos(5)+3-8, Y: tileToPos(5)+6-11,
							Width: 10, Height: 9,
						},
						AnimationFrames: []uint{0},
						AnimationIndex: 0,
						AnimationCountdown: 0,
						Sprite: AdversarySprite,
						Direction: DirDown,
						OnInteract: func(self *OverworldEntity) {
							State.Status = StatusMessage
							SetPallete(PalleteRustGold)
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "I Am The\n Forgotten Soul",
										X: 5, Y: 10,  DrawColors: 0x2 },
									{ Text: "You Shall Not Pass\n The Gate!",
										X: 5, Y: 130, DrawColors: 0x2 },
								},
								Images: []MessageImage{
									{ Sprite: BossFaceSprite, X: 57, Y: 47 },
								},
								After: func() {
									State.Status = StatusBoss
								},
							}
						},
						Data: nil,
					},
				},
				Events: PositionalEventList{
					{
						Hitbox: Hitbox{
							X: tileToPos(4), Y: tileToPos(9),
							Width: 32, Height: 16,
						},
						OnInteract: func () {
							State.Status = StatusMessage
							State.CurrentMessage = Message {
								Texts: []MessageText {
									{ Text: "A menacing presence\n prevents you from\n proceeding...",
								  	X: 5, Y: 70,  DrawColors: 0x2 },
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

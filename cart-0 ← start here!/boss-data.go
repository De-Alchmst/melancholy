package main

import "cart/w4"

type Soul struct {
	Hitbox Hitbox
	Sprite Sprite
}


type BossPartList []BossPart
type BossPart struct {
	Hitbox Hitbox
	DrawOffsetX, DrawOffsetY int
	Sprite Sprite
	Update func(self *BossPart)
	Data any
}

type BossAttackList []BossAttack
type BossAttack struct {
	Hitbox Hitbox
	Draw func(self *BossAttack)    // not all attacks need to be sprites
	Update func(self *BossAttack) //  I mean like, beams are pretty cool, right?
	Data any
}

type BossConfig struct {
	Pallete Pallete
	HP int
	Soul Soul
	BossParts BossPartList
	BossAttacks BossAttackList
	Draw func(self *BossConfig)
	Update func(self *BossConfig)
}


func (s Soul) Draw() {
	*w4.DRAW_COLORS = s.Sprite.DrawColors
	x := int  (s.Hitbox.X)
	y := int  (s.Hitbox.Y)
	w := uint (s.Hitbox.Width)
	h := uint (s.Hitbox.Height)

	w4.BlitSub(&s.Sprite.Data[0], x, y, w, h, 0, 0, s.Sprite.ArchWidth, s.Sprite.Flags)
}


var (
	TheLostDarknessBoss = BossConfig {
		Pallete: PalleteRustGold,
		HP: 100,
		Soul: Soul {
			Hitbox: Hitbox {
				X: 78, Y: 78,
				Width: 7, Height: 7,
			},
			Sprite: SoulSprite,
		},
		BossParts: BossPartList{},
		BossAttacks: BossAttackList{},
		Draw: func(self *BossConfig) {
			*w4.DRAW_COLORS = BossFaceSprite.DrawColors
			w4.Blit(&BossFaceSprite.Data[0], 57, 47, uint(BossFaceSprite.PiceWidth), uint(BossFaceSprite.PiceHeight), BossFaceSprite.Flags)

			*w4.DRAW_COLORS = 0x2
			DrawRectLines(40, 40, 80, 80)

			self.Soul.Draw()
		},
		Update: func(self *BossConfig) {
		},
	}

	CurrentBossData = TheLostDarknessBoss
)

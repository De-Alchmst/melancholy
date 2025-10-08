package main

import "cart/w4"

type Soul struct {
	Hitbox Hitbox
	Sprite Sprite
	Direction Direction
	Cooldown int
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

type SoulShotList []SoulShot
type SoulShot struct {
	Direction	Direction
	Hitbox Hitbox
	Damage int
}

type BossConfig struct {
	Pallete Pallete
	HP int
	Soul Soul
	BossParts BossPartList
	BossAttacks BossAttackList
	SoulShots SoulShotList
	Draw func(self *BossConfig)
	Update func(self *BossConfig)
}

var (
	TheForgottenSoulBoss = BossConfig {
		Pallete: PalleteRustGold,
		HP: 100,
		Soul: Soul {
			Hitbox: Hitbox {
				X: 78, Y: 78,
				Width: 7, Height: 7,
			},
			Sprite: SoulSprite,
			Cooldown: 0,
		},
		BossParts: BossPartList{},
		BossAttacks: BossAttackList{},
		SoulShots: SoulShotList{},
		Draw: func(self *BossConfig) {
			*w4.DRAW_COLORS = BossFaceSprite.DrawColors
			w4.Blit(&BossFaceSprite.Data[0], 57, 47, uint(BossFaceSprite.PiceWidth), uint(BossFaceSprite.PiceHeight), BossFaceSprite.Flags)

			*w4.DRAW_COLORS = 0x2
			DrawRectLines(40, 40, 80, 80)

			self.Soul.Draw()
			self.SoulShots.Draw()
		},
		
		Update: func(self *BossConfig) {
			self.Soul.Update(self)
			self.BossParts.Update(self)
			self.BossAttacks.Update(self)
			self.SoulShots.Update(self)
			
		},
	}
	CurrentBossData = TheForgottenSoulBoss
)

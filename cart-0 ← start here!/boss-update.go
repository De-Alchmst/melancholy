package main

func UpdateBoss() {
	CurrentBossData.Update(&CurrentBossData)
	CurrentBossData. Draw (&CurrentBossData)
}


func (s *Soul) Update(b *BossConfig) {
	UpdateSoulDirections(s)
	UpdateSoulShooting(s, b)
	UpdateMusic()
}


func (p *BossPartList) Update(b *BossConfig) {
	for i := range *p {
		(*p)[i].Update(&(*p)[i])
	}
}


func (a *BossAttackList) Update(b *BossConfig) {
	i := 0
	for i < len(*a) {
		if BossAttackCollision((*a)[i], b) || (*a)[i].Update(&(*a)[i]) {
			*a = RemoveAtIndex(*a, i)
		} else {
			i++
		}

	}
}


func (s *SoulShotList) Update(b *BossConfig) {
	MoveSoulShots(s)
	KillSoulShots(s)
}

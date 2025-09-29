package main

func UpdateOverworld() {
	MovePlayer()
	HandleInteractions()
	DrawBackground()
	State.OverworldEntsWithPlayer.Draw()
}

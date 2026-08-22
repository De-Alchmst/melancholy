package main

Game_Entity :: struct {
	pos: Point,
	type: enum {
		Cultist,
		Nihil  ,
		Dead   ,
	}
}

update_game :: proc "c" () {

}

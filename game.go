package main

type Game struct {
	t1_picks [5]*Champion
	t1_bans  [3]*Champion
	t2_picks [5]*Champion
	t2_bans  [3]*Champion
}

func NewGame() *Game {
	return &Game{}
}

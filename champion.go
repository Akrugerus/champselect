package main

type Champion struct {
	version string
	key     string
	name    string
	id      string
	title   string
	blurb   string
	info    ChampionInfo
	image   ChampionImage
}

type ChampionInfo struct {
	attack     int
	defense    int
	magic      int
	difficulty int
}

type ChampionImage struct {
	full   string
	sprite string
	group  string
	x      int
	y      int
	w      int
	h      int
}

func DefaultChampion() Champion {
	return Champion{
		name: "Mordekaiser",
	}
}

func LoadChampions() []*Champion {
	return []*Champion{}
}

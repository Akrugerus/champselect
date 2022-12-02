package main

import "encoding/json"

type Champion struct {
	Version string `json:"version"`
	Key     string `json:"key"`
	Name    string `json:"name"`
	Id      string `json:"id"`
	Title   string `json:"title"`
	Blurb   string `json:"blurb"`
	Info    ChampionInfo
	// image   ChampionImage
}

type ChampionInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}

type ChampionImage struct {
	Full   string
	Sprite string
	Group  string
	X      int
	Y      int
	W      int
	H      int
}

type ChampionJSON struct {
	ChampType string              `json:"type"`
	Format    string              `json:"format"`
	Version   string              `json:"version"`
	Data      map[string]Champion `json:"data"`
}

// LoadChampions pulls general champion data out of the dragondata archive and returns
// a ChampionJSON struct
func LoadChampions() (*ChampionJSON, error) {

	// Load the dragondata tarball
	archive, err := GetDragonArchive()
	if err != nil {
		return nil, err
	}

	// Pull the champion.json file out of the archive
	champinfo, err := ExtractFile("12.22.1/data/en_US/champion.json", archive)
	if err != nil {
		return nil, err
	}

	// Parse the champion.json file into a metadata object and return
	var champs = ChampionJSON{}
	if err := json.Unmarshal(champinfo, &champs); err != nil {
		return nil, err
	}
	return &champs, nil
}

func ChampionList() []string {
	champjson, err := LoadChampions()
	if err != nil {
		panic(err)
	}

	list := make([]string, len(champjson.Data))
	n := 0
	for k := range champjson.Data {
		list[n] = k
		n += 1
	}
	return list
}

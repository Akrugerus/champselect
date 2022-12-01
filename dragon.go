package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const temp_dragon_url = "https://ddragon.leagueoflegends.com/cdn/dragontail-12.22.1.tgz"
const dragon_path = "."

func FetchDragon() (*os.File, error) {
	fmt.Println("Downloading dragon tarball")
	// Fetch the dragon tarball
	res, err := http.Get(temp_dragon_url)
	if err != nil {
		fmt.Println("Unable to fetch dragon")
		return nil, err
	}
	fmt.Printf("%v", res.Header)

	// Open the dragon tarball file handler
	file, err := os.Create("dragon.tar.gz")
	if err != nil {
		return nil, err
	}

	// Copy the contents of the response body into the file
	if _, err := io.Copy(file, res.Body); err != nil {
		return nil, err
	}

	// Reset the file handler
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}

	// Return the file handler
	return file, nil
}

func LoadDragon() error {
	// declare a basic reader
	f, err := os.Open("dragon.tar.gz")
	if err != nil {
		f, err = FetchDragon()
		if err != nil {
			return err
		}
	}

	// Load the champ info from the archive
	champinfo, err := UntarSingleFile("12.22.1/data/en_US/champion.json", f)
	if err != nil {
		return err
	}

	// decode the root champ info to json
	var a = make(map[string]any, 0)
	if err = json.Unmarshal(champinfo, &a); err != nil {
		return err
	}

	fmt.Printf("%v", a["data"])

	// Unzip the dragon tarball
	// fmt.Println("Unzipping tarball")
	// exec.Command("tar", "-xzvf", "dragon.tar.gz")

	return nil
}

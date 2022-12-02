package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const temp_dragon_url = "https://ddragon.leagueoflegends.com/cdn/dragontail-12.22.1.tgz"

func FetchDragon() (*os.File, error) {
	fmt.Println("Downloading dragon tarball")
	// Fetch the dragon tarball
	res, err := http.Get(temp_dragon_url)
	if err != nil {
		fmt.Println("Unable to fetch dragon")
		return nil, err
	}
	fmt.Printf("%v\n", res.Header)

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

func GetDragonArchive() (*os.File, error) {
	f, err := os.Open("dragon.tar.gz")
	if err != nil {
		f, err = FetchDragon()
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

// func LoadDragon() error {
// 	// declare a basic reader
// 	f, err := os.Open("dragon.tar.gz")
// 	if err != nil {
// 		f, err = FetchDragon()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Load the champ info from the archive

// 	// decode the root champ info to map
// 	var a = make(map[string][]byte, 0)
// 	if err = json.Unmarshal(champinfo, &a); err != nil {
// 		return err
// 	}

// 	// Print the map keys which should be champion names
// 	var champs = make([]map[string][]byte, 0)
// 	if err = json.Unmarshal(a["data"], &champs)
// 		// Unzip the dragon tarball
// 	// fmt.Println("Unzipping tarball")
// 	// exec.Command("tar", "-xzvf", "dragon.tar.gz")

// 	return nil
// }

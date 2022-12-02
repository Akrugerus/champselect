package main

import "fmt"

func main() {
	clist := ChampionList()
	for _, x := range clist {
		fmt.Printf("%v\n", x)
	}
	fmt.Printf("Number of champions: %v\n", len(clist))
}

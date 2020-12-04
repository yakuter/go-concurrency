package main

import (
	"fmt"
	"time"
)

func main() {
	heroes := []string{"Marvel", "Flash", "Thanos", "Eagle", "Hulk", "Thor"}
	spaceships := []string{"Battlecruiser", "Battleship", "Cruiseship", "Her Majesty's Ship", "Imperial Spaceship"}
	go listHeroes(heroes)
	go listSpaceships(spaceships)

	<-time.After(time.Second * 10)
}

func listHeroes(items []string) {
	for i := range items {
		fmt.Printf("Hero: %s\n", items[i])
		time.Sleep(time.Second)
	}

}

func listSpaceships(items []string) {
	for i := range items {
		fmt.Printf("Spaceship: %s\n", items[i])
		time.Sleep(time.Second)
	}
}

/* OUTPUT
// Output always changes

Hero: Marvel
Spaceship: Battlecruiser
Spaceship: Battleship
Hero: Flash
Spaceship: Cruiseship
Hero: Thanos
Spaceship: Her Majesty's Ship
Hero: Flash
Hero: Hulk
Spaceship: Imperial Spaceship
Hero: Thor

*/

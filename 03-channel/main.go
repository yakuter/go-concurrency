// https://ednsquare.com/story/learn-go-introduction-to-channels-in-golang------bocBYu

package main

import (
	"fmt"
	"time"
)

func main() {
	theMine := [3]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)
	// Finder
	go func(mine [3]string) {
		for _, item := range mine {
			oreChan <- item //send
		}
	}(theMine)

	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //receive
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
	}()

	<-time.After(time.Second * 5) // Again, ignore this for now
}

package main

import (
	"fmt"
	"time"
)

func numbers(start chan bool) {
	// Waiting for start chan to have a value or closed
	<-start
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets(start chan bool) {
	// Waiting for start chan to have a value or closed
	<-start
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func main() {

	// In this example, we used channel close to
	// start and stop goroutines

	start := make(chan bool)

	go numbers(start)
	go alphabets(start)

	time.Sleep(2 * time.Second)
	close(start)
	fmt.Println("Channel closed after 2 second")

	<-time.After(time.Second * 5)
}

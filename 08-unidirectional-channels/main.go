package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	// Send only channel
	sendch := make(chan<- int)

	go sendData(sendch)

	// Panic line
	fmt.Println(<-sendch)
}

// This program can not build because sendch is send only channel.
// You can not receive value from send only channel

// OUTPUT
// ./main.go:12:14: invalid operation: <-sendch (receive from send-only type chan<- int)

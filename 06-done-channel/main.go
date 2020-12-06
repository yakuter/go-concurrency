package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func main() {

	// We can use channels to wait for goroutines
	// This is called done channel which is in fact normal channel

	done := make(chan bool)
	go hello(done)

	<-done
	fmt.Println("main function")
}

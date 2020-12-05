// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "fmt"

func main() {

	value := "PassWall"

	firstChan := make(chan string, 1)

	sender(firstChan, value)

	found := receiver(firstChan)

	fmt.Println(found)

	// direction transports data from one channel to another
	// secondChan := make(chan string, 1)
	// direction(firstChan, secondChan)

}

func sender(channel chan<- string, message string) {
	channel <- message
}

func receiver(channel <-chan string) string {
	message := <-channel
	return message
}

func direction(receiver <-chan string, sender chan<- string) {
	message := <-receiver
	sender <- message
}

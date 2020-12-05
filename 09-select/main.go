// https://golangbyexample.com/select-statement-golang/
package main

import (
	"fmt"
	"time"
)

func main() {

	// Once any receive operation is complete
	// on any of the channels it is executed
	// and select exits. So as seen from the output,
	// in the above program, it prints the received
	// value from one of the channels and exits.

	ch1 := make(chan string)
	ch2 := make(chan string)

	go goTwo(ch2)
	go goOne(ch1)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	<-time.After(1 * time.Second)
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	// You can try with sleep to be sure
	// time.Sleep(time.Second * 1)
	ch <- "From goTwo goroutine"
}

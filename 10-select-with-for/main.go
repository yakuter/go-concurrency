// https://golangbyexample.com/select-statement-golang/
package main

import "fmt"

func main() {

	//It is also possible to wait for receive
	// operation to complete on both the channels
	// by using a for loop across the select
	// statement. Let's see a program for that

	ch1 := make(chan string)
	ch2 := make(chan string)
	go goOne(ch1)
	go goTwo(ch2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	ch <- "From goTwo goroutine"
}

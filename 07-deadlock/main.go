package main

func main() {
	// If a Goroutine (including main process) is sending data on a channel,
	// then it is expected that some other Goroutine
	// should be receiving the data. If this does not happen,
	// then the program will panic at runtime with Deadlock.

	// If a Goroutine is waiting to receive data from a channel,
	// then some other Goroutine is expected to write data on that channel,
	// else the program will panic.

	ch := make(chan int)
	ch <- 5
}

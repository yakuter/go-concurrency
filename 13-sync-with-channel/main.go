package main

import (
	"fmt"
	"sync"
)

func main() {

	var v string
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan bool)
	go func() {
		v = "Channel value"
		ch <- true
		wg.Done()
	}()
	go func() {
		<-ch
		fmt.Println(v)
		wg.Done()
	}()
	wg.Wait()
}

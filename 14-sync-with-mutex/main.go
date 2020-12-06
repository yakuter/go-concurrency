package main

import (
	"fmt"
	"sync"
)

func main() {

	var v string
	var wg sync.WaitGroup
	wg.Add(2)

	var m sync.Mutex
	m.Lock()

	go func() {
		v = "Channel value"
		m.Unlock()
		wg.Done()
	}()

	go func() {
		m.Lock()
		fmt.Println(v)
		wg.Done()
	}()

	wg.Wait()
}

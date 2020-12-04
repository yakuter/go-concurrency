package main

import (
	"fmt"
	"time"
)

func main() {

	// Anonymous function
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("%d ", i)
		}
	}()

	// Anonymous function
	go func() {
		for i := 'a'; i <= 'k'; i++ {
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%c ", i)
		}
	}()

	<-time.After(time.Second * 5)
}

/* OUTPUT
// Output always changes
1 a 2 3 b 4 c 5 6 d 7 8 e 9 f 10 g h i j k %
*/

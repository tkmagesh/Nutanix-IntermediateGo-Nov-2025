package main

import (
	"fmt"
	"time"
)

// communicate by sharing memory
var stop bool = false

func main() {
	ch := genNos()

	fmt.Println("Hit ENTER to stop...!")
	go func() {
		fmt.Scanln()
		fmt.Println("Stopping initiated")
		stop = true
	}()

	for no := range ch {
		fmt.Println("No :", no)
	}

	fmt.Println("Done!")
}

func genNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			if stop {
				close(ch)
				break
			}
			ch <- (i + 1) * 10
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return ch
}

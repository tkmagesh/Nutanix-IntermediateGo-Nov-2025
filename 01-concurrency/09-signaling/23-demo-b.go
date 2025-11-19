package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan bool)
	ch := genNos(stopCh)

	fmt.Println("Hit ENTER to stop...!")
	go func() {
		fmt.Scanln()
		fmt.Println("Stopping initiated")
		stopCh <- true
	}()

	for no := range ch {
		fmt.Println("No :", no)
	}

	fmt.Println("Done!")
}

func genNos(stopCh <-chan bool) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				close(ch)
				break LOOP
			case ch <- (i + 1) * 10:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	return ch
}

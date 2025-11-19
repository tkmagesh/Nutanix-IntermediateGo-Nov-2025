package main

import (
	"fmt"
	"time"
)

func main() {

	stopCh := time.After(5 * time.Second)
	ch := genNos(stopCh)

	for no := range ch {
		fmt.Println("No :", no)
	}

	fmt.Println("Done!")
}

// The following can be replaced with time.After()
/*
	func timeout(d time.Duration) <-chan time.Time {
		stopCh := make(chan time.Time)
		go func() {
			time.Sleep(d)
			fmt.Println("Stopping initiated")
			stopCh <- time.Now()
		}()
		return stopCh
	}
*/
func genNos(stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			case ch <- (i + 1) * 10:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

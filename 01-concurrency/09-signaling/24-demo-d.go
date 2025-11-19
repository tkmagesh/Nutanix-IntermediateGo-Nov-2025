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

func genNos(stopCh <-chan time.Time) <-chan int {
	ch := make(chan int)
	ticker := time.Tick(500 * time.Millisecond)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-stopCh:
				break LOOP
			case <-ticker:
				ch <- (i + 1) * 10
			}
		}
		close(ch)
	}()
	return ch
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		break
	}
	fmt.Println("Done!")
}

func genNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	fmt.Printf("[genNos]  - producing %d values \n", count)
	go func() {
		for i := range count {
			ch <- (i + 1) * 10
		}
		fmt.Println("[genNos] - all data produced, closing the channel")
		close(ch)
	}()
	return ch
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(<-ch)
	}
}

func genNos() <-chan int {
	ch := make(chan int)
	count := rand.Intn(20)
	go func() {
		for i := range count {
			ch <- (i + 1) * 10
		}
	}()
	return ch
}

/*
Modify the "consumer" logic in such a way the program prints

ch2 : 200
ch1 : 100
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// producer
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	// producer
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	fmt.Println("ch1 : ", <-ch1)
	fmt.Println("ch2 : ", <-ch2)
}

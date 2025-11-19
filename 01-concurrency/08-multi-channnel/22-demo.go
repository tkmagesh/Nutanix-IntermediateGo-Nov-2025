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

	// consumer
	ch3 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		d3 := <-ch3
		fmt.Println("data from ch3 :", d3)
	}()

	// simulating the success of all the channel operations in 'select-case'
	// time.Sleep(6 * time.Second)

	// consumers
	for range 3 {
		select {
		case d1 := <-ch1:
			fmt.Println("ch1 :", d1)
		case d2 := <-ch2:
			fmt.Println("ch2 :", d2)
		case ch3 <- 300:
			fmt.Println("sent data to ch3")
			/*
				default:
					fmt.Println("no channel operations were successfull!")
			*/
		}
	}

	fmt.Println("Done!")

}

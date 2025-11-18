/*
Modify the program to follow "Share Memory by Communicating" (rather than "Communicate by sharing memory")
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	primesCh := make(chan int)
	start, end := 2, 100
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkIfPrime(no, wg, primesCh)
	}
	go func() {
		wg.Wait()
		close(primesCh)
	}()

	for primeNo := range primesCh {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func checkIfPrime(no int, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	if isPrime(no) {
		ch <- no
	}
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

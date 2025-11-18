package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	workerCount := 5
	primesCh := make(chan int)
	dataCh := getDataFeeder(2, 100)

	for range workerCount {
		wg.Add(1)
		go primeWorker(wg, dataCh, primesCh)
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

func getDataFeeder(start, end int) <-chan int {
	dataCh := make(chan int)
	go func() {
		for no := start; no <= end; no++ {
			dataCh <- no
		}
		close(dataCh)
	}()
	return dataCh
}

func primeWorker(wg *sync.WaitGroup, dataCh <-chan int, primesCh chan<- int) {
	defer wg.Done()
	for no := range dataCh {
		if isPrime(no) {
			primesCh <- no
		}
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

/*
Modify the program to follow "Share Memory by Communicating" (rather than "Communicate by sharing memory")
*/

package main

import (
	"fmt"
	"sync"
)

// communicate by sharing memory
var primes []int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	start, end := 2, 100
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkIfPrime(no, wg)
	}
	wg.Wait()
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func checkIfPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	if isPrime(no) {
		// DO NOT print if the number is a prime number
		// fmt.Printf("Prime No : %d\n", no)
		mutex.Lock()
		{
			primes = append(primes, no)
		}
		mutex.Unlock()
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

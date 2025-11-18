/*
Print the prime numbers found in the "main()" function
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	start, end := 2, 100
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkIfPrime(no, wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func checkIfPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	if isPrime(no) {
		// DO NOT print if the number is a prime number
		// fmt.Printf("Prime No : %d\n", no)
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

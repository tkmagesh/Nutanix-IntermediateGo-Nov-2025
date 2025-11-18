/* convert the solution to execute the prime check logic concurrently */
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
		go func() {
			defer wg.Done()
			printIfPrime(no)
		}()
	}
	wg.Wait()
	fmt.Println("Done")
}

func printIfPrime(no int) {
	if isPrime(no) {
		fmt.Printf("Prime No : %d\n", no)
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

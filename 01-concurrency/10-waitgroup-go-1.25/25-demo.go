package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for range 20 {
		wg.Go(f1)
	}
	wg.Wait() // block the execution of "this" function until the wg counter becomes 0 (default = 0)
	fmt.Println("Done!")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")

}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // increment the wg counter by 1
	go f1()   //scheduling the execution of "f1()" through the scheduler to be executed in "future"

	f2()

	wg.Wait() // block the execution of "this" function until the wg counter becomes 0 (default = 0)
}

func f1() {
	fmt.Println("f1 started")
	// simulate a time consuming operation
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}

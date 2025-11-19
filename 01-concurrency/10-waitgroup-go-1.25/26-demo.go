package main

import (
	"fmt"
	"sync"
	"time"
)

type MyWaitGroup struct {
	sync.WaitGroup
}

func (myWG *MyWaitGroup) MyGo(f func()) {
	myWG.Add(1)
	go func() {
		defer myWG.Done()
		f()
	}()
}

func main() {
	wg := &MyWaitGroup{}
	for range 20 {
		wg.MyGo(f1)
	}
	wg.Wait() // block the execution of "this" function until the wg counter becomes 0 (default = 0)
	fmt.Println("Done!")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")

}

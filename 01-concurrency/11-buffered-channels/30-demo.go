package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// handle the error
	/*
		errCh := printNos(wg)
		fmt.Println(<-errCh)
	*/

	// consumer has choosen NOT to handle the error
	printNos(wg)
	wg.Wait()
}

func printNos(wg *sync.WaitGroup) <-chan error {
	// errCh := make(chan error) // result in a deadlock if the consumer has chosen NOT to receive the error
	errCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		var no int
		for {
			no = rand.Intn(100)
			if no%7 == 0 {
				fmt.Println("[error] no :", no)
				errCh <- errors.New("encountered multiple of 7")
				break
			}
			fmt.Println("no :", no)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return errCh
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	doneCh := printNumbers(cancelCtx)
	<-doneCh
	fmt.Println("Done!")
}

func printNumbers(ctx context.Context) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		evenTimeoutCtx, _ := context.WithTimeout(ctx, 5*time.Second)
		go printEven(evenTimeoutCtx, wg)

		wg.Add(1)
		oddTimeoutCtx, _ := context.WithTimeout(ctx, 10*time.Second)
		go printOdd(oddTimeoutCtx, wg)

		wg.Add(1)
		go printRandom(ctx, wg)

		wg.Wait()
		close(doneCh)
	}()
	return doneCh
}

func printRandom(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()
	randomTicker := time.Tick(700 * time.Millisecond)
LOOP:
	for {
		select {
		case <-randomTicker:
			fmt.Println("[printRandom] no :", rand.Intn(100))
		case <-ctx.Done():
			fmt.Println("[printRandom] cancellation signal received")
			break LOOP
		}
	}

}

func printEven(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	evenTicker := time.Tick(300 * time.Millisecond)
LOOP:
	for i := 0; ; i += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[printEven] cancellation signal received")
			break LOOP
		case <-evenTicker:
			fmt.Println("[printEven] no :", i)
		}

	}
}

func printOdd(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	oddTicker := time.Tick(500 * time.Millisecond)
LOOP:
	for i := 1; ; i += 2 {
		select {
		case <-ctx.Done():
			fmt.Println("[printOdd] cancellation signal received")
			break LOOP
		case <-oddTicker:
			fmt.Println("[printOdd] no :", i)
		}

	}
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	/*
		rootCtx := context.Background()
		valCtx := context.WithValue(rootCtx, "root-key", "root-value")
		cancelCtx, cancel := context.WithCancel(valCtx)
	*/

	ctx := context.Background()
	ctx = context.WithValue(ctx, "root-key", "root-value")
	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("Hit ENTER to stop...")
	go func() {
		fmt.Scanln()
		cancel()
	}()
	doneCh := printNumbers(ctx)
	<-doneCh
	fmt.Println("Done!")
}

func printNumbers(ctx context.Context) <-chan struct{} {
	fmt.Printf("[printNumbers] root-key : %v\n", ctx.Value("root-key"))
	doneCh := make(chan struct{})
	go func() {
		wg := &sync.WaitGroup{}

		// overriding the root-key
		newRootKey := context.WithValue(ctx, "root-key", "new-root-value")

		numberCtx := context.WithValue(newRootKey, "branch-key", "branch-value")

		wg.Add(1)
		evenTimeoutCtx, _ := context.WithTimeout(numberCtx, 5*time.Second)
		go printEven(evenTimeoutCtx, wg)

		wg.Add(1)
		oddTimeoutCtx, _ := context.WithTimeout(numberCtx, 10*time.Second)
		go printOdd(oddTimeoutCtx, wg)

		wg.Add(1)
		go printRandom(numberCtx, wg)

		wg.Wait()
		close(doneCh)
	}()
	return doneCh
}

func printRandom(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("[printRandom] root-key : %v\n", ctx.Value("root-key"))
	fmt.Printf("[printRandom] branch-key : %v\n", ctx.Value("branch-key"))
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
	fmt.Printf("[printEven] root-key : %v\n", ctx.Value("root-key"))
	fmt.Printf("[printEven] branch-key : %v\n", ctx.Value("branch-key"))
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
	fmt.Printf("[printOdd] root-key : %v\n", ctx.Value("root-key"))
	fmt.Printf("[printOdd] branch-key : %v\n", ctx.Value("branch-key"))
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

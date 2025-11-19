/* programmatic cancellation */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)

	ch := genNos(cancelCtx)

	fmt.Println("Hit ENTER to stop...!")
	go func() {
		fmt.Scanln()
		fmt.Println("Stopping initiated")
		cancel()
	}()

	for no := range ch {
		fmt.Println("No :", no)
	}

	fmt.Println("Done!")
}

func genNos(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("[genNos] cancellation signal received")
				break LOOP
			case ch <- (i + 1) * 10:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

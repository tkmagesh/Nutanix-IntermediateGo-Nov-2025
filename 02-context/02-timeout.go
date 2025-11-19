/* time based (relative) cancellation */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(rootCtx, 5*time.Second)

	ch := genNos(timeoutCtx)

	fmt.Println("Will stop after 5 secs.. Hit ENTER to stop before 5 secs...!")
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
				switch ctx.Err() {
				case context.Canceled:
					fmt.Println("[genNos] programmatic cancellation signal received")
				case context.DeadlineExceeded:
					fmt.Println("[genNos] time based cancellation signal received")
				}
				break LOOP
			case ch <- (i + 1) * 10:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

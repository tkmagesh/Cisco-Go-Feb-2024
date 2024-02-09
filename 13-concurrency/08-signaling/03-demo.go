/* generate fibonocci series until the user hits ENTER key */

package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genFib(stopCh)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-stopCh:
				fmt.Println("stop signal received")
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- x
			}
		}
		close(ch)
	}()
	return ch
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */

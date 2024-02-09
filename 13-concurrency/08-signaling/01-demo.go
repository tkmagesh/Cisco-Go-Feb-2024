/* generate fibonocci series for specific time */

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	ch := make(chan int)
	timeoutCh := timeout(10 * time.Second)
	go func() {
	LOOP:
		for x, y := 0, 1; ; x, y = y, x+y {
			select {
			case <-timeoutCh:
				fmt.Println("timeout signal received")
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

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}

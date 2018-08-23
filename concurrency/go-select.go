package concurrency

import (
	"fmt"
	"time"
)

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/*
GoSelect - The select statement lets a goroutine wait on multiple communication operations.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/
func GoSelect() {
	fmt.Print("****Running concurrency.GoSelect(), Go select statement \n")
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

/*
DefaultSelect - The default case in a select is run if no other case is ready.

Use a default case to try a send or receive without blocking:
*/
func DefaultSelect() {
	fmt.Print("****Running concurrency.DefaultSelect(), Go default select statement \n")
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

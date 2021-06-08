package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib1(c, quit)
	/*
	0
	1
	1
	2
	3
	5
	8
	13
	21
	34
	quit
	*/

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <- tick:
			fmt.Println("tick...")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("...sleep every 50 ms")
			time.Sleep(50 * time.Millisecond)
		}
	}
	/*
	...sleep every 50 ms
	...sleep every 50 ms
	tick...
	...sleep every 50 ms
	...sleep every 50 ms
	tick...
	...sleep every 50 ms
	...sleep every 50 ms
	tick...
	...sleep every 50 ms
	tick...
	...sleep every 50 ms
	...sleep every 50 ms
	tick...
	BOOM!
	*/
}

func fib1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

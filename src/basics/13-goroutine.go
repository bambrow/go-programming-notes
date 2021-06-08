package main

import (
	"fmt"
	"time"
)

func main() {
	go say("a")
	say("b")
	/*
	a 0
	b 0
	b 1
	a 1
	a 2
	b 2
	b 3
	a 3
	a 4
	b 4
	*/
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

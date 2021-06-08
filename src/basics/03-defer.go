package main

import "fmt"

func main() {
	fmt.Println("Counting...")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done!")

	/*
	Counting...
	Done!
	9
	8
	7
	6
	5
	4
	3
	2
	1
	0
	*/
}

package main

import (
	"fmt"
	"math"
)

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot)) // 2.23606797749979
	fmt.Println(compute(math.Pow)) // 1

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
	/*
	0 0
	1 -2
	3 -6
	6 -12
	10 -20
	15 -30
	21 -42
	28 -56
	36 -72
	45 -90
	*/

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	/*
	1
	1
	2
	3
	5
	8
	13
	21
	34
	55
	*/
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(1, 2)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

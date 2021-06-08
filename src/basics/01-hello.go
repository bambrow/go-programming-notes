package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var k1, k2 bool = true, false
const j1 = "constant"

func main() {
	fmt.Println("hello world")
	fmt.Println("The time is", time.Now())
	fmt.Println("A random number", rand.Intn(10))
	fmt.Println("3 + 5 =", add(3, 5))
	x, y := swap("x", "y")
	fmt.Println("Swap x and y:", x, y) // Swap x and y: y x
	fmt.Println(foo(25)) // 5 20
	var k3, k4 = 1, "hello"
	fmt.Println(k1, k2, k3, k4) // true false 1 hello

	var k5 float64 = 1.2345
	var k6 complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println(k5, k6)  // 1.2345 (2+3i)
	fmt.Printf("Type: %T, value: %v\n", k5, k5) // Type: float64, value: 1.2345
	fmt.Printf("Type: %T, value: %v\n", k6, k6) // Type: complex128, value: (2+3i)
	fmt.Printf("Type: %T, content: with quote %q without quote %v\n", k4, k4, k4)
	// Type: string, content: with quote "hello" without quote hello

	k7, k8 := 3, 4
	fmt.Println("Sqrt(", k7, "*", k7, "+", k8, "*", k8, ") =", math.Sqrt(float64(k7 * k7 + k8 * k8)))
	// Sqrt( 3 * 3 + 4 * 4 ) = 5

	fmt.Println("j1 is", j1)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func foo(x int) (y, z int) {
	y = x % 10
	z = x - y
	return
}
package main

import "fmt"

func main() {

	i, j := 42, 81

	p := &i // p points to i
	fmt.Println(*p) // use pointer to read i, 42
	*p = 21 // use pointer to set i
	fmt.Println(i) // 21

	p = &j // p points to j
	*p = *p / 9 // use pointer to calculate
	fmt.Println(j) // 9

}

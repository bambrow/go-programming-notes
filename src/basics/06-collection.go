package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a) // [hello world]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// [2 3 5 7 11 13]

	fmt.Println(primes[1:4])
	// [3 5 7]

	names := [4]string{"Amy", "Bob", "Cathy", "David"}
	b := names[1:]
	b[0] = "Bran"
	fmt.Println(names)
	// [Amy Bran Cathy David]

	s := []struct {
		i int
		b bool
	} {
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(s)
	// [{1 true} {2 false} {3 true}]

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(s1) // len = 8, cap = 8, [1 2 3 4 5 6 7 8]
	s1 = s1 [:0]
	printSlice(s1) // len = 0, cap = 8, []
	s1 = s1[:4]
	printSlice(s1) // len = 4, cap = 8, [1 2 3 4]
	s1 = s1[2:]
	printSlice(s1) // len = 2, cap = 6, [3 4]

	var s2 []int
	printSlice(s2) // len = 0, cap = 0, []
	fmt.Println(s2 == nil) // true

	s3 := make([]int, 5)
	printSlice(s3) // len = 5, cap = 5, [0 0 0 0 0]
	s4 := make([]int, 0, 5)
	printSlice(s4) // len = 0, cap = 5, []
	s5 := s4[:2]
	printSlice(s5) // len = 2, cap = 5, [0 0]
	s6 := s5[2:5]
	printSlice(s6) // len = 3, cap = 3, [0 0 0]

	var r1 []int
	r1 = append(r1, 0)
	printSlice(r1) // len = 1, cap = 1, [0]
	r1 = append(r1, 1, 2, 3, 4)
	printSlice(r1) // len = 5, cap = 6, [0 1 2 3 4]

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2 ^ %d = %d\n", i, v)
	}
	/*
	2 ^ 0 = 1
	2 ^ 1 = 2
	2 ^ 2 = 4
	2 ^ 3 = 8
	2 ^ 4 = 16
	2 ^ 5 = 32
	2 ^ 6 = 64
	2 ^ 7 = 128
	*/

	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // 2 ^ i
	}
	for _, v := range pow {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// 1 2 4 8 16 32 64 128 256 512

}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, %v\n", len(s), cap(s), s)
	// len is the length of array or slice
	// cap is the slice length till the end of underlying array
}

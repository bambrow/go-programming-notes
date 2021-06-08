package main

import (
	"fmt"
	"strings"
)

type Vertex1 struct {
	X int
	Y int
}

func main() {
	var m map[string]Vertex1

	m = make(map[string]Vertex1)
	m["a"] = Vertex1{ 1, 2 }
	fmt.Println(m["a"]) // {1 2}

	var m1 = map[string]Vertex1 {
		"a": {
			1, 2,
		},
	}
	fmt.Println(m1["a"]) // {1 2}

	fmt.Println(m1["b"]) // {0 0}
	v, ok := m1["a"]
	fmt.Println(v, ok) // {1 2} true

	fmt.Println(WordCount("a b c d a b c a b a"))
	// map[a:4 b:3 c:2 d:1]
}

func WordCount(s string) map[string]int {
	var a = strings.Fields(s)
	var m = make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	return m
}

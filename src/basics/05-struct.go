package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	fmt.Println(Vertex{1, 2}) // {1 2}
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v) // {4 2}

	p := &v
	p.X = 9
	fmt.Println(v) // {9 2}

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p1 = &Vertex{2, 1}
	)
	fmt.Println(v1, v2, v3, p1)
	// {1 2} {1 0} {0 0} &{2 1}

}

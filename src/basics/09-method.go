package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex2) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func ScaleFunc(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs()) // 5
	v.Scale(10)
	fmt.Println(v.Abs()) // 50

	p := &Vertex2{3, 4}
	ScaleFunc(p, 10)
	fmt.Println(p, p.Abs()) // &{30 40} 50

	var a Abser
	a = v
	fmt.Println(a.Abs()) // 50
}

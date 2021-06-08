package main

import (
	"fmt"
	"math"
	"strconv"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func desc(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v)", p.Name, p.Age)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M() // Hello

	i = F(math.Pi)
	describe(i) // (3.141592653589793, main.F)
	i.M() // 3.141592653589793

	var t *T
	i = t
	describe(i) // (<nil>, *main.T)
	i.M() // <nil>

	var i1 interface{}
	desc(i1) // (<nil>, <nil>)

	var i2 interface{} = "a"
	s, ok := i2.(string)
	fmt.Println(s, ok) // a true
	f, ok := i2.(float64)
	fmt.Println(f, ok) // 0 false

	do(21) // int type 21
	do(true) // other type bool

	k, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("cannot convert number %v\n", err)
	} else {
		fmt.Printf("converted number %v\n", k)
	}
	// converted number 42

	p := Person{"David", 30}
	fmt.Println(p) // David (30)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int type %v\n", v)
	case string:
		fmt.Printf("string type %q\n", v)
	default:
		fmt.Printf("other type %T\n", v)
	}
}

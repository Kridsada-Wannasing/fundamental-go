package main

import (
	"fmt"
)

func main() {
	var a interface{}

	a = 10
	fmt.Printf("%T %v\n", a, a)

	if i, ok := a.(float32); ok {
		fmt.Println(i)
	}

	var x, y Obj
	x = rectangle{w: 4, l: 8}
	y = triangle{w: 4, h: 8}

	fmt.Println(x.Area())
	fmt.Println(y.Area())

	if v, ok := y.(triangle); ok {
		v.h = 5
	}
}

type Obj interface {
	Area() float64
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return tri.h * tri.w * 0.5
}

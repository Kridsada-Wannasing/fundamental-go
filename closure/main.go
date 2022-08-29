package main

import "fmt"

func main() {
	fn := newCounterFuc()

	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func newCounterFuc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

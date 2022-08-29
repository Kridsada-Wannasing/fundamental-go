package main

import (
	"fmt"
	"fundamental-go/pack/book"
)

func main() {
	b := book.New()

	fmt.Printf("%T\n", b)
}

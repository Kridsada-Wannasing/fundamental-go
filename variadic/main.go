package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	// s... ตือ spread operator
	variadicString(s...)
}

// ...s ใน func จะเห็นเป็น slice
func variadicString(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

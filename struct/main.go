package main

import "fmt"

type rectangle struct {
	w, l float64
}

// function style
func Area(rec rectangle) float64 {
	return rec.l * rec.w
}

// method style
func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

// หาก reciever เป็น type ธรรมดา (rec rectangle) จะเป็นการ copy rec แทน และจะถูกทำลายเมื่อหมด scope
func (rec *rectangle) SetWidth(w float64) {
	rec.w = w
}

func main() {
	rec := rectangle{w: 4, l: 5}

	rec.w = 6

	fmt.Println(rec.l)
	fmt.Println(rec.w)
	fmt.Println(Area(rec))
	fmt.Println(rec.Area())

	rec.SetWidth(6)
}

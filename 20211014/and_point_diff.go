package main

import "fmt"

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) size() float64 {
	return r.width * r.width
}

func main() {
	fmt.Println(&Rect{10, 2}) // &{10 2}

	var r = &Rect{10, 3}
	fmt.Println(r)  // &{10 2}
	fmt.Println(*r) //{10 3}
	fmt.Println(&r) // 0xc00000e030
}

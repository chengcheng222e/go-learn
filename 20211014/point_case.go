package main

import "fmt"

func main() {
	a := "vernonchen"
	b := &a

	fmt.Println("a的值", a)
	fmt.Println("a的地址", b)
	fmt.Println("b的值", b)
	fmt.Println("b的指针指向", *b)

	*b = "not vernonchen"
	fmt.Println("修改后，b的指针指向", *b)
}

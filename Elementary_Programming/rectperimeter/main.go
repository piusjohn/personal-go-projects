package main

import(
	"fmt"
)

func RectPerimeter(w, h int) int {
	if h < 0 || w < 0{
		return -1
	}
	result := (w+h)*2
	//final := result * h
	return result
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
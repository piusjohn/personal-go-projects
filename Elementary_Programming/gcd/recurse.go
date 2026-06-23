package main

import "fmt"

func countdown(n int) int{
	if n < 0 {
		return 0
	}
	fmt.Println(n)
	return countdown(n-1)
}

// func main() {
// 	countdown(8)
// }

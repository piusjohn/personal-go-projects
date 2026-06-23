package main

import (
	"fmt"
)
func DigitLen(n, base int) int {
	count := 0
	if base < 2 || base > 32{
		return -1
	}
	if n < 0{
		n = -n
	}
	for n > 0{
		n = n/base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
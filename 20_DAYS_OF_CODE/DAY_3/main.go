package main

import (
	"fmt"
)

// func RepeatAlpha(s string) string {}
func main() {
	//var i int
	res := ""
	index := []int{}
	alphabet := "boy"
	//alphabet := strings.Split(s2, "")
	//alphabet = []string{"b", "o", "y"}
	s := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for i := 0; i < len(alphabet); i++ {
		for j := 0; j < len(s); j++ {
			if string(alphabet[i]) == s[j] {
				index = append(index, j+1)
				fmt.Println(index)
			}
		}
		for _, ir := range index {
			for k := 0; k < ir; k++ {
				res += string(alphabet[i])
			}
		}
	}

	// for _, ir := range index {
	// 	for k := 0; k < ir; k++{
	// 		res += alphabet[i]
	// 	}
	// }
	fmt.Println(res)
}

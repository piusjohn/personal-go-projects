package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello, boy, born how are you"
	ghost := "b"
	fmt.Println(Ghost(s, ghost))
}

func Ghost(s, ghost string) string {

	//var new string

	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		var new string
		if string(word[i][0]) == ghost {
			runes := []rune(word[i])
			for i := len(runes) - 1; i >= 0; i-- {

				new += string(runes[i])
			}
			word[i] = new
		}

	}
	return strings.Join(word, " ")
}

package main

import(
	"fmt"
)

func FirstWord(s string) string {
	var runes []rune
    for _, r := range s{
		if r == ' '{
			break
		}
		runes = append(runes, r)
	}
	return string(runes) + "\n"
}

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}
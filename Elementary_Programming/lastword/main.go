package main

import (
	"fmt"
)

func LastWord(s string) string{
	var  finalresult []rune
	var currentRune rune
	for i := len(s)-1; i >= 0; i-- {
		currentRune = rune(s[i])
		if currentRune == ' '{
			break
		}
		finalresult = append(finalresult, currentRune)
	}
	return string(finalresult)
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
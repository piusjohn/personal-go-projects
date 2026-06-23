package main

import (
	"fmt"
)

// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.

// If the resulting character is unprintable add 33 to it.

func HashCode(dec string) string {
	var currentRune int
	var result []rune
	for _, char := range dec{
		currentRune = int(char) + len(dec)%127
		if currentRune < 32 || currentRune > 127{
			currentRune = currentRune + 33
		}
		result = append(result, rune(currentRune))
	}
	return string(result)
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
package main

import (
	//"fmt"
	"strings"
)
func GenerateArt(input string, banner map[rune][]string) string {
	r, err := ValidateInput(input)
	if r != 0 && err != nil {
		return ("unsupported character")
	}
	var b strings.Builder
	data := SplitInput(input)
	for i, char := range data{
		if char != "" {
			b.WriteString(strings.Join(RenderLine(char, banner), "\n") + "\n")
		} else if i < len(data)-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
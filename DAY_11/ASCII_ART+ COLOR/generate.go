package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	dat, err := ValidateInput(input)
	if dat != 0 && err != nil{
		return "unsupported character"
	}
	data := SplitInput(input)
	
	var b strings.Builder
	for i, char := range data{
		if char != "" {
			b.WriteString(strings.Join(RenderLine(char, banner), "\n") + "\n")
		} else if i < len(data)-1 {
			b.WriteString("\n")
		}
	}
	finalres := b.String()
	return finalres
}



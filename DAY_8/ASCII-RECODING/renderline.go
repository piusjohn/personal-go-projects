package main

import "strings"

func RenderLine(input string, banner map[rune][]string) []string {
	
	var result []string
	for i := 0; i < 8; i++ {
		var b strings.Builder
		for _,  char := range input{
			b.WriteString(banner[char][i])
		}
		result = append(result, b.String())
	}
	return result
}
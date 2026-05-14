package main

import "strings"

func RenderLine(input string, banner map[rune][]string) []string {
	var result []string
	for i := 0; i < 8; i++ {
		var build strings.Builder
		for _, char := range input {
			build.WriteString(banner[char][i])
		}
		result = append(result, build.String())
	}
	return result

}

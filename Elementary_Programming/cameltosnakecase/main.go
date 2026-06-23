package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string{

	var result []rune
	if len(s) == 0{
		return ""
	}
	lastIndex := len(s)-1
	if s[lastIndex] >= 'A' && s[lastIndex] <= 'Z'{
		return s
	}
	for _, r := range s{
		if r >= '0' && r <= '9'{
			return s
		}
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++{
		if runes[i] >= 'A' && runes[i] <= 'Z' && runes[i+1] >= 'A' && runes[i+1] <= 'Z'{
			return s
		}
	}
	for i, char := range s{
		if i !=0{
			if char >= 'A' && char <= 'Z' {
				result = append(result, '_')
			}
		}
		result = append(result, char)
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
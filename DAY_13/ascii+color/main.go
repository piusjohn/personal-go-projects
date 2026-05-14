package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
}

func Colorize(colorstring, substring, fullstring string) string {
	colorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"reset":   "\033[0m",
	}
	if len(os.Args) == 3 || len(os.Args) == 4 {
		colorstring = os.Args[1]
		substring = os.Args[2]

		color := strings.TrimPrefix(colorstring, "--color=")
		//fmt.Println(color)

		cucolor := colorMap[color]
		r := colorMap["reset"]
		if len(os.Args) == 3 {
			fullstring := os.Args[2]
			out := fmt.Sprintf("%s%s%s\n", cucolor, fullstring, r)
			return out
		} else if len(os.Args) == 4 {
			fullstring := os.Args[3]
			colored := fmt.Sprintf("%s%s%s", cucolor, substring, r)
			new := strings.Replace(fullstring, substring, colored, -1)
			return new
		}

	}
	return ("Usage: go run . [OPTION] [STRING]\n") + ("EX: go run . --color=<color> <substring to be colored> something\n")

}

func LoadBanner(input string) (map[rune][]string, error){
	file, _ := os.Open(input)
	file.Close()
	content, _  := io.ReadAll(file)
	splitted := strings.Split(string(content), "\n")
	asciiMap := make(map[rune][]string)
	currentRune := ' '
	for i := 1; i  < len(splitted); i += 9 {
		if currentRune > '~'{
			break
		}
		asciiMap[currentRune] = splitted[i : i + 8]
		currentRune++
	}
	return asciiMap, nil
}

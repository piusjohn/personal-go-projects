package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	colorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"reset":   "\033[0m",
	}

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("USAGE: go run . helloworld\\n standard.txt")
		return
	}
	if !strings.HasPrefix(os.Args[1], "--color="){
		bannername := "standard"
	if len(os.Args) > 2 {
		bannername = os.Args[2]
	}
	
	userinput := os.Args[1]
	if userinput == ""{
		return 
	}

	bannerpath := fmt.Sprintf("%s.txt", bannername)
	banner, err := LoadBanner(bannerpath)
	if err != nil {
		fmt.Printf("error: Banner %s not found or invalid\n", bannername)
		return
	  
	}
	result := GenerateArt(userinput,  banner, "", "")
	  fmt.Println(result)

	}
	

		if strings.HasPrefix(os.Args[1], "--color="){
			fmt.Println("Type preffered Banner Style: 1. standard 2. shadow 3. thinkertoy")
			var bannerstyle string
			fmt.Scan(&bannerstyle)
			banner, _ := LoadBanner(bannerstyle + ".txt")
			color := strings.TrimPrefix(os.Args[1], "--color=")
			currentcolor, exists := colorMap[color]
			if !exists {
				fmt.Println("invalid color")
			}
			fmt.Println(GenerateArt(os.Args[3], banner, currentcolor, os.Args[2]))
	}

}


// func Colorize(colorstring, substring, fullstring string) string {
// 	colorMap := map[string]string{
// 		"red":     "\033[31m",
// 		"green":   "\033[32m",
// 		"yellow":  "\033[33m",
// 		"blue":    "\033[34m",
// 		"magenta": "\033[35m",
// 		"cyan":    "\033[36m",
// 		"reset":   "\033[0m",
// 	}
// 		color := strings.TrimPrefix(colorstring, "--color=")
// 		cucolor := colorMap[color]
// 		r := colorMap["reset"]
// 		if len(substring) == 0 {
// 			out := fmt.Sprintf("%s%s%s", cucolor, fullstring, r)
// 			return out
// 		} else if len(substring) != 0 {
// 			colored := fmt.Sprintf("%s%s%s", cucolor, substring, r)
// 			output := strings.Replace(fullstring, substring, colored, -1)
// 			return output
// 		}
// 		return ("Usage: go run . [OPTION] [STRING]\n") + "EX: go run . --color=<color> <substring to be colored> something"
			
// 	}


func LoadBanner(input string) (map[rune][]string, error){
	file, _ := os.Open(input)
	defer file.Close()
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

func RenderLine(input string, banner map[rune][]string) []string {
	var result []string
	for i := 0;  i < 8; i++ {
		var build string
		for _, char := range input{
		  build += banner[char][i]
		}
		 result = append(result, build)
	}
	
	return result
}

func Parse(input string)[]string {
	return strings.Split(input, "\\n")
}

func GenerateArt(input string,  banner map[rune][]string, color, substring string) string {
	reset := "\033[0m"
	start := strings.Index(input, substring)
	end := start + len(substring)
	word := Parse(input)
	var result string
	//var finalresult []string
	for i := 0; i < 8; i++ {
	for _, char := range word{
		for j, r := range char{
				if j >= start && j < end{
			result += color + banner[r][i] + reset
		} else{
			result += banner[r][i]
		}
		}
		result += "\n"
	}
}
	return result
}


package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	colorMap := map[string]string{
		"red": "\033[31m",
		"green": "\033[32m",
		"yellow": "\033[33m",
		"blue": "\033[34m",
		"magenta": "\033[35m",
		"cyan": "\033[36m",
		"reset": "\033[0m",
	}
	if len(os.Args) == 3 || len(os.Args) == 4 {
	colorstring := os.Args[1]
	substring := os.Args[2]
	
	color := strings.TrimPrefix(colorstring, "--color=")
	fmt.Println(color)

	cucolor := colorMap[color]
	r := colorMap["reset"]
		if len(os.Args) == 3 {
			fullstring := os.Args[2]
			fmt.Printf("%s%s%s\n", cucolor, fullstring, r)
		} else if len(os.Args) == 4{
			fullstring := os.Args[3]
			colored := fmt.Sprintf("%s%s%s", cucolor, substring, r)
			new := strings.Replace(fullstring, substring, colored, -1)
			fmt.Println(new)
		}

	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println("EX: go run . --color=<color> <substring to be colored> something")
		return
	}	

//	fmt.Printf("%s%s%s\n", cucolor, fullstring,r )

}
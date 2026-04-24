package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var res string
	if len(os.Args) < 2 {
		log.Fatal("please input a text")
		os.Exit(1)
	}

	bannername := "standard"
	if len(os.Args) > 2 && os.Args[2] == "standard" || os.Args[2] == "shadow" || os.Args[2] == "thinkeytoy" {
		bannername = os.Args[2]
	}
	bannerpath := fmt.Sprintf("banners/%s.txt", bannername)
	data, _ := os.ReadFile(bannerpath)
	startlines := strings.Split(string(data), "\n")
	userinput := strings.Split(os.Args[1], "\\n")

	for _, char := range userinput {
		for i := 0; i < 8; i++ {
			for _, word := range char {
				res += startlines[i+(int(word-' ')*9)+1]
			}
			res += "\n"
		}
	}
	fmt.Println(res)
}

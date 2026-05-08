package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("USAGE: go run . helloworld\\n standard.txt")
		return
	}
	bannername := "thinkertoy"
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
	
	fmt.Println(GenerateArt(userinput, banner))
}

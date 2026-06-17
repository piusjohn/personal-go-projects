package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: go run . 'userinput' standard.txt")
		return
	} 
	userinput := os.Args[1]
	if userinput == ""{
		return
	}

	bannername := "standard"
	if len(os.Args) > 2 {
		bannername = os.Args[2]
	}
	bannerpath := fmt.Sprintf("%s.txt", bannername)	
	
	
	_, err := LoadBanner(bannerpath)
	if err != nil {
		fmt.Printf("Error: Banner '%s' not found or invalid.\n", bannername)
		return
	}

	data, _ := LoadBanner(bannerpath)
	
	result := GenerateArt(userinput, data)
	fmt.Println(result)
	
}

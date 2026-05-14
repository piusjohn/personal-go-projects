package main

import (
	"fmt"
	"os"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

func main() {
	fmt.Println("Choose a color for your 'Hello, World!' message:")
	fmt.Println("1. Red")
	fmt.Println("2. Green")
	fmt.Println("3. Yellow")
	fmt.Println("4. Blue")
	fmt.Println("5. Magenta")
	fmt.Println("6. Cyan")
	fmt.Println("enter the numer of your choice")
	
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	
	if err != nil {
		fmt.Println("Invalid Input, Exiting")
		os.Exit(1)
	}
	var chosenchoice string

	switch choice {
	case 1:
		chosenchoice = Red
	case 2:
		chosenchoice = Green
	case 3:
		chosenchoice = Yellow
	case 4:
		chosenchoice = Blue
	case 5:
		chosenchoice = Magenta
	case 6:
		chosenchoice = Cyan
	default:
		fmt.Println("Invalid choice. Defaulting to standard text.")
		chosenchoice = Reset
	}


	fmt.Printf("%sGod is Good%s\n", chosenchoice, Reset)
}

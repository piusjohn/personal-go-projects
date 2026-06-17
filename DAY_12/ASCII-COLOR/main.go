package main

import (
	"fmt"
	"os"
)

// ANSI color escape codes
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
	fmt.Print("Enter the number of your choice: ")

	var choice int
	_, err := fmt.Scanf("%d", &choice)

	// Handle invalid inputs (like letters) or scanning errors
	if err != nil {
		fmt.Println("Invalid input. Exiting.")
		os.Exit(1)
	}

	var chosenColor string

	// Map user choice to the corresponding ANSI color code
	switch choice {
	case 1:
		chosenColor = Red
	case 2:
		chosenColor = Green
	case 3:
		chosenColor = Yellow
	case 4:
		chosenColor = Blue
	case 5:
		chosenColor = Magenta
	case 6:
		chosenColor = Cyan
	default:
		fmt.Println("Invalid choice. Defaulting to standard text.")
		chosenColor = Reset
	}

	// Print the colored message and immediately reset terminal colors
	fmt.Printf("%sHello, World!%s\n", chosenColor, Reset)
}

package main
import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    scanner.Scan()
    
    inputString := scanner.Text()
    fmt.Println("Hello, World.")
    
    fmt.Println(inputString)
}
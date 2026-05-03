// package main

// import "fmt"

// func compareTriplets(a []int, b []int) []int {
//     aliceScore := 0
//     bobScore := 0

//     for i := 0; i < 3; i++ {
//         if a[i] > b[i] {
//             aliceScore++
//         } else if a[i] < b[i] {
//             bobScore++
//         }
//     }

//     return []int{aliceScore, bobScore}
// }

// func main() {
//     a := []int{5, 6, 7}
//     b := []int{3, 6, 10}

//     result := compareTriplets(a, b)
//     fmt.Println(result)
// }


// package main

// import (
// 	"fmt"
// 	//"strings"
// )

// func main() {
	// fmt.Println("please enter a word")
	// var word string

	// fmt.Scan(&word)

// 	fmt.Println("enter an operation [1] lastword [2] capitalize [3] delete index")
// 	var operator string
// 	fmt.Scan(&operator)

// 	if operator == "1" {
// 		lastword := string(word[len(word)-1]) 
// 		fmt.Println((lastword))
// 	} else if operator == "2"{
// 		fmt.Println(strings.ToUpper(word))
// 	} else if  operator == "3" {
// 		fmt.Println("Please enter index to delete")
// 		var index int
// 		fmt.Scan(&index)

// 		word = word[:index] + word[index+1:]
// 		fmt.Println(word)
// 	}
// }



// arr := [][]int{
//     {2,3,4,5}, // Row 0
//     {4,5,6,7}, // Row 1
//     {4,5,3,4}, // Row 2
// }

// fmt.Println(len(arr)) // Output: 3

// }


// package main

// import (
//     "bufio"
//     "fmt"
//     "io"
//     "os"
//     "strconv"
//     "strings"
// )

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// func plusMinus(arr []int32) {
//     // Write your code here
//     var neg int
//     var plus int
//     n := len(arr)
//     for i := 0; i < n; i++{
//         if arr[i] < 0
//         neg = append(neg)
//     } else if arr[i] => 1{
//         plus = append(plus)
//     } 
//     result1 := len(neg) / len(arr)
//     result2 := len(plus) / len(arr)
//     return result1
//     return result2
//     return result3

// }

// f

package main

// import (
// 	"fmt"
// 	"strings"
// )
// func main() {
// 	daa := "hello good morning"
// 	dat := strings.Index(daa, "good")
// 	fmt.Println(dat)
// }




package main

import "fmt"
import "sort"

func main() {
	arr := []string{"aa", "bbbbbb", "c", "ddd", "ee"}
	fmt.Println(RemainderSorting(arr))
}
// func RemainderSorting(strArr []string) []string {
//     var reint []int
//     var result []string
//     for i := 0; i < len(strArr); i++ {
//         var res int
//         res = len(strArr[i]) % 3
//         reint = append(reint, res)
//     }

//     for i := 0; i < len(strArr)-1; i++{
//         for j := i+1; j < len(strArr); j++ {
//             if reint[i] > reint[j]{
//                 reint[i], reint[j] = reint[j], reint[i] 
//                 strArr[i], strArr[j] = strArr[j], strArr[i]
//             }
//         }
//     }
//     result = append(result, strArr...)
//     return result
// }


/////// SECOND APPROACH <<<<<<<<<<<<<<<


// func RemainderSorting(strArr []string) []string {
//     var result []string
//     var reint []int
//     //fmt.Print(strArr)
//     for i := 0; i < len(strArr); i++{
//         var res int
//        res = len(strArr[i]) % 3 
//        fmt.Print(res)
//        reint = append(reint, res)
//        if reint[i] == len(strArr[i]) %3 {
//          slices.Insert(result,  res, string(strArr[i]))
//        }
      
//     }
//     //slices.Sort(reint)
//     return result
// }
 

/////// IDIOMATIC APPROACH >>>>>>>>

func RemainderSorting(strArr []string) []string {
	sort.Slice(strArr, func(i, j int) bool {
		return len(strArr[i])%3 < len(strArr[j])%3
	})
	return strArr
}
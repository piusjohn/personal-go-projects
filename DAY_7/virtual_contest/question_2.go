package main

import (
	"fmt"

	//"golang.org/x/text/number"
)

func main() {
	var arr [5]int32
	//var result []int32
	arr = [5]int32{7, 69, 2, 221, 8974}

	 //  var maxsum int32
	//  min := 0
	//  max := 0
   sum := []int64{}
   add := int64(0)
    for i := 0; i < len(arr); i++ { 
		copy := append([]int32{}, arr[:i]...)
		copy = append(copy, arr[i+1:]...)

		for _, intg := range copy{
			add += int64(intg)

		}
		sum = append(sum, add)
		fmt.Println(sum)
		add = 0

    }
	min := sum[0]
	max := sum[0] 
	for _, char := range sum {
		if char > max {
			max = char
		}
		if char < min{
			min = char 
		}
	}
	fmt.Println(min, max)
}
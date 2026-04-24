package main

import (
	"fmt"
	"time"
)

// <<<<<< CONCURRENCY (GOROUTINES) >>>>>>
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3000 * time.Millisecond)
		fmt.Println(s)
	}
}

// func main() {
// 	go say("world")
// 	say("hello")


// }


//<<<<<<<<<<<< CHANNELS >>>>>>>>>>>>>>


func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	num := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(num[:len(num)/2], ch)
	go sum(num[len(num)/2:], ch)

	x, y := <-ch, <-ch 

	fmt.Println(x, y, x+y)

}


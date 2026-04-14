package main

import (
	"fmt"
)

type vehicle struct {
	brand string
	speed float64
}

type car struct {
	vehicle
	doors string
}

type boat struct {
	vehicle
	hasAnchor string
}

type mover interface {
	move() string
}

type describer interface {
	describe() string
}

func (c car) move() string {
	return "Car is driving"
}

func (b boat) move() string {
	return "Boat is sailing"
}

func (c car) describe() string {
	return c.brand + " "+ "has" + " " + c.doors + " " + "doors"

}

func (b boat) describe() string {
	return b.brand + " " + "has" + " " + b.hasAnchor +" " + "anchors"
}

func showInfo(des []describer) {
	for _, char := range des{
		fmt.Println(char)
	}
}

func startMoving(s []mover) {
	for _, char := range s {
		fmt.Println(char.move())
	}
}

func main() {
	c := car{vehicle: vehicle{brand: "volvo", speed: 100}, doors: "4"}
	b := boat{vehicle: vehicle{brand: "volvo", speed: 100}, hasAnchor: "No"}
	slice1 := []mover{b, c}
	slice2 := []describer{b, c}
	startMoving(slice1)
	showInfo(slice2)
}
package main

import (
	"fmt"
)

type Animal interface {
	sound() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (D Dog) sound() string {
	return "Woof"
}

func (C Cat) sound() string {
	return "Meow"
}

func makeNoise(A []Animal) {
	for _, char := range A {
		fmt.Println(char.sound())
	}
}

func main() {
	d := Dog{name: "tig"}
	c := Cat{name: "lacer"}
	s := []Animal{d, c}
	makeNoise(s)
}

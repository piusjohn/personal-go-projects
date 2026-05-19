package main

import "fmt"

// Combine small interfaces to create a larger one
type Athlete interface {
    Walker
    Runner
}

func main() {
    var athlete Athlete = Human{}

    athlete.Walk()
    athlete.Run()
}

type Walker interface { // Interface for a walker
    Walk()
}

type Runner interface { // Interface for a runner
    Run()
}

type Human struct{} // Implement the interfaces in a concrete type

func (h Human) Walk() { // Implement the Walker interface
    fmt.Println("Human is walking.")
}

func (h Human) Run() { // Implement the Runner interface
    fmt.Println("Human is running.")
}
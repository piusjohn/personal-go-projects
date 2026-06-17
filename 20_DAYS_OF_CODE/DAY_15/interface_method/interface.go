package main

import "fmt"

type Vehicle interface{
	CountRentalcost() int
}
func PrintRentalCost(v Vehicle) {
	fmt.Printf("Rental Cost: $%d\n", v.CountRentalcost())
} 

func main() {
  var vehicle Vehicle 
  vehicle = Car{hours: 2, miles: 4}
  PrintRentalCost(vehicle)
	// s := Scooter{hours: 5}
	// c := Car{hours:5, miles: 5}
	// PrintRentalCost(s)
	// PrintRentalCost(c)
}

type Scooter struct{
	hours int
}

type Car struct{
	miles int
	hours int
}

func(s Scooter) CountRentalcost() int{
	const hourlyRate int = 5
	return s.hours * hourlyRate
}

func(c Car) CountRentalcost() int{
	const hourlyRate int = 10 
	const mileageRate int = 5
	return (c.hours * hourlyRate) + (c.miles * mileageRate)
}
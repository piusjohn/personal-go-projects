// package main

// import "fmt"

// type Speaker interface {
//     Speak() string
// }

// type Dog struct{}
// func (d Dog) Speak() string {
//     return "Woof"
// }

// type Cat struct{}
// func (c Cat) Speak() string {
//     return "Meow"
// }
// type goat struct{}
// func (g goat) Speak() string {
// 	return "meeeeh"
// }

// func MakeSound(s Speaker) {
//     fmt.Println(s.Speak())
// }

// func main() {
//     k := Dog{}
//     c := Cat{}
// 	d := goat{}

// 	MakeSound(d)
//     MakeSound(k)
//     MakeSound(c)
// }


//##2

// package main

// import "fmt"

// type Card struct{}
// func (c Card) Pay(amount float64) {
//     fmt.Println("Paid with card:", amount)
// }

// type BankTransfer struct{}
// func (b BankTransfer) Pay(amount float64) {
//     fmt.Println("Paid with bank transfer:", amount)
// }

// func main() {
//     card := Card{}
//     bank := BankTransfer{}

//     card.Pay(100)
//     bank.Pay(200)
// }

package main

import (
	"fmt"
	"time"
)

	type User struct {
    ID        int
    Name      string
    Email     string
    CreatedAt time.Time
}
func (u User) DisplayName() string {
    return u.Name
}

func main () {

	user := User {
		ID: 1,
		Name: "Pius John",
		Email: "pius622472@gmail.com",
		CreatedAt: time.Now(),

	}
	val := user.DisplayName()
	fmt.Println(val)
}   
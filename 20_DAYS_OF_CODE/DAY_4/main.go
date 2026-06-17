package main
import (
	"fmt"
	"sort"
)
func main() {
	menu := map[string]float64{
    "eggs":    1.75,
    "bacon":   3.22,
    "sausage": 1.89,
}
// var total float64
// for dish, price := range menu{
// 	total += price
// 	fmt.Printf("Remember to buy %s\n", dish)	
// }
// fmt.Printf("Remember to pay %.2f for the meal\n", total)

//>>>>>>>>>>>> ORDERING IN MAP <<<<<<<<<<<<<
delete(menu, "bacon")
menu["biscuit"] = 0.00
var dishes []string
for dish := range menu{
	dishes = append(dishes, dish)
}
sort.Strings(dishes)
for _, char := range dishes {
	fmt.Println(char)
}

//>>>>>> CHECKING IF A KEY EXIST IN A MAP <<<<<<<
if _, ok := menu["biscuit"]; ok {
	fmt.Println("key exists")
} else {
	fmt.Println("key doesn't exist")
}
}

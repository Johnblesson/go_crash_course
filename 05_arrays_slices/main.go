package main
import "fmt"

func main() {
	// Arrays
	// var fruitArr [3]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"
	// fruitArr[2] = "Mango"

	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
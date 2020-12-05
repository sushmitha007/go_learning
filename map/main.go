package main

import "fmt"

func main() {
	//Creating map with literal way
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	//2nd type
	// var colors map[string]string

	//3rd type
	colors := make(map[string]string)

	colors["white"] = "ffffff"
	fmt.Println(colors)

	//Deleting key value pair
	delete(colors, "white")
	fmt.Println(colors)
}

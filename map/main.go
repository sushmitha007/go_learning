package main

import "fmt"

func main() {
	//Creating map with literal way
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//2nd type
	// var colors map[string]string

	//3rd type
	// colors := make(map[string]string)

	// updating map
	// colors["white"] = "ffffff"
	// fmt.Println(colors)

	//Deleting key value pair
	// delete(colors, "white")

	fmt.Println(colors)
	printMap(colors)
}

// iterating over maps
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is", hex)
	}
}

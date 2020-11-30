package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	//Declaring the new slice
	// cards := []string{"ace of diamonds", newCard()}

	// Declaring as type deck
	cards := deck{"ace of diamonds", newCard()}

	// Appending new element into slice
	cards = append(cards, "Six of Space")

	// // Iterating over a slice
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// call the function from deck.go
	cards.print()

	fmt.Println(cards)
}
func newCard() string {
	return "Five of Diamonds"
}

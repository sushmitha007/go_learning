package main

func main() {
	// 	// var card string = "Ace of spades"
	// 	// card := "Ace of Spades"
	// 	// card = "Five of Diamonds"
	// 	// card := newCard()

	// 	//Declaring the new slice
	// 	// cards := []string{"ace of diamonds", newCard()}

	// 	// Declaring as type deck
	// 	cards := deck{"ace of diamonds", newCard()}

	// 	// Appending new element into slice
	// 	cards = append(cards, "Six of Space")

	// 	// // Iterating over a slice
	// 	// for i, card := range cards {
	// 	// 	fmt.Println(i, card)
	// 	// }

	// 	// call the function from deck.go
	// 	cards.print()

	// 	// fmt.Println(cards)

	// 	//Creating hand of cards
	// 	hand, remainingDeck := deal(cards, 2)
	// 	hand.print()
	// 	remainingDeck.print()

	// Type conversion
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// converting deck into strings

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// save cards to hardrive
	// cards.saveToFile("myCards")

	// Load the cards from drive
	// cards := newDeckFromFile("myCards")
	// cards.print()

	// load the wrong file
	// cards := newDeckFromFile("my")
	// cards.print()

	//shuffle the deck
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

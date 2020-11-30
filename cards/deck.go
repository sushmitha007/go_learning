package main

import "fmt"

// Create new type of deck, whcih is slice of strings

type deck []string

//Reciever function
// Any Variable of type "deck" gets acces to "print" method
// here "d" is receiver of type "deck"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Returning multiple values and using range inside slice
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

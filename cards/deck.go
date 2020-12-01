package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {

	// type conversion from deck to slice of strings
	// Join the slices into single slice
	return strings.Join([]string(d), ",")

}

// function to create deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	// replace the variable with underscore when the varibles are not used anywhere in the code
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}

	}
	return cards
}

//saves a file into hardrive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//get cards from hard drive
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option1: log the error and go back newDeck() function
		// option2: log the error and entirely quit the program

		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

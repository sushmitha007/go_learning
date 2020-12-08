package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// Problem without interface

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb)
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb)
// }

// With Interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating english greeting
	return "Hi,There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

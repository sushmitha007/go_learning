package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
func printGreeting(eb englishBot) {
	fmt.Println(eb)
}
func printGreeting(sb spanishBot) {
	fmt.Println(sb)
}
func (englishBot) getGreeting() string {
	//VERY custom logic for generating english greeting
	return "Hi,There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

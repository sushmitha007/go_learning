package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	//Declaring Structs
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// other way.this assigns empty strings
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 56738,
		},
	}
	// fmt.Printf("%+v", jim)
	//Creating the pointer to modify the original content
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// Alternative way of declaring and defining pointer.go will automatically turn into pointer
	jim.updateName("Jimmy")
	// prints the values in struct
	// fmt.Println(*&jim)
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}

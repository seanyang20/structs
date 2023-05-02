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
	contactInfo // we don't have to specify field name for embedded structs if we don't want to
}

func main() {
	// --------------
	// declaring structs
	// ---------------

	// best way of declaring struct
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// another way of declaring struct
	// alex := person{"Alex", "Anderson"}

	// one more way of declaring struct
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// --------------
	// embedding structs
	// ---------------

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // unlike other languages, every line needs a comma
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// You can call this function that we are about to define on any type of type person
func (p person) print() {
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type contactInfo struct {
	email       string
	postal_code string
}

type Person struct {
	firstName string
	surname   string
	contactInfo
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
	fmt.Printf("\n")
}

func updateSlice(s []string, i int, v string) {
	s[i] = v
}

func updateName(n string, v string) {
	n = v
}

func printPointer(namePointer *string) {
	// This memory address will be different from the memory address passed into
	// the function as Go is a pass by value language
	fmt.Println(&namePointer)
}

func main() {
	person := Person{
		firstName: "Bill",
		surname:   "Beaumont",
		contactInfo: contactInfo{
			email:       "bill.b@domain.com",
			postal_code: "AA1 1AA",
		},
	}

	person.print()

	// get a pointer to the person struct
	personPointer := &person
	fmt.Println(personPointer)
	// use the pointer with the receiver function
	personPointer.updateName("Billy")
	// print out the update
	person.print()

	// This will result in the slice being updated...as a slice data structure has
	// a pointer to head included...
	// slice data structure are known as reference types
	s := []string{
		"Hi",
		"there",
		"how",
		"are",
		"you",
	}
	updateSlice(s, 0, "Bye")
	fmt.Println(s)

	// This will result in the value not being updated...
	name := "Nichola"
	updateName(name, "Nicky")
	fmt.Println(name)
	// same as above...
	fmt.Println(*&name)

	namePointer := &name
	fmt.Println(namePointer)
	printPointer(namePointer)
}

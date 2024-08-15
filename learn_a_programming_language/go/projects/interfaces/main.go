package main

import "fmt"

type bot interface {
	// If you have a function called getGreeting that returns a string
	// you can use this interface and the subsequent functions of type bot
	// e.g. printGreeting with b.getGreeting() would execute getGreeting() for
	// that receiver. Interfaces can contain many functions which each type must have to
	// utilise this interface
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Printf("%s \n\n", b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

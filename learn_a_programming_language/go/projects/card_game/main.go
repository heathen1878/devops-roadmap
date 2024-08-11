package main

import (
	"fmt"
)

// variables can be defined outside of the main function

func main() {
	cards := newDeck()

	cards.shuffle()

	hand, _ := deal(cards, 5)

	fmt.Println("Dealt hand...")
	hand.print()
	fmt.Println()
	//fmt.Println("Remainder of deck...")
	//remainingDeck.print()

	//cards.writeToFile("cards.txt")

	//deckFromFile := readFromFile("cards.txt")
	//fmt.Println()
	//fmt.Println("Deck from file...")
	//fmt.Println(deckFromFile)
}

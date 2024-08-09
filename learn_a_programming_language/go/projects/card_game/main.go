package main

import "fmt"

// variables can be defined outside of the main function

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	fmt.Println(cards.toString())

	cards.writeToFile("cards.txt")
}

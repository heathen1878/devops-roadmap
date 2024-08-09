package main

import (
	"fmt"
	"os"
	"strings"
)

// new type which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Clubs",
		"Diamonds",
		"Hearts",
		"Spades",
	}

	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// grant any variables of type deck get access to this function / method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(file string) error {
	return os.WriteFile(file, []byte(d.toString()), 0666)
}

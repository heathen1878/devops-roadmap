package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
			// loop through the card suits and card values and concatenate the suit and values
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// grant any variables of type deck get access to this function / method
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	//return to decks, one with 0 to handsize and one with handsize and remainder
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	// Join the slice of string with a comma
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(file string) error {
	// convert the deck to a string using toString() then convert to byte slice
	return os.WriteFile(file, []byte(d.toString()), 0666)
}

func readFromFile(file string) deck {
	bs, _err := os.ReadFile(file)
	if _err != nil {
		// Something went wrong
		fmt.Println("Error: ", _err)
		os.Exit(1)
	}
	// Byte slice to string then split by comma and assign to ss
	ss := strings.Split(string(bs), ",")
	// return type of deck
	return deck(ss)
}

func (d deck) shuffle() {
	// Create an Int64 to be the seed of rand
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	//  loop through the deck getting the index
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

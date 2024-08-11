package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewDeck(t *testing.T) {
	t.Parallel()
	d := newDeck()

	// The deck should have 52 cards
	want_deck_length := 52
	got_deck_length := len(d)

	if !cmp.Equal(want_deck_length, got_deck_length) {
		t.Errorf("Expected %v cards got %v", want_deck_length, got_deck_length)
	}

	want_first_card := "Ace of Clubs"
	got_first_card := d[0]

	if !cmp.Equal(want_first_card, got_first_card) {
		t.Errorf("Expected %v got %v", want_first_card, got_first_card)
	}

	want_last_card := "King of Spades"
	got_last_card := d[51]

	if !cmp.Equal(want_last_card, got_last_card) {
		t.Errorf("Expected %v got %v", want_last_card, got_last_card)
	}
}

func TestDeal(t *testing.T) {
	t.Parallel()

	d := newDeck()
	handSize := 5

	hand, remainder := deal(d, handSize)

	want_hand := 5
	got_hand := len(hand)

	want_remainer := 47
	got_remainer := len(remainder)

	if !cmp.Equal(want_hand, got_hand) {
		t.Errorf("Expected %v got %v", want_hand, got_hand)
	}

	if !cmp.Equal(want_remainer, got_remainer) {
		t.Errorf("Expected %v got %v", want_remainer, got_remainer)
	}
}

func TestToString(t *testing.T) {
	t.Parallel()

	d := newDeck()
	s := d.toString()

	want_type := reflect.TypeOf("string")
	got_type := reflect.TypeOf(s)

	if want_type != got_type {
		t.Errorf("Expected a %v got a %v", want_type, got_type)
	}
}

func TestCreateAndReadFromFile(t *testing.T) {
	_err := os.Remove("_decktesting")
	if _err != nil {
		fmt.Printf("Failed to %s", _err.Error())
	}

	d := newDeck()

	_err = d.writeToFile("_decktesting")

	if _err != nil {
		t.Errorf("Expected error of nil got %v", _err)
	}

	ld := readFromFile("_decktesting")

	want := 52
	got := len(ld)

	if !cmp.Equal(want, got) {
		t.Errorf("Expected %v got %v", want, got)
	}

	_err = os.Remove("_decktesting")
	if _err != nil {
		fmt.Printf("Failed to %s", _err.Error())
	}
}

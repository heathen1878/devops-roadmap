# Card Game

## Main

The entry point is main.go and all the of the functions live in deck.go.

### Functions

#### Receiver

```go
func (r receiverType) function_name () {
    // some code
}
```

#### Argument

```go
func function_name (a argumentType) {
    // some code
} 
```

a function can combine both receiver and arguments. 

## Tests

All tests live in deck_test.go.

```go
// type Card struct {
// 	suit  string
// 	value string
// }

// func NewDeck() Card {
// 	cards := map[int]Card{}

// 	suits := []string{
// 		"Clubs",
// 		"Diamonds",
// 		"Hearts",
// 		"Spades",
// 	}

// 	values := []string{
// 		"Ace",
// 		"Two",
// 		"Three",
// 		"Four",
// 		"Five",
// 		"Six",
// 		"Seven",
// 		"Eight",
// 		"Nine",
// 		"Ten",
// 		"Jack",
// 		"Queen",
// 		"King",
// 	}

// 	for i, suit := range suits {
// 		for _, value := range values {
// 			// loop through the card suits and card values and concatenate the suit and values
// 			cards[i].suit(suit)
// 			cards[i].value(value)
// 		}
// 	}

// 	return cards
// }
```
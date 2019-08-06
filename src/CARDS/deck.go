package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cars := deck{"Ace of Spades", "Two of Spades"}
}

func (myDeck deck) print() {
	for i, card := range myDeck {
		fmt.Println(i, card)
	}
}

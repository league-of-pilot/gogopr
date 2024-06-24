package main

import "fmt"

type deck []string

// receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// d -> convention
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// demo only -> go ko có KN destructing như js, bản thân việc return multi value đã tương tự array rồi
func dealor(d deck, handSize int) [2]deck {
	return [2]deck{d[:handSize], d[handSize:]}
}

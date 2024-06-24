package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// var another = "Ten of Diamonds"
	// cardFoo := newCard()

	// another = "Queen of Hearts"
	// fmt.Println(card, another, cardFoo)
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Queen of Hearts")

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

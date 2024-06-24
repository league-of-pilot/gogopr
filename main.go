package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// var another = "Ten of Diamonds"
	// cardFoo := newCard()

	// another = "Queen of Hearts"
	// fmt.Println(card, another, cardFoo)
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Queen of Hearts")

	// for _, card := range cards {
	// 	fmt.Println(card)
	// }
	cards := newDeck()
	// cards.print()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()

	// ko hay
	// lor := dealor(cards, 5)
	// // Hàm print ko return mà print thẳng nên ko dùng log như js được
	// lor[0].print()
	// lor[1].print()
}

func newCard() string {
	return "Five of Diamonds"
}

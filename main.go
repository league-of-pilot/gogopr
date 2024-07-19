package main

import (
	"fmt"
	"go-learn/chanogo"
	"go-learn/talkConc"
	"go-learn/udego"
	"os"
)

func main() {
	args := os.Args
	var ex string

	if len(args) > 1 {
		ex = args[1]
	} else {
		ex = "method-19"
	}

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
	// cards := newDeck()
	// cards.print()

	// Deal card
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// demo lor deal card
	// ko hay
	// lor := dealor(cards, 5)
	// // Hàm print ko return mà print thẳng nên ko dùng log như js được
	// lor[0].print()
	// lor[1].print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards._temp_.txt")

	switch ex {
	case "39":
		ex39()
	case "59":
		udego.Ex59()
	case "62":
		udego.HttpV62()
	case "method-19":
		udego.Method19()
	case "20":
		udego.Main20()
	case "23":
		udego.MainMethod23()
	case "demoReader":
		udego.PlayReader()
	case "chano":
		chanogo.CheckLinks()
	case "chantour":
		chanogo.ChanTour()
	case "chanbuf":
		chanogo.ChanBuf()
	case "selecterr":
		chanogo.SelectErr()
	case "selectno":
		chanogo.SelectUnknown()
	case "conc22":
		talkConc.MainTalkConc22()
	case "fanin":
		talkConc.XmainFanIn()
	case "struct":
		structPrivate := chanogo.CreatePrivate()
		newAge := structPrivate.Age + 20
		println(newAge)

		// var pubStruct chanogo.StructPub

		pubStruct := chanogo.StructPub{
			PubAge: 20,
		}
		priStr := chanogo.structExport{
			Age: 20,
		}
		createByFunc := chanogo.CreatePublic()
		println(pubStruct, priStr, createByFunc)
	default:
		deckFromFile := udego.NewDeckFromFile("my_cards._temp_.txt")
		// deckFromFile.print()
		println(len(deckFromFile))
		// exercise39()

		demoStruct()
	}
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

func ex39() {
	// ints := []int
	// len := 10
	// var ints [len]int ❌
	len := 11
	ints := make([]int, len)

	for i := 0; i < len; i++ {
		// ints = append(ints, i) -> ko hay
		ints[i] = i
	}

	for _, i := range ints {
		if i%2 == 0 {
			println(i, " is even")
		} else {
			println(i, " is odd")
		}
	}
}

type pInfo struct {
	age  int
	mail string
}

type person struct {
	firstName string
	lastName  string
	pInfo     // shorthand for pInfo pInfo
}

func demoStruct() {
	person1 := person{firstName: "Alex", lastName: "Anderson", pInfo: pInfo{
		age:  20,
		mail: "mail",
	}} // , cuối ở đây thì skip được -> nhờ editor hint error syntax

	person2 := person{"John", "Doe", pInfo{
		age:  20,
		mail: "mail", // bắt buộc phải có dấu , ở cuối
	}} // valid but not recommended

	fmt.Println(person1, person2)

	var person3 person // zero value chứ ko phải undefined như js
	fmt.Printf("%+v", person3)
	person3 = person2
	person2.lastName = "rename"

	fmt.Printf("%+v", person3) // -> ko bị reference
	person2.print("person2")
}

// https://bojanz.github.io/optional-parameters-go/
// go ko support built-in -> lí do ở trên
func (p person) print(s string) {
	fmt.Printf("%s %+v", s, p)
}

package udego

// syntax import bằng newline, ko dùng dấu phẩy
import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	// https://go.dev/tour/basics/13
	// ko thấy docs cụ thể
	return strings.Join([]string(d), ",")

	// return fmt.Sprint(d) -> ❌ , return [Ace of Spades Two of Spades Three of Spades Four of Spades Ace of Diamonds ... ]
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func NewDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	// https://pkg.go.dev/math/rand@go1.22.4#New
	// https://pkg.go.dev/time@go1.22.4#Time.UnixNano
	// https://pkg.go.dev/math/rand@go1.22.4#Source

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

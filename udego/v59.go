package V59

import "fmt"

// v59 interface
type botInterstellar interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getBye() {
	fmt.Println("Bye!")
}

func Ex() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	eb.getBye()
}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b botInterstellar) {
	fmt.Println(b.getGreeting())
}

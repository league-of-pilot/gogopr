package udego

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) String() string {
	return fmt.Sprintf("Success ?? %v", e.What)
}

// Nếu comment đi thì String() ở trên sẽ được gọi
// Tuy nhiên ko phải do order
// The Error() method has a higher precedence when dealing with types that implement the error interface in Go.
// type error Bị implement ngầm
// Phải explicit gọi .String()
func (e MyError) Error() string {
	return fmt.Sprintf("mine %v, %s",
		e.When, e.What)
}

// Vì struct value là non-pointer , ko thể là nil
// Phải return về dạng pointer mới so sánh với nil được
func run() *MyError {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Method19() {
	okela := run()

	if okela != nil {
		fmt.Println(okela)
	}

	if err := runWeb(); err != nil {
		fmt.Println(err)
	}
}

// ====================
type MyErrorWeb struct {
	When time.Time
	What string
}

func (e *MyErrorWeb) Error() string {
	return fmt.Sprintf("web %v, %s",
		e.When, e.What)
}

func runWeb() error {
	return &MyErrorWeb{
		time.Now(),
		"it didn't work",
	}
}

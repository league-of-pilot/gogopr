package chanogo

import (
	"fmt"
	"time"
)

type tour struct {
	sum int
	arr []int
}

const INT = 7
const DELAY = false

func sum(s []int, c chan tour) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	if s[0] != INT && DELAY {
		fmt.Println("delay", s)
		time.Sleep(100 * time.Millisecond)
	}
	c <- tour{sum, s}
}

func ChanTour() {
	s := []int{INT, 2, 8, -9, 8, 20, 5, -53, -20}

	c := make(chan tour)
	half := len(s) / 2
	fmt.Println("half", half, "len", len(s))
	s1 := s[:half]
	s2 := s[half:]
	go sum(s1, c)
	go sum(s2, c)
	x, y := <-c, <-c // receive from c

	fmt.Println(s1, s2)
	fmt.Println(x, y, x.sum+y.sum)
}

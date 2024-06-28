package chanogo

import "fmt"

// https://go.dev/tour/concurrency/5
func fibonacciSelectUnknown(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func SelectUnknown() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 1000
	}()
	go func() {
		fmt.Println("b4")
		fmt.Println(<-c + 100)
		fmt.Println("a8")
	}()
	go func() {
		for {
			fmt.Println(<-c + 1000)
		}
	}()
	fibonacciSelectUnknown(c, quit)
}

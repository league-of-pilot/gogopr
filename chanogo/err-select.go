package chanogo

import "fmt"

func fibonacci_err(c chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			if x > 40 {
				println("stop ", x)
				return
			}
		}
	}
}

// Điều kiện stop ko ổn định,

// https://go.dev/tour/concurrency/5
func SelectErr() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	fibonacci_err(c)
}

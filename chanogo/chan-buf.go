package chanogo

import "fmt"

func ChanBuf() {
	// try to remove buffer value or change to see different log
	ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// ch <- 3 // deadlock

	// run this without goroutine will cause deadlock
	go push(ch, []int{1, 2, 3})
	fmt.Println(<-ch + 100)
	fmt.Println(<-ch + 200)
}

func push(ch chan int, ints []int) {
	for _, i := range ints {
		println("pre-push", i)
		ch <- i
		println("post-push", i)
	}
	close(ch)
}

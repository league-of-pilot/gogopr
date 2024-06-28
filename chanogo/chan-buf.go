package chanogo

import "fmt"

func ChanBuf() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3
	fmt.Println(<-ch * 10)
	fmt.Println(<-ch * 7)
}

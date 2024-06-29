package talkConc

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller.
}

// https://go.dev/talks/2012/concurrency.slide#26
func MainTalkConc22() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// https://go.dev/talks/2012/concurrency.slide#27
// Multiplexing ??
// ý tưởng hay nhưng khi log ra sẽ ko kiểm soát được thứ tự nữa
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	kill := time.After(5 * time.Second)
	// go func() {
	// 	for {
	// 		c <- <-input1
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c <- <-input2
	// 	}
	// }()

	// gọn hơn với select
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case <-time.After(1 * time.Second):
				// instant kill if no message in 1 second
				// demo syntax for this case only, cannot reach in this setup
				fmt.Println("You're too slow.")
				return
			case <-kill:
				fmt.Println("kill")
				// deadlock main vì main vẫn tiếp tục receive
				return
			}
		}
	}()
	return c
}

func XmainFanIn() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

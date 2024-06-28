package chanogo

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// main cũng là 1 goroutine nên có thể dùng time.Sleep
const DELAY_TIME = time.Second * 5
const CASE = 3
const INDEX_SKIP = 9000

// const ko support slice, map, array
var links = []string{
	// "http://youtube.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://google.com",
	"http://golang.org",
	"http://amazon.com",
}

// go routine fireup seperately and init log does not comes in order
func checkLink(link string, c chan string, i int) {
	if i != INDEX_SKIP {
		println("init routine", i)
	}
	// Set delay ở đây ko đúng về logic vì add side effect vào func
	// time.Sleep(DELAY_TIME)
	_, err := http.Get(link)

	// Có thể gom style if-else res nhưng tách ra để dễ hack
	var res string
	if err != nil {
		res = strconv.Itoa(i) + " Error:" + err.Error()
		println(res)
		c <- link
		return
	}
	// Deadlock if put here
	// tmp := <-c
	// fmt.Println(tmp)

	res = strconv.Itoa(i) + " -> " + link + " is up!"
	println(res)
	c <- link
}

func CheckLinks() {
	c := make(chan string)

	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for i, link := range links {
		go checkLink(link, c, i)
		// Đặt đây ko khác gì synchronous
		// fmt.Printf(<-c)
	}

	// Nếu ko comment loop trên lại thì ko còn value receive từ channel
	switch CASE {
	case 1:
		simpleLoop(len(links), c)
	case 2:
		infiniteLoop(c)
	case 3:
		alternativeLoopSyntax(c)
	default:
		simpleReceiver(c)
	}

	// ??? nếu đặt 2 inf loop ???
	// Vì chạy tuần tự nên loop trên sẽ chặn đứng vòng chạy và ko xuống dưới đây được
	// Phải set loop trong 1 goroutine khác ?!
	println("FINAL EXIT SWITCH CASE")
	anotherReceiver(c)
}

// ========================================
// Receiver setup
// ========================================

func anotherReceiver(c chan string) {
	mes := <-c
	println("WHAT HAPPENED ", mes)
}

func simpleReceiver(c chan string) {
	mes := <-c
	fmt.Println(mes)
	fmt.Printf(<-c)
	println("Done - 1")
	fmt.Printf(<-c)
	println("Done - 2")
	fmt.Printf(<-c)
}

func simpleLoop(len int, c chan string) {
	for i := 0; i < len; i++ {
		// This log COULD (not always) come up 1st before goroutine init complete
		// goroutine fire up took some delay

		println("Prepare receiving", i)
		fmt.Printf(<-c)
		println(" ==> Receive from channel", i)
	}
}

func infiniteLoop(c chan string) {
	for {
		// delay ở đây ko đúng vì các request sau chạy xong phải + 5s của main nữa mới chạy tiếp
		// time.Sleep(DELAY_TIME)
		// function literal ~ anonymous function in JS
		go checkLink(<-c, c, INDEX_SKIP)
	}
}

// Prefer syntax này hơn
func alternativeLoopSyntax(c chan string) {
	// Vì go tự ngầm hiểu <-c trả về string nhưng code nhìn sẽ rối
	// Bọc lại để dễ hình dung
	// l := <-c
	for l := range c {
		go func() {
			time.Sleep(DELAY_TIME)
			checkLink(l, c, INDEX_SKIP)
		}()
	}
}

package chanogo

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetLink() []string {
	links := []string{
		"http://youtube.com",
		// "http://facebook.com",
		"http://stackoverflow.com",
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
	}
	return links
}

// go routine fireup seperately and init log does not comes in order
func CheckLink(link string, c chan string, i int) {
	println("init routine", i)
	_, err := http.Get(link)

	// Có thể gom style if-else res nhưng tách ra để dễ hack
	var res string
	if err != nil {
		res = strconv.Itoa(i) + " Error:" + err.Error()
		c <- res
		return
	}
	// Deadlock if put here
	// tmp := <-c
	// fmt.Println(tmp)

	res = strconv.Itoa(i) + " -> " + link + " is up!"
	c <- res
}

func CheckLinks() {
	links := GetLink()
	c := make(chan string)

	println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	for i, link := range links {
		go CheckLink(link, c, i)
		// Đặt đây ko khác gì synchronous
		// fmt.Printf(<-c)
	}

	// mes := <- c
	// fmt.Println(mes)
	// fmt.Printf(<-c)
	// println("Done - 1")
	// fmt.Printf(<-c)
	// println("Done - 2")
	// fmt.Printf(<-c)

	for i := 0; i < len(links); i++ {
		// This log COULD (not always) come up 1st before goroutine init complete
		// goroutine fire up took some delay

		println("Prepare receiving", i)
		fmt.Printf(<-c)
		println("Receive from channel", i)
	}
}

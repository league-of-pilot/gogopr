package chanogo

import (
	"fmt"
	"net/http"
)

func GetLink() []string {
	links := []string{
		"http://youtube.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://google.com",
		"http://golang.org",
		// "http://amazon.com",
	}
	return links
}

func CheckLink(link string, c chan string, i int) {
	var res string
	_, err := http.Get(link)
	if err != nil {
		res = string(i) + " Error:" + err.Error()
		c <- res
		return
	}
	// Deadlock if put here
	// tmp := <-c
	// fmt.Println(tmp)

	res = string(i) + link + " is up!"
	c <- res
}

func CheckLinks() {
	links := GetLink()
	c := make(chan string)

	for i, link := range links {
		go CheckLink(link, c, i)
		// Đặt đây ko khác gì synchronous
		// fmt.Printf(<-c)
	}

	// mes := <- c
	// fmt.Println(mes)
	fmt.Printf(<-c)
}

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
		"http://golang.org",
		"http://google.com",
		"http://amazon.com",
	}
	return links
}

func CheckLink(link string, c chan string, i int) {
	var res string
	_, err := http.Get(link)
	if err != nil {
		res = string(i) + " Error:" + err.Error()
		fmt.Println(res)
		c <- res
		return
	}
	res = string(i) + link + " is up!"
	fmt.Println(res)
	c <- res
}

func CheckLinks() {
	links := GetLink()
	c := make(chan string)

	for i, link := range links {
		go CheckLink(link, c, i)
	}

	// mes := <- c
	// fmt.Println(mes)
	fmt.Printf(<-c)
}

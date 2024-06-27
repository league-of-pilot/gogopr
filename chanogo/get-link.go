package chanogo

import (
	"fmt"
	"net/http"
)

func GetLink() []string {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	return links
}

func CheckLink(link string, i int) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(i, "Error:", err)
		return
	}
	fmt.Println(i, link, "is up!")
}

func CheckLinks() {
	links := GetLink()
	for i, link := range links {
		CheckLink(link, i)
	}
}

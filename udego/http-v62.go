package udego

import (
	"net/http"
	"os"
)

func Http() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	println(resp)
}

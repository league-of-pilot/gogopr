package udego

import (
	"fmt"
	"net/http"
	"os"
)

func HttpV62() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
	println("Status code:", resp.StatusCode)
	// println(resp) // -> trả pointer
	// println(&resp) // -> vẫn là pointer

}

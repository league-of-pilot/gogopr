package udego

import (
	"io"
)

type demoRead struct {
	content string
}
type demoWrite struct{}

func PlayReader() {
	io.Copy(demoWrite{}, demoRead{
		content: "This demo some custom read implementation but may have bugs, check with GPT",
	})
}

func (demoRead) Read(b []byte) (n int, err error) {
	return len(b), nil
}

func (demoWrite) Write(bs []byte) (int, error) {
	return len(bs), nil
}

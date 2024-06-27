package udego

import (
	"fmt"
	"io"
)

type demoRead struct {
	content string
	pos     int
}
type demoWrite struct{}

func PlayReader() {
	io.Copy(demoWrite{}, &demoRead{
		content: "This demo some custom read implementation but may have bugs, check with GPT",
	})
}

func (p *demoRead) Read(b []byte) (n int, err error) {
	l := len(p.content)
	if p.pos >= l {
		println("EOF")
		return 0, io.EOF
	}

	for i := 0; i < l && i < len(b); i++ {
		iter := p.content[i]
		b[i] = iter
		p.pos++
	}

	return p.pos, nil
}

func (demoWrite) Write(bs []byte) (int, error) {
	fmt.Println(string(bs), len(bs))
	return len(bs), nil
}

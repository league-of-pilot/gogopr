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
		// n vì define ở dạng return nên initial sẽ là 0
		return n, io.EOF
	}

	// track theo pos của content < l chứ ko phải i
	// mỗi loop sẽ tăng pos lên theo i
	// i chỉ là biến chạy ảo chỉ ràng buộc lỏng với len(b)
	// Dùng closure quản lý thì hơi căng

	for i := 0; p.pos < l && i < len(b); i++ {
		iter := p.content[p.pos]
		b[i] = iter
		p.pos++
		n++
	}

	// The Read method should return the number of bytes read in this call, not the total position (p.pos).
	return n, nil
}

func (demoWrite) Write(bs []byte) (int, error) {
	fmt.Println(string(bs), len(bs))
	return len(bs), nil
}

package udego

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	// Chỉ đảm bảo cho r bên trong
	// rot13.r.Read() method -> exist
	// Nhưng rot13.Read() chưa có, cần implement
	r io.Reader
}

type logWriter struct{}

// Bản thân rot13Reader cần implement Read() method để cùng có dạng io.Reader
// arr b[] byte sẽ được io.Copy() tạo buffer và truyền vào
func (sctRot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = sctRot.r.Read(b)
	// b sau khi Read() sẽ được populate data từ string
	// bắt đầu dùng n để loop ngược lại b[] và decode
	for i := 0; i < n; i++ {
		// thuật toán thì ref wiki, copy ko sao
		// modified trực tiếp b slice
		if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
			b[i] += 13
		} else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return
}

func MainMethod23() {
	// s với setup này đã thỏa interface io.Reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

	// string.NewReader() trả về *strings.Reader nên ko thể dùng r2:=r
	// Bản thân expression trên trả về copy value nhưng valua là pointer lại cùng trỏ về 1 biến
	// Khi này strings đã ở EOF, io.Copy() gặp EOF sẽ dừng, ko chạy tiếp vào Write nên ko bắt được trong logic nữa
	// muốn đọc lại thì phải reset con trỏ position
	justTest := strings.NewReader("Just print me!")
	io.Copy(logWriter{}, justTest)
	io.Copy(logWriter{}, justTest) // nothing happen
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs), len(bs))
	// bs là []byte, nếu convert sang string thì sẽ có thể dùng string method
	// nhưng ko cần thiết, chỉ cần dùng len() để lấy size
	// trả về size và nil
	return len(bs), nil
}

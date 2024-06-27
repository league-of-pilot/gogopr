package udego

import (
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
}

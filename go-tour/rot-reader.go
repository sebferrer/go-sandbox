package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	length, err := r.r.Read(b)

	for i := 0; i < length; i++ {
		if b[i] <= 'Z' && b[i] >= 'A' {
			b[i] = (b[i]-'A'+13)%26 + 'A'
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = (b[i]-'a'+13)%26 + 'a'
		}
	}

	return length, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/** Output

You cracked the code!
*/

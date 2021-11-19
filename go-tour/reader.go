package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type ErrEmptyBuffer struct {
	length int
}

func (e ErrEmptyBuffer) Error() string {
	return fmt.Sprintf("Read error: slice len can't be %v", e.length)
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	bLength := len(b)

	if bLength == 0 {
		return 0, ErrEmptyBuffer{length: bLength}
	}

	for i := 0; i < bLength; i++ {
		b[i] = 'A'
	}
	return bLength, nil
}

func main() {
	r := MyReader{}

	b := make([]byte, 8)
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])

	reader.Validate(MyReader{})
}

/** Output

n = 8 err = <nil> b = [65 65 65 65 65 65 65 65]
b[:n] = "AAAAAAAA"
OK!
*/

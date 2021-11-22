package util

import (
	"bufio"
	"log"
	"os"
)

func Read(inputFile string) string {
	s := ""

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	i := 0
	for scanner.Scan() {
		if i > 0 {
			s += "\n"
		}
		s += scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return s
}

/**
https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go

There is one caveat: Scanner will error with lines longer than 65536 characters.
If you know your line length is greater than 64K, use the Buffer() method to increase the scanner's capacity:


scanner := bufio.NewScanner(file)

const maxCapacity = longLineLen  // your required line length
buf := make([]byte, maxCapacity)
scanner.Buffer(buf, maxCapacity)

for scanner.Scan() {
*/

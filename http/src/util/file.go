package util

import (
	"bufio"
	"errors"
	"io"
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

func Copy(inputFile string, outputFile string) {
	// open input file
	fi, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		panic(err)
	}
}

/**
https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
*/

func Write(outputFile string, content string, append bool) {
	var f *os.File
	var err error

	if append {
		f, err = os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		f, err = os.Create(outputFile)
	}

	defer f.Close()

	val := content
	data := []byte(val)

	_, err2 := f.Write(data)

	if err2 != nil {
		log.Fatal(err2)
	}
}

func FileExists(fileFullPath string) bool {
	if _, err := os.Stat(fileFullPath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func CopyLineByLine(inputFileFullPath string, outputFileFullPath string) {
	inputFile, err1 := os.Open(inputFileFullPath)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer inputFile.Close()

	outputFile, err2 := os.OpenFile(outputFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

	i := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		data := []byte(scanner.Text())

		if i > 0 {
			outputFile.Write([]byte("\n"))
		}
		_, err2 := outputFile.Write(data)

		if err2 != nil {
			log.Fatal(err2)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

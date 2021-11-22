package main

import (
	"file/src/util"
	"fmt"
)

func main() {
	const SAMPLE_FILE = "resources/sample.txt"

	fmt.Println("Start")

	util.Read(util.RootDir() + SAMPLE_FILE)
}

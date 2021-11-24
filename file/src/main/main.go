package main

import (
	"file/src/util"
	"fmt"
)

func main() {
	const SAMPLE_FILE = "resources/sample.txt"
	const OUTPUT_FILE = "output/output.txt"

	fmt.Println("Start")

	fmt.Println(util.Read(util.RootDir() + SAMPLE_FILE))
	// util.Copy(util.RootDir()+SAMPLE_FILE, util.RootDir()+OUTPUT_FILE)
	// util.Write(util.RootDir()+OUTPUT_FILE, "Hello, World!", true)
}

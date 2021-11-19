package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// pic := make([][]uint8, dy)
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		// line := make([]uint8, dx)
		var line []uint8
		for j := 0; j < dx; j++ {
			//line = append(line, uint8((i+j)/2))
			//line = append(line, uint8(i*j))
			line = append(line, uint8(i^j))
		}
		pic = append(pic, line)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}

/** Output

A very cool picture
*/

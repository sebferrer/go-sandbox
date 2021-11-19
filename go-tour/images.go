package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
	color         uint8
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.color + uint8(x), img.color + uint8(y), 255, 255}
}

func main() {
	m := Image{Width: 125, Height: 125, color: 128}
	pic.ShowImage(m)
}

/** Output

A very cool picture
*/

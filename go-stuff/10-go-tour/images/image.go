// Exercise: Images
// Link: https://go.dev/tour/methods/25

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h, x, y int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(img.x, img.y, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	c := uint8(x + y)
	return color.RGBA{c, c, c, 255}
}

func main() {
	m := Image{w: 256, h: 256}
	pic.ShowImage(m)
}

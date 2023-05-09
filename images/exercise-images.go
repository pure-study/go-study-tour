package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	rect image.Rectangle
}

func main() {
	m := Image{image.Rect(0, 0, 256, 256)}
	pic.ShowImage(m)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.rect
}

func (img Image) At(x, y int) color.Color {
	//v := (uint8(x) + uint8(y)) / uint8(2)
	//v := uint8(x) * uint8(y)
	v := uint8(x) ^ uint8(y)
	//return color.RGBA{v, v, 255, 255}
	return color.RGBA{255, v, v, 255}
}

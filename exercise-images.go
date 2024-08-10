package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct {
    w, h int
}

func (i Image) Bounds() image.Rectangle{
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model{ 
	return color.RGBAModel
}

func (i Image) At(w int, h int) color.Color {
	return color.RGBA{uint8(w), uint8(h), 255, 255}
}


func main() {
	m := Image{16, 16}
	pic.ShowImage(m)
}

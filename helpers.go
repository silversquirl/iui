package iui

import (
	"image"
	"image/color"
)

func scissorBox(box image.Rectangle) (x, y, w, h int32) {
	return int32(box.Min.X), int32(box.Min.Y), int32(box.Dx()), int32(box.Dy())
}

func rgba(c color.Color) (r, g, b, a float32) {
	ri, gi, bi, ai := c.RGBA()
	return float32(ri) / 0xffff, float32(gi) / 0xffff, float32(bi) / 0xffff, float32(ai) / 0xffff
}

package iui

import (
	"image"
	"image/color"

	"github.com/vktec/gll"
)

type Color struct{ color.Color }

func (Color) Size(avail image.Point) image.Point {
	return avail
}
func (c Color) Draw(ctx DrawContext) {
	r, g, b, a := c.RGBA()
	ctx.ClearColor(float32(r)/0xffff, float32(g)/0xffff, float32(b)/0xffff, float32(a)/0xffff)
	ctx.Clear(gll.COLOR_BUFFER_BIT)
}

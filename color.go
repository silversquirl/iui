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
	ctx.ClearColor(rgba(c.Color))
	ctx.Clear(gll.COLOR_BUFFER_BIT)
}

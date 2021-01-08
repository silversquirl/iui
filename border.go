package iui

import (
	"image"
	"image/color"

	"github.com/vktec/gll"
)

type Border struct {
	Color color.Color
	Width int
	Child Component
}

func (border Border) Size(avail image.Point) image.Point {
	w2 := image.Pt(2*border.Width, 2*border.Width)
	size := border.Child.Size(avail.Sub(w2))
	return size.Add(w2)
}

func (border Border) Draw(ctx DrawContext) {
	ctx.ClearColor(rgba(border.Color))
	ctx.Clear(gll.COLOR_BUFFER_BIT)

	box := ctx.Box.Inset(border.Width)
	ctx.Scissor(scissorBox(box))
	border.Child.Draw(ctx.WithBox(box))

	ctx.Scissor(scissorBox(ctx.Box))
}

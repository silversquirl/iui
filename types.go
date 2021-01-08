package iui

import (
	"image"

	"github.com/vktec/gll"
)

type Component interface {
	Size(avail image.Point) image.Point
	Draw(ctx DrawContext)
}

type DrawContext struct {
	gll.GL300
	ShaderRegistry
	Box    image.Rectangle
	Mouse  image.Point
	Clicks []image.Point
}

func (ctx DrawContext) WithBox(box image.Rectangle) DrawContext {
	mouse := ctx.Mouse.Add(ctx.Box.Min)
	if mouse.In(box) {
		ctx.Mouse = mouse.Sub(box.Min)
	} else {
		ctx.Mouse = image.Pt(-1, -1)
	}

	clicks := make([]image.Point, 0, len(ctx.Clicks))
	for _, click := range ctx.Clicks {
		click = click.Add(ctx.Box.Min)
		if click.In(box) {
			clicks = append(clicks, click.Sub(box.Min))
		}
	}
	ctx.Clicks = clicks

	ctx.Box = box
	return ctx
}

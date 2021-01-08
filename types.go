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
	Box    image.Rectangle
	Mouse  image.Point
	Clicks []image.Point
}

func (ctx DrawContext) WithBox(box image.Rectangle) DrawContext {
	ctx.Box = box

	if !ctx.Mouse.In(ctx.Box) {
		ctx.Mouse = image.Pt(-1, -1)
	}

	clicks := make([]image.Point, 0, len(ctx.Clicks))
	for _, click := range ctx.Clicks {
		if click.In(ctx.Box) {
			clicks = append(clicks, click)
		}
	}
	ctx.Clicks = clicks

	return ctx
}

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
	Mouse  Vec2
	Clicks []Vec2
}
type Vec2 struct{ X, Y float64 }

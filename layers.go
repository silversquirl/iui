package iui

import (
	"image"
)

type Layers []Component

func (layers Layers) Size(avail image.Point) (size image.Point) {
	for _, comp := range layers {
		compSize := comp.Size(avail)
		if compSize.X > size.X {
			size.X = compSize.X
		}
		if compSize.Y > size.Y {
			size.Y = compSize.Y
		}
	}
	return size
}

func (layers Layers) Draw(ctx DrawContext) {
	for _, comp := range layers {
		compSize := comp.Size(ctx.Box.Size())
		dx := ctx.Box.Dx() - compSize.X
		dy := ctx.Box.Dy() - compSize.Y
		box := ctx.Box
		box.Min.X += dx / 2
		box.Min.Y += dy / 2
		box.Max.X -= dx - dx/2
		box.Max.Y -= dy - dy/2

		ctx.Scissor(scissorBox(box))
		comp.Draw(ctx.WithBox(box))
	}
	ctx.Scissor(scissorBox(ctx.Box))
}

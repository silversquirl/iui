package iui

import "image"

type Hover struct {
	Normal, Hover Component
}

func (hover Hover) Size(avail image.Point) image.Point {
	size := hover.Normal.Size(avail)
	size2 := hover.Hover.Size(avail)
	if size2.X > size.X {
		size.X = size2.X
	}
	if size2.Y > size.Y {
		size.Y = size2.Y
	}
	return size
}

func (hover Hover) Draw(ctx DrawContext) {
	if ctx.Mouse == image.Pt(-1, -1) {
		hover.Normal.Draw(ctx)
	} else {
		hover.Hover.Draw(ctx)
	}
}

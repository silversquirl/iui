package iui

import "image"

type Col []Component

func (col Col) Size(avail image.Point) image.Point {
	maxx := 0
	for _, size := range col.sizes(avail) {
		if size.X > maxx {
			maxx = size.X
		}
	}
	return image.Pt(maxx, avail.Y)
}

func (col Col) sizes(avail image.Point) []image.Point {
	sizes := make([]image.Point, len(col))
	// FIXME: super naive sizing algorithm
	height := avail.Y / len(col)
	extra := avail.Y % len(col)
	for i, comp := range col {
		h := height
		if extra > 0 {
			extra--
			h++
		}
		sizes[i] = comp.Size(image.Pt(avail.X, h))
	}
	return sizes
}

func (col Col) Draw(ctx DrawContext) {
	sizes := col.sizes(ctx.Box.Size())
	width := ctx.Box.Dx()
	y := ctx.Box.Min.Y
	for i, comp := range col {
		xoff := (width - sizes[i].X) / 2
		if xoff < 0 {
			panic("Child of col is wider than permitted")
		}
		box := image.Rect(ctx.Box.Min.X+xoff, y, ctx.Box.Max.X-xoff, y+sizes[i].Y)
		y += sizes[i].Y

		ctx.Scissor(scissorBox(box))
		comp.Draw(ctx.WithBox(box))
	}
	ctx.Scissor(scissorBox(ctx.Box))
}

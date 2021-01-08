package iui

import "image"

type Col []Component

func (col Col) Size(avail image.Point) image.Point {
	heights := col.heights(avail.Y)
	maxx := 0
	for i, comp := range col {
		size := comp.Size(image.Pt(avail.X, heights[i]))
		if size.X > maxx {
			maxx = size.X
		}
	}
	return image.Pt(maxx, avail.Y)
}

func (col Col) heights(total int) []int {
	heights := make([]int, len(col))
	// FIXME: super naive sizing algorithm
	height := total / len(col)
	extra := total % len(col)
	for i := range heights {
		heights[i] = height
		if extra > 0 {
			extra--
			heights[i]++
		}
	}
	return heights
}

func (col Col) Draw(ctx DrawContext) {
	heights := col.heights(ctx.Box.Dy())
	width := ctx.Box.Dx()
	y := ctx.Box.Max.Y
	for i, comp := range col {
		size := comp.Size(image.Pt(width, heights[i]))
		xoff := (width - size.X) / 2
		yoff := (heights[i] - size.Y) / 2
		if xoff < 0 {
			panic("Child of col is wider than permitted")
		}
		y -= yoff
		box := image.Rect(ctx.Box.Min.X+xoff, y-size.Y, ctx.Box.Max.X-xoff, y)
		y -= size.Y + yoff

		ctx.Scissor(scissorBox(box))
		comp.Draw(ctx.WithBox(box))
	}
	ctx.Scissor(scissorBox(ctx.Box))
}

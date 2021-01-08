package iui

import "image"

type Row []Component

func (row Row) Size(avail image.Point) image.Point {
	widths := row.widths(avail.X)
	maxy := 0
	for i, comp := range row {
		size := comp.Size(image.Pt(widths[i], avail.Y))
		if size.Y > maxy {
			maxy = size.Y
		}
	}
	return image.Pt(avail.X, maxy)
}

func (row Row) widths(total int) []int {
	widths := make([]int, len(row))
	// FIXME: super naive sizing algorithm
	width := total / len(row)
	extra := total % len(row)
	for i := range widths {
		widths[i] = width
		if extra > 0 {
			extra--
			widths[i]++
		}
	}
	return widths
}

func (row Row) Draw(ctx DrawContext) {
	widths := row.widths(ctx.Box.Dx())
	height := ctx.Box.Dy()
	x := ctx.Box.Min.X
	for i, comp := range row {
		size := comp.Size(image.Pt(widths[i], height))
		xoff := (widths[i] - size.X) / 2
		yoff := (height - size.Y) / 2
		if yoff < 0 {
			panic("Child of row is taller than permitted")
		}
		x += xoff
		box := image.Rect(x, ctx.Box.Min.Y+yoff, x+size.X, ctx.Box.Max.Y-yoff)
		x += size.X + xoff

		ctx.Scissor(scissorBox(box))
		comp.Draw(ctx.WithBox(box))
	}
	ctx.Scissor(scissorBox(ctx.Box))
}

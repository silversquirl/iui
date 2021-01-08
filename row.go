package iui

import "image"

type Row []Component

func (row Row) Size(avail image.Point) image.Point {
	maxy := 0
	for _, size := range row.sizes(avail) {
		if size.Y > maxy {
			maxy = size.Y
		}
	}
	return image.Pt(avail.X, maxy)
}

func (row Row) sizes(avail image.Point) []image.Point {
	sizes := make([]image.Point, len(row))
	// FIXME: super naive sizing algorithm
	width := avail.X / len(row)
	extra := avail.X % len(row)
	for i, comp := range row {
		w := width
		if extra > 0 {
			extra--
			w++
		}
		sizes[i] = comp.Size(image.Pt(w, avail.Y))
	}
	return sizes
}

func (row Row) Draw(ctx DrawContext) {
	sizes := row.sizes(ctx.Box.Size())
	height := ctx.Box.Dy()
	x := ctx.Box.Min.X
	for i, comp := range row {
		yoff := (height - sizes[i].Y) / 2
		if yoff < 0 {
			panic("Child of row is taller than permitted")
		}
		box := image.Rect(x, ctx.Box.Min.Y+yoff, x+sizes[i].X, ctx.Box.Max.Y-yoff)
		x += sizes[i].X

		ctx.Scissor(scissorBox(box))
		comp.Draw(ctx.WithBox(box))
	}
	ctx.Scissor(scissorBox(ctx.Box))
}
